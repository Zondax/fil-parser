package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/Zondax/zindexer/components/connections/data_store"
	"github.com/bytedance/sonic"
	"github.com/filecoin-project/go-state-types/abi"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/spf13/cobra"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/types"

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

func parse(c *cli.CLI, cmd *cobra.Command, _ []string) {
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
	height, err := cmd.Flags().GetUint64("height")
	if err != nil {
		zap.S().Errorf("Error loading heights: %s", err)
		return
	}
	useDataStore, err := cmd.Flags().GetBool("useDataStore")
	if err != nil {
		zap.S().Errorf("Error loading useDataStore: %s", err)
		return
	}
	actorAddress, err := cmd.Flags().GetString("actorAddress")
	if err != nil {
		zap.S().Errorf("Error loading actorAddress: %s", err)
		return
	}
	actorName, err := cmd.Flags().GetString("actorName")
	if err != nil {
		zap.S().Errorf("Error loading actorName: %s", err)
		return
	}
	actorMethod, err := cmd.Flags().GetString("actorMethod")
	if err != nil {
		zap.S().Errorf("Error loading actorMethod: %s", err)
		return
	}
	parseSubTxs, err := cmd.Flags().GetBool("parseSubTxs")
	if err != nil {
		zap.S().Errorf("Error loading parseSubTxs: %s", err)
		return
	}
	txCid, err := cmd.Flags().GetString("txCid")
	if err != nil {
		zap.S().Errorf("Error loading txCid: %s", err)
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
		WithInitialDuration(1*time.Second).Linear())
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
	var tipset *lotusChainTypes.TipSet
	var data []byte
	if useDataStore {
		tipset, data, err = downloadTraceFromDataStore(int64(height), tracesPath, dataStore, rpcClient, config, logger)
	} else {
		tipset, data, err = downloadTraceIfNotExists(int64(height), tracesPath, rpcClient)
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

	tmp, _ := sonic.Marshal(computeState)
	os.WriteFile("min.json", tmp, fs.ModePerm)

	resp := []map[string]any{}
	subcalls := []typesV2.ExecutionTraceV2{}
	actorParser := v2.NewActorParser(config.NetworkName, helper, logger, filMetrics.NewMetricsClient(metrics2.NewNoopMetrics()))

	for traceId, trace := range computeState.Trace {
		if txCid != "" && trace.MsgCid.String() != txCid {
			fmt.Println("skipping tx cid", trace.Msg.Cid().String())
			continue
		}
		if len(trace.ExecutionTrace.Subcalls) > 0 {
			subcalls = append(subcalls, trace.ExecutionTrace.Subcalls...)
		}
		if actorAddress != "" && trace.Msg.To.String() != actorAddress {
			continue
		}
		if trace.Msg.To.String() == "f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a" {
			fmt.Println("trace.Msg.To", trace.Msg.To.String())
			fmt.Println("trace.Msg.From", trace.Msg.From.String())
			fmt.Println("trace.Msg.Method", trace.Msg.Method)
			fmt.Println("trace.Msg.Params", trace.Msg.Params)
			fmt.Println("trace.MsgRct.ExitCode", trace.MsgRct.ExitCode)
		}
		_, foundActorName, err := helper.GetActorNameFromAddress(trace.Msg.To, int64(height), tipset.Key())
		if err != nil {
			logger.Error(err.Error())
			return
		}
		if actorName != "" && actorName != foundActorName {
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
		methodName, err := v2.GetMethodName(context.Background(), msg.Method, foundActorName, int64(height), config.NetworkName, helper, logger)
		// methodName, err := helper.GetMethodName(&msg, int64(height), tipset.Key())
		if err != nil {
			logger.Error(err.Error())
			return
		}
		if actorMethod != "" && methodName != actorMethod {
			continue
		}

		_, got, addressInfo, err := actorParser.GetMetadata(context.Background(), foundActorName, methodName, &msg, trace.Msg.Cid(), &rct, int64(height), tipset.Key())
		if err != nil {
			logger.Error(err.Error())
		}
		if addressInfo != nil {
			actorsCache.StoreAddressInfo(*addressInfo)
		}
		et := &types.ExtendedTipSet{
			TipSet: *tipset,
		}
		blockCid, err := v2.GetBlockCidFromMsgCid(trace.MsgCid.String(), methodName, got, et, logger)
		if err != nil {
			logger.Error(err.Error())
		}
		resp = append(resp, map[string]any{
			"actorName":   foundActorName,
			"addressTo":   trace.Msg.To.String(),
			"addressFrom": trace.Msg.From.String(),
			"methodName":  methodName,
			"methodNum":   trace.Msg.Method,
			"traceId":     traceId,
			"got":         got,
			"blockCid":    blockCid,
			"addrsInfo":   addressInfo,
		})
	}

	if parseSubTxs {
		for _, trace := range subcalls {

			fmt.Printf("parsing subcalls: %d\n", len(trace.Subcalls))
			resps, _ := parseSubCall(0, int64(height), config.NetworkName, actorName, actorMethod, actorAddress, trace, tipset, helper, logger)
			resp = append(resp, resps...)
		}
	}

	gotStr, _ := json.MarshalIndent(resp, "", "  ")
	os.WriteFile(fmt.Sprintf("%s/resp_%d.json", tracesPath, height), gotStr, fs.ModePerm)

}

func parseSubCall(level, height int64, network, actorName, actorMethod string, actorAddress string, trace typesV2.ExecutionTraceV2, tipset *lotusChainTypes.TipSet, helper *helper.Helper, logger *logger.Logger) ([]map[string]any, []typesV2.ExecutionTraceV2) {
	for i := 0; i < int(level); i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("level: %d, actorName: %s, actorAddress: %s, subcalls: %d\n", level, actorName, trace.Msg.To.String(), len(trace.Subcalls))
	subcalls := []typesV2.ExecutionTraceV2{}
	n := level + 1
	res := []map[string]any{}
	for _, subcall := range trace.Subcalls {
		subResp, _ := parseSubCall(n, height, network, actorName, actorMethod, actorAddress, subcall, tipset, helper, logger)
		res = append(res, subResp...)
	}
	if actorAddress != "" && trace.Msg.To.String() != actorAddress {
		return res, subcalls
	}
	if trace.Msg.To.String() == "f067253" {
		fmt.Println("trace.Msg.To", trace.Msg.To.String())
		fmt.Println("trace.Msg.From", trace.Msg.From.String())
		fmt.Println("trace.Msg.Method", trace.Msg.Method)
		fmt.Println("trace.Msg.Params", trace.Msg.Params)
		fmt.Println("trace.MsgRct.ExitCode", trace.MsgRct.ExitCode)
	}

	_, foundActorName, err := helper.GetActorNameFromAddress(trace.Msg.To, int64(height), tipset.Key())
	if err != nil || foundActorName == "" {
		if trace.InvokedActor != nil {
			foundActorName, err = helper.GetActorNameFromCid(trace.InvokedActor.State.Code, int64(tipset.Height()))
			if err != nil {
				logger.Errorf("Error when trying to get actor name from cid in tx cid using invoked actor: %v", err)
			}
		} else {
			logger.Errorf("error getting actor name from address: %s : %s", trace.Msg.To.String(), err)
			return res, subcalls
		}
	}
	if actorName != "" && actorName != foundActorName {
		return res, subcalls
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
	// methodName, err := helper.GetMethodName(&msg, int64(height), tipset.Key())
	if err != nil {
		logger.Errorf("error getting method name: %s", err)
		return res, subcalls
	}
	if actorMethod != "" && methodName != actorMethod {
		return res, subcalls
	}

	_, got, addrsInfo, err := actorParser.GetMetadata(context.Background(), foundActorName, methodName, &msg, cid.Undef, &rct, int64(height), tipset.Key())
	if err != nil {
		logger.Errorf("error getting metadata: %s, actorName: %s, address:%s, methodName: %s, traceId: %s", err, foundActorName, trace.Msg.To.String(), methodName, "SUBCALL")
		// fmt.Println(hex.EncodeToString(trace.Msg.Params))
		// return mr, subcalls
	}
	if len(got) == 0 {
		if got == nil {
			got = map[string]any{}
		}
		got["ParamsRaw"] = trace.Msg.Params
		got["ReturnRaw"] = trace.MsgRct.Return
	}
	res = append(res, map[string]any{
		"actorName":   foundActorName,
		"addressTo":   trace.Msg.To.String(),
		"addressFrom": trace.Msg.From.String(),
		"methodName":  methodName,
		"methodNum":   trace.Msg.Method,
		"traceId":     fmt.Sprintf("SUBCALL_%d", level),
		"got":         got,
		"addrsInfo":   addrsInfo,
	})
	return res, subcalls
}

func downloadTraceIfNotExists(height int64, outPath string, rpcClient *RPCClient) (*lotusChainTypes.TipSet, []byte, error) {
	logger.Infof("Downloading trace from node for height: %d", height)
	traceFile := fmt.Sprintf("%s/traces_%d.json.gz", outPath, height)
	tipset, err := rpcClient.client.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(height), lotusChainTypes.EmptyTSK)
	if err != nil {
		return nil, nil, err
	}
	if _, err := os.Stat(traceFile); os.IsNotExist(err) {
		zap.S().Infof("Downloading trace for height: %d", height)
		tipset, dataJson, err := getTrace(uint64(height), "traces", rpcClient)
		if err != nil {
			return nil, nil, err
		}
		out, err := compress("gz", dataJson)
		if err != nil {
			return nil, nil, err
		}
		os.WriteFile(traceFile, out, fs.ModePerm)
		return tipset, dataJson, nil
	}
	zap.S().Infof("Trace file already exists for height: %d", height)
	data, err := readGzFile(traceFile)
	if err != nil {
		return nil, nil, err
	}

	return tipset, data, nil
}

func downloadTraceFromDataStore(height int64, outPath string, dataStore *data_store.DataStoreClient, rpcClient *RPCClient, config *Config, logger *logger.Logger) (*lotusChainTypes.TipSet, []byte, error) {
	logger.Infof("Downloading trace from data store for height: %d", height)
	traceFile := fmt.Sprintf("%s/traces_%d.json.gz", outPath, height)
	tipset, err := rpcClient.client.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(height), lotusChainTypes.EmptyTSK)
	if err != nil {
		return nil, nil, err
	}
	if _, err := os.Stat(traceFile); os.IsNotExist(err) {
		dataJson, err := getTraceFromDataStore(uint64(height), "traces", dataStore, config)
		if err != nil {
			os.Exit(1)
			return nil, nil, err
		}
		out, err := compress("gz", dataJson)
		if err != nil {
			return nil, nil, err
		}
		os.WriteFile(traceFile, out, fs.ModePerm)
		return tipset, dataJson, nil
	}

	logger.Infof("Trace file already exists for height: %d", height)
	data, err := readGzFile(traceFile)
	if err != nil {
		return nil, nil, err
	}

	return tipset, data, nil
}
