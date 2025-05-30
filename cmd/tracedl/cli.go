package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	logger2 "github.com/zondax/fil-parser/logger"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/bytedance/sonic"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/klauspost/compress/s2"
	"github.com/spf13/cobra"
	"github.com/zondax/golem/pkg/cli"
)

func GetStartCommand(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get",
		Run: func(cmd *cobra.Command, args []string) {
			get(c, cmd, args)
		},
	}
	cmd.Flags().String("type", "traces", "--type trace")
	cmd.Flags().String("outPath", ".", "--outPath ../")
	cmd.Flags().String("compress", "gz", "--compress s2")
	cmd.Flags().Uint64("height", 0, "--height 387926")
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
		data, err = getTraceFileByHeight(height, rpcClient.client)
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

func writeToFile(path, filename string, data []byte) error {
	tmp, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("%s/%s", tmp, filename), data, fs.ModePerm)
}

func compress(format string, data []byte) ([]byte, error) {
	// Compress data using s2
	var b bytes.Buffer
	dataBuff := bytes.NewBuffer(data)

	var enc io.WriteCloser
	switch format {
	case "s2":
		enc = s2.NewWriter(&b)
	case "gz":
		enc = gzip.NewWriter(&b)
	default:
		return nil, fmt.Errorf("invalid format,expected s2 or gz")
	}

	_, err := io.Copy(enc, dataBuff)
	if err != nil {
		_ = enc.Close()
		return nil, err
	}
	// Blocks until compression is done.
	_ = enc.Close()

	return b.Bytes(), nil
}
