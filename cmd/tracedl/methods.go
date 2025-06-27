package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/bytedance/sonic"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	filMetrics "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/golem/pkg/cli"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/zcache"
	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"go.uber.org/zap"
)

func methods(c *cli.CLI, cmd *cobra.Command, _ []string) {
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

	heightsFile, err := cmd.Flags().GetString("heights")
	if err != nil {
		zap.S().Errorf("Error loading heights: %s", err)
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
		zap.S().Error(err)
		return
	}
	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(rpcClient.client)
	helper := helper.NewHelper(lib, actorsCache, rpcClient.client, logger, nil)
	dataStore, err := getDataStoreClient(config)
	if err != nil {
		zap.S().Error(err)
		return
	}
	data, err := os.ReadFile(heightsFile)
	if err != nil {
		zap.S().Error(err)
		return
	}

	heights := map[string][]map[string]float64{}
	json.Unmarshal(data, &heights)

	resps := [][]map[string]any{}
	for _, tmpheight := range heights["heights"] {
		height := int64(tmpheight["height"])
		var tipset *lotusChainTypes.TipSet
		var data []byte

		tipset, data, err = downloadTraceFromDataStore(int64(height), tracesPath, dataStore, rpcClient, config, logger)

		if err != nil {
			zap.S().Errorf("error downloading trace: %s", err)
			return
		}

		var computeState *typesV2.ComputeStateOutputV2
		err = sonic.Unmarshal(data, &computeState)
		if err != nil {
			zap.S().Errorf("error unmarshalling traces: %s", err)
			return
		}

		resp := []map[string]any{}
		subcalls := []typesV2.ExecutionTraceV2{}
		for _, trace := range computeState.Trace {
			if len(trace.ExecutionTrace.Subcalls) > 0 {
				subcalls = append(subcalls, trace.ExecutionTrace.Subcalls...)
			}

			_, foundActorName, err := helper.GetActorNameFromAddress(trace.Msg.To, int64(height), tipset.Key())
			if err != nil {
				zap.S().Errorf("error getting actor name: %s", err)
				return
			}

			msg := parser.LotusMessage{
				From:   trace.Msg.From,
				To:     trace.Msg.To,
				Method: trace.Msg.Method,
				Params: trace.Msg.Params,
			}

			txType, err := helper.GetMethodName(&msg, int64(height), tipset.Key())
			if err != nil {
				zap.S().Errorf("error getting method name: %s", err)
				fmt.Println("parser-fix: ", foundActorName, msg.Method, height)
				resp = append(resp, map[string]any{
					"actor":  foundActorName,
					"method": msg.Method,
					"height": height,
					"error":  err.Error(),
				})
			}

			if err == nil && txType == parser.UnknownStr {
				resp = append(resp, map[string]any{
					"actor":  foundActorName,
					"method": msg.Method,
					"height": height,
					"error":  "unknown method",
				})
				fmt.Println("parser-fix: ", foundActorName, msg.Method, height)
			}
		}
		resps = append(resps, resp)
	}

	tmp, _ := json.MarshalIndent(resps, "", "  ")
	os.WriteFile("methods.json", tmp, 0644)
}
