package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/bytedance/sonic"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/spf13/cobra"

	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	v2 "github.com/zondax/fil-parser/actors/v2"
	logger2 "github.com/zondax/fil-parser/logger"
	filMetrics "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/golem/pkg/cli"
	"github.com/zondax/golem/pkg/logger"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/zcache"
	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"go.uber.org/zap"
)

type BatchHeight struct {
	TxTo   string `json:"tx_to"`
	Height int    `json:"height"`
}

type UnknownMethod struct {
	Height  int64  `json:"height"`
	Method  int64  `json:"method"`
	Address string `json:"address"`
	Actor   string `json:"actor"`
}

func updateFile(file string, data []UnknownMethod) error {
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, out, fs.ModePerm)
}

func batch(c *cli.CLI, cmd *cobra.Command, _ []string) {
	zap.S().Infof(c.GetVersionString())
	logger := logger2.GetSafeLogger(nil)
	config, err := cli.LoadConfig[Config]()
	if err != nil {
		zap.S().Errorf("Error loading config: %s", err)
		return
	}
	tracesPath, err := cmd.Flags().GetString("tracesPath")
	if err != nil {
		zap.S().Errorf("Error loading tracesPath: %s", err)
		return
	}
	useDataStore, err := cmd.Flags().GetBool("useDataStore")
	if err != nil {
		zap.S().Errorf("Error loading useDataStore: %s", err)
		return
	}
	inFile, err := cmd.Flags().GetString("inFile")
	if err != nil {
		zap.S().Errorf("Error loading inFile: %s", err)
		return
	}
	outFile, err := cmd.Flags().GetString("outFile")
	if err != nil {
		zap.S().Errorf("Error loading outFile: %s", err)
		return
	}

	heights := []BatchHeight{}
	inData, err := os.ReadFile(inFile)
	if err != nil {
		zap.S().Errorf("Error reading inFile: %s", err)
		return
	}
	err = json.Unmarshal(inData, &heights)
	if err != nil {
		zap.S().Errorf("Error unmarshalling inFile: %s", err)
		return
	}

	parseSubTxs, err := cmd.Flags().GetBool("parseSubTxs")
	if err != nil {
		zap.S().Errorf("Error loading parseSubTxs: %s", err)
		return
	}

	rpcClient, err := newFilecoinRPCClient(config.NodeURL, config.NodeToken)
	if err != nil {
		zap.S().Error(err)
		return
	}

	actorsCache, err := cache.SetupActorsCache(common.DataSource{
		Node: rpcClient.client,
		Config: common.DataSourceConfig{
			NetworkName: config.NetworkName,
			Cache: &common.CacheConfig{
				CombinedConfig: &zcache.CombinedConfig{
					Remote: &zcache.RemoteConfig{
						Addr:     config.RedisAddr,
						Password: config.RedisPassword,
					},
					Local:              &zcache.LocalConfig{},
					IsRemoteBestEffort: true,
					GlobalPrefix:       config.NetworkName,
					GlobalMetricServer: metrics2.NewNoopMetrics(),
				},
			},
		},
	}, logger, filMetrics.NewMetricsClient(metrics2.NewNoopMetrics()), golemBackoff.New().
		WithMaxAttempts(3).
		WithMaxDuration(1*time.Second).
		WithInitialDuration(1*time.Second))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(rpcClient.client)
	helper := helper.NewHelper(lib, actorsCache, rpcClient.client, logger, filMetrics.NewMetricsClient(metrics2.NewNoopMetrics()))
	dataStore, err := getDataStoreClient(config)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	unknownMethods := []UnknownMethod{}
	unknownMethodsMap := map[string]bool{}
	for _, height := range heights {
		var tipset *lotusChainTypes.TipSet
		var data []byte
		if useDataStore {
			tipset, data, err = downloadTraceFromDataStore(int64(height.Height), tracesPath, dataStore, rpcClient, config, logger)
		} else {
			tipset, data, err = downloadTraceIfNotExists(int64(height.Height), tracesPath, rpcClient)
		}
		if err != nil {
			logger.Error(err.Error())
			return
		}

		var computeState *typesV2.ComputeStateOutputV2
		err = sonic.Unmarshal(data, &computeState)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		resp := []map[string]any{}
		subcalls := []typesV2.ExecutionTraceV2{}
		for traceId, trace := range computeState.Trace {
			if len(trace.ExecutionTrace.Subcalls) > 0 {
				subcalls = append(subcalls, trace.ExecutionTrace.Subcalls...)
			}
			if height.TxTo != "" && trace.Msg.To.String() != height.TxTo {
				continue
			}
			_, foundActorName, err := helper.GetActorNameFromAddress(trace.Msg.To, int64(height.Height), tipset.Key())
			if err != nil {
				logger.Error(err.Error())
				continue
			}

			msg := parser.LotusMessage{
				From:   trace.Msg.From,
				To:     trace.Msg.To,
				Method: trace.Msg.Method,
				Params: trace.Msg.Params,
			}
			rct := parser.LotusMessageReceipt{
				ExitCode: trace.MsgRct.ExitCode,
				Return:   trace.MsgRct.Return,
			}
			actorParser := v2.NewActorParser(config.NetworkName, helper, logger, filMetrics.NewMetricsClient(metrics2.NewNoopMetrics()))

			methodName, err := v2.GetMethodName(context.Background(), msg.Method, foundActorName, int64(height.Height), config.NetworkName, helper, logger)
			if methodName == parser.UnknownStr {
				unknown := UnknownMethod{
					Height:  int64(height.Height),
					Method:  int64(msg.Method),
					Address: trace.Msg.To.String(),
					Actor:   foundActorName,
				}
				key := fmt.Sprintf("%s:%d", unknown.Actor, unknown.Method)
				if _, ok := unknownMethodsMap[key]; !ok {
					unknownMethods = append(unknownMethods, unknown)
					unknownMethodsMap[key] = true
				}
			}
			// methodName, err := helper.GetMethodName(&msg, int64(height), tipset.Key())
			if err != nil {
				logger.Error(err.Error())
				continue
			}

			_, got, _, err := actorParser.GetMetadata(context.Background(), foundActorName, methodName, &msg, trace.Msg.Cid(), &rct, int64(height.Height), tipset.Key())
			if err != nil {
				logger.Error(err.Error())
			}
			resp = append(resp, map[string]any{
				"actorName":   foundActorName,
				"addressTo":   trace.Msg.To.String(),
				"addressFrom": trace.Msg.From.String(),
				"methodName":  methodName,
				"traceId":     traceId,
				"got":         got,
			})
		}

		if parseSubTxs {
			for _, trace := range subcalls {
				fmt.Printf("parsing subcalls: %d\n", len(trace.Subcalls))
				resps, _, resUnknown := parseSubCall2(0, int64(height.Height), config.NetworkName, "", "", height.TxTo, trace, tipset, helper, logger)
				resp = append(resp, resps...)
				for _, unknown := range resUnknown {
					key := fmt.Sprintf("%s:%d", unknown.Actor, unknown.Method)
					if _, ok := unknownMethodsMap[key]; !ok {
						unknownMethods = append(unknownMethods, unknown)
						unknownMethodsMap[key] = true
					}
				}
			}
		}

		gotStr, _ := json.MarshalIndent(resp, "", "  ")
		os.WriteFile(fmt.Sprintf("%s/resp_%d.json", tracesPath, height.Height), gotStr, fs.ModePerm)
		updateFile(outFile, unknownMethods)
	}

}
func parseSubCall2(level, height int64, network, actorName, actorMethod string, actorAddress string, trace typesV2.ExecutionTraceV2, tipset *lotusChainTypes.TipSet, helper *helper.Helper, logger *logger.Logger) ([]map[string]any, []typesV2.ExecutionTraceV2, []UnknownMethod) {
	for i := 0; i < int(level); i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("level: %d, actorName: %s, actorAddress: %s, subcalls: %d\n", level, actorName, trace.Msg.To.String(), len(trace.Subcalls))
	subcalls := []typesV2.ExecutionTraceV2{}
	n := level + 1
	res := []map[string]any{}
	resUnknown := []UnknownMethod{}
	for _, subcall := range trace.Subcalls {
		subResp, _, subUnknown := parseSubCall2(n, height, network, actorName, actorMethod, actorAddress, subcall, tipset, helper, logger)
		res = append(res, subResp...)
		resUnknown = append(resUnknown, subUnknown...)
	}
	if actorAddress != "" && trace.Msg.To.String() != actorAddress {
		return res, subcalls, resUnknown
	}
	_, foundActorName, err := helper.GetActorNameFromAddress(trace.Msg.To, int64(height), tipset.Key())
	if err != nil {
		logger.Errorf("error getting actor name from address: %s : %s", trace.Msg.To.String(), err)
		return res, subcalls, resUnknown
	}
	if actorName != "" && actorName != foundActorName {
		return res, subcalls, resUnknown
	}
	msg := parser.LotusMessage{
		From:   trace.Msg.From,
		To:     trace.Msg.To,
		Method: trace.Msg.Method,
		Params: trace.Msg.Params,
	}
	rct := parser.LotusMessageReceipt{
		ExitCode: trace.MsgRct.ExitCode,
		Return:   trace.MsgRct.Return,
	}
	actorParser := v2.NewActorParser(network, helper, logger, filMetrics.NewMetricsClient(metrics2.NewNoopMetrics()))
	methodName, err := v2.GetMethodName(context.Background(), msg.Method, foundActorName, int64(height), network, helper, logger)

	if methodName == parser.UnknownStr {
		resUnknown = append(resUnknown, UnknownMethod{
			Height:  height,
			Method:  int64(msg.Method),
			Address: trace.Msg.To.String(),
			Actor:   foundActorName,
		})
	}
	// methodName, err := helper.GetMethodName(&msg, int64(height), tipset.Key())
	if err != nil {
		logger.Errorf("error getting method name: %s", err)
		return res, subcalls, resUnknown
	}

	if actorMethod != "" && methodName != actorMethod {
		return res, subcalls, resUnknown
	}

	_, got, _, err := actorParser.GetMetadata(context.Background(), foundActorName, methodName, &msg, cid.Undef, &rct, int64(height), tipset.Key())
	if err != nil {
		logger.Errorf("error getting metadata: %s, actorName: %s, address:%s, methodName: %s, traceId: %s", err, foundActorName, trace.Msg.To.String(), methodName, "SUBCALL")
		// fmt.Println(hex.EncodeToString(trace.Msg.Params))
		// return mr, subcalls
	}
	if len(got) == 0 {
		got["ParamsRaw"] = trace.Msg.Params
		got["ReturnRaw"] = trace.MsgRct.Return
	}
	res = append(res, map[string]any{
		"actorName":   foundActorName,
		"addressTo":   trace.Msg.To.String(),
		"addressFrom": trace.Msg.From.String(),
		"methodName":  methodName,
		"traceId":     fmt.Sprintf("SUBCALL_%d", level),
		"got":         got,
	})
	return res, subcalls, resUnknown
}
