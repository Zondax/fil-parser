package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bytedance/sonic"
	"github.com/filecoin-project/go-address"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"

	logger2 "github.com/zondax/fil-parser/logger"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/golem/pkg/cli"
	"go.uber.org/zap"
)

func miner(c *cli.CLI, cmd *cobra.Command, _ []string) {
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
	fmt.Println("useDataStore", useDataStore)
	minerAddress, err := cmd.Flags().GetString("minerAddress")
	if err != nil {
		zap.S().Errorf("Error loading actorAddress: %s", err)
		return
	}

	addr, err := address.NewFromString(minerAddress)
	if err != nil {
		zap.S().Error(err)
		return
	}
	rpcClient, err := newFilecoinRPCClient(config.NodeURL, config.NodeToken)
	if err != nil {
		zap.S().Error(err)
		return
	}

	dataStore, err := getDataStoreClient(config)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var tipset *lotusChainTypes.TipSet
	var data []byte
	if false {
		tipset, data, err = downloadTraceFromDataStore(int64(height), tracesPath, dataStore, rpcClient, config, logger)
	} else {
		tipset, data, err = downloadTraceIfNotExists(int64(height), tracesPath, rpcClient)
	}
	if err != nil {
		logger.Error(err.Error())
		return
	}

	tipData, err := json.Marshal(tipset)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	os.WriteFile(fmt.Sprintf("tipset_%d.json", height), tipData, os.ModePerm)
	var computeState *typesV2.ComputeStateOutputV2
	err = sonic.Unmarshal(data, &computeState)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	ctx := cmd.Context()
	results := map[string]any{}
	actor, err := rpcClient.client.StateGetActor(ctx, addr, tipset.Key())
	if err != nil {
		logger.Error(err.Error())
	}
	results["actor"] = actor
	info, err := rpcClient.client.StateMinerInfo(ctx, addr, tipset.Key())
	if err != nil {
		logger.Error(err.Error())

	}
	results["minerInfo"] = info
	as, err := rpcClient.client.StateMinerActiveSectors(ctx, addr, tipset.Key())
	if err != nil {
		logger.Error(err.Error())

	}
	results["activeSectors"] = as
	sc, err := rpcClient.client.StateMinerSectorCount(ctx, addr, tipset.Key())
	if err != nil {
		logger.Error(err.Error())

	}
	results["sectorCount"] = sc
	ac, err := rpcClient.client.StateMinerSectors(ctx, addr, nil, tipset.Key())
	if err != nil {
		logger.Error(err.Error())

	}
	results["allSectors"] = ac
	mp, err := rpcClient.client.StateMinerPower(ctx, addr, tipset.Key())
	if err != nil {
		logger.Error(err.Error())

	}
	results["minerPower"] = mp

	out, _ := json.MarshalIndent(results, "", " ")
	os.WriteFile(fmt.Sprintf("miner_%d.json", height), out, os.ModePerm)
}
