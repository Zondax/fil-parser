package main

import (
	"fmt"

	logger2 "github.com/zondax/fil-parser/logger"

	"github.com/bytedance/sonic"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"
	"github.com/zondax/golem/pkg/cli"
)

func GetStartCommand(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get",
		Run: func(cmd *cobra.Command, args []string) {
			download(c, cmd, args)
		},
	}
	cmd.Flags().String("type", "traces", "--type traces")
	cmd.Flags().String("outPath", ".", "--outPath ../")
	cmd.Flags().String("compress", "gz", "--compress s2")
	cmd.Flags().UintSlice("heights", []uint{387926}, "--heights 387926")
	cmd.Flags().Bool("useDataStore", false, "--useDataStore true")
	return cmd
}

func GetActorParseCommand(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "actor-parse",
		Short: "Actor Parse",
		Run: func(cmd *cobra.Command, args []string) {
			parse(c, cmd, args)
		},
	}
	cmd.Flags().String("tracesPath", ".", "--tracesPath .")
	cmd.Flags().Uint64("height", 387926, "--height 387926")
	cmd.Flags().Bool("useDataStore", false, "--useDataStore true")
	cmd.Flags().String("actorAddress", "", "--actorAddress f01")
	cmd.Flags().String("actorName", "", "--actorName account")
	cmd.Flags().String("actorMethod", "", "--actorMethod Constructor")
	cmd.Flags().Bool("parseSubTxs", false, "--parseSubTxs true")
	cmd.Flags().String("txCid", "", "--txCid xx")
	return cmd
}

func GetMinerStateCommand(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "miner",
		Short: "miner",
		Run: func(cmd *cobra.Command, args []string) {
			miner(c, cmd, args)
		},
	}
	cmd.Flags().String("tracesPath", ".", "--tracesPath .")
	cmd.Flags().Uint64("height", 387926, "--height 387926")
	cmd.Flags().Bool("useDataStore", false, "--useDataStore true")
	cmd.Flags().String("minerAddress", "", "--minerAddress f01")

	return cmd
}

func GetActorCommand(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "actors",
		Short: "Actor",
		Run: func(cmd *cobra.Command, args []string) {
			actorscmd(c, cmd, args)
		},
	}
	cmd.Flags().String("tracesPath", ".", "--tracesPath .")
	cmd.Flags().Bool("useDataStore", false, "--useDataStore true")
	cmd.Flags().String("actors", "", "--actors ./actors.json")
	cmd.Flags().String("out", "out.json", "--out out.json")

	return cmd
}

func GetBatchParseCommand(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch-parse",
		Short: "Batch Parse",
		Run: func(cmd *cobra.Command, args []string) {
			batch(c, cmd, args)
		},
	}
	cmd.Flags().String("tracesPath", ".", "--tracesPath .")
	cmd.Flags().Bool("useDataStore", false, "--useDataStore true")
	cmd.Flags().String("inFile", "", "--inFile ./fallback.json")
	cmd.Flags().String("outFile", "./unknown_methods.json", "--outFile ./unknown_methods.json")
	cmd.Flags().Bool("parseSubTxs", false, "--parseSubTxs true")
	return cmd
}

func get(c *cli.CLI, cmd *cobra.Command, _ []string) {
	logger := logger2.GetSafeLogger(nil)
	logger.Infof(c.GetVersionString())

	config, err := cli.LoadConfig[Config]()
	if err != nil {
		logger.Errorf("Error loading config: %s", err)
		return
	}
	logType, err := cmd.Flags().GetString("type")
	if err != nil {
		logger.Errorf("Error loading type: %s", err)
		return
	}
	outPath, err := cmd.Flags().GetString("outPath")
	if err != nil {
		logger.Errorf("Error loading outPath: %s", err)
		return
	}
	format, err := cmd.Flags().GetString("compress")
	if err != nil {
		logger.Errorf("Error loading compress: %s", err)
		return
	}
	height, err := cmd.Flags().GetUint64("height")
	if err != nil {
		logger.Errorf("Error loading height: %s", err)
		return
	}

	rpcClient, err := newFilecoinRPCClient(config.NodeURL, config.NodeToken)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	var data any
	switch logType {
	case "traces":
		_, data, err = getTraceFileByHeight(height, rpcClient.client)
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
		logger.Error(err.Error())
		return
	}

	dataJson, err := sonic.Marshal(data)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	out := dataJson
	fname := fmt.Sprintf("%s_%d.json", logType, height)
	if format != "" {
		out, err = compress(format, dataJson)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		fname = fmt.Sprintf("%s_%d.json.%s", logType, height, format)
	}

	if err := writeToFile(outPath, fname, out); err != nil {
		logger.Error(err.Error())
		return
	}

}

func GetUploadCommand(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload",
		Run: func(cmd *cobra.Command, args []string) {
			upload(c, cmd, args)
		},
	}
	cmd.Flags().String("traceFile", "", "--traceFile ./tmp.txt")
	cmd.Flags().String("outPath", "", "--outPath /tmp/")

	return cmd
}
