package main

import (
	"fmt"

	"github.com/Zondax/zindexer/components/connections/data_store"
	"github.com/bytedance/sonic"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"
	"github.com/zondax/golem/pkg/cli"
	"go.uber.org/zap"
)

func download(c *cli.CLI, cmd *cobra.Command, _ []string) {
	zap.S().Infof(c.GetVersionString())

	config, err := cli.LoadConfig[Config]()
	if err != nil {
		zap.S().Errorf("Error loading config: %s", err)
		return
	}
	logType, err := cmd.Flags().GetString("type")
	if err != nil {
		zap.S().Errorf("Error loading type: %s", err)
		return
	}
	outPath, err := cmd.Flags().GetString("outPath")
	if err != nil {
		zap.S().Errorf("Error loading outPath: %s", err)
		return
	}
	format, err := cmd.Flags().GetString("compress")
	if err != nil {
		zap.S().Errorf("Error loading compress: %s", err)
		return
	}

	heights, err := cmd.Flags().GetUintSlice("heights")
	if err != nil {
		zap.S().Errorf("Error loading heights: %s", err)
		return
	}

	rpcClient, err := newFilecoinRPCClient(config.NodeURL, config.NodeToken)
	if err != nil {
		zap.S().Error(err)
		return
	}

	for _, h := range heights {
		_, dataJson, err := getTrace(uint64(h), logType, rpcClient)
		if err != nil {
			zap.S().Error(err)
			continue
		}

		out := dataJson
		fname := fmt.Sprintf("%s_%d.json", logType, h)
		if format != "" {
			out, err = compress(format, dataJson)
			if err != nil {
				zap.S().Error(err)
				return
			}
			fname = fmt.Sprintf("%s_%d.json.%s", logType, h, format)
		}

		if err := writeToFile(outPath, fname, out); err != nil {
			zap.S().Error(err)
			return
		}

	}

}

func getTrace(height uint64, logType string, rpcClient *RPCClient) (*lotusChainTypes.TipSet, []byte, error) {
	var data any
	var tipset *lotusChainTypes.TipSet
	var err error
	switch logType {
	case "traces":
		tipset, data, err = getTraceFileByHeight(height, rpcClient.client)
	case "tipset":
		data, err = getTipsetFileByHeight(height, lotusChainTypes.EmptyTSK, rpcClient.client)
	case "ethlog":
		data, err = getEthLogsByHeight(height, rpcClient.client)
	case "nativelog":
		data, err = getNativeLogsByHeight(height, rpcClient.client)
	case "metadata":
		data, err = getMetadata(rpcClient)
	}

	if err != nil {
		return nil, nil, err
	}

	dataJson, err := sonic.Marshal(data)
	if err != nil {
		return nil, nil, err
	}
	return tipset, dataJson, nil
}

func getTraceFromDataStore(height uint64, logType string, dsClient *data_store.DataStoreClient, config *Config) ([]byte, error) {
	storePath := fmt.Sprintf("%s/%s", config.S3Bucket, config.S3RawDataPath)
	name := fmt.Sprintf("%s_%012d.json.s2", logType, height)
	zap.S().Infof("Getting trace from data store: %s/%s", storePath, name)

	data, err := dsClient.Client.GetFile(name, storePath)
	zap.S().Infof("Data: %v", len(data))
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	decompressed, err := decompress(data)
	if err != nil {
		return nil, err
	}

	return decompressed, nil
}
