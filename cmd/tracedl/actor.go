package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/filecoin-project/go-address"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	filMetrics "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/cli"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/zcache"
	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"go.uber.org/zap"
)

type ActorList struct {
	ActorName string `json:"tx_to"`
	Height    uint64 `json:"height"`
}

func actorscmd(c *cli.CLI, cmd *cobra.Command, _ []string) {
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
	actors, err := cmd.Flags().GetString("actors")
	if err != nil {
		zap.S().Errorf("Error loading actors: %s", err)
		return
	}
	outFile, err := cmd.Flags().GetString("out")
	if err != nil {
		zap.S().Errorf("Error loading actors: %s", err)
		return
	}

	// read file actors
	actorsFile, err := os.ReadFile(actors)
	if err != nil {
		zap.S().Errorf("Error reading actors: %s", err)
		return
	}
	actorsList := []ActorList{}
	err = json.Unmarshal(actorsFile, &actorsList)
	if err != nil {
		zap.S().Errorf("Error unmarshalling actors: %s", err)
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

	var outActors []ActorList
	for _, actor := range actorsList {
		var tipset *lotusChainTypes.TipSet
		if useDataStore {
			tipset, _, err = downloadTraceFromDataStore(int64(actor.Height), tracesPath, dataStore, rpcClient, config, logger)
		} else {
			tipset, _, err = downloadTraceIfNotExists(int64(actor.Height), tracesPath, rpcClient)
		}
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		actorAddress, err := address.NewFromString(actor.ActorName)
		if err != nil {
			zap.S().Errorf("Error parsing actor address: %s", err)
			continue
		}
		fmt.Println(actorAddress)
		_, actorName, err := helper.GetActorNameFromAddress(actorAddress, int64(actor.Height), tipset.Key())
		if err != nil {
			zap.S().Errorf("Error getting actor name from address: %s", err)
			continue
		}
		fmt.Println(actorName)
		outActors = append(outActors, ActorList{
			ActorName: actorName,
			Height:    actor.Height,
		})
	}

	gotStr, _ := json.MarshalIndent(outActors, "", "  ")
	os.WriteFile(fmt.Sprintf("%s/%s", tracesPath, outFile), gotStr, fs.ModePerm)

}
