package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/bytedance/sonic"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/klauspost/compress/s2"
	"github.com/spf13/cobra"
	"github.com/zondax/fil-parser/actors/account"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/golem/pkg/cli"
	"go.uber.org/zap"
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

	cmd.Flags().UintSlice("heights", []uint{0}, "--heights 387926,387927,387928")
	return cmd
}

func getTrace(height uint64, logType string, rpcClient *RPCClient) ([]byte, error) {
	var data any
	var err error
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
		return nil, err
	}

	dataJson, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataJson, nil
}

func get(c *cli.CLI, cmd *cobra.Command, _ []string) {
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
	if len(heights) > 0 {
		for _, tmp := range heights {
			height := uint64(tmp)
			dataJson, err := getTrace(height, logType, rpcClient)
			if err != nil {
				zap.S().Error(err)
				return
			}
			actorJson, err := DownloadActorJSON(height, dataJson)
			if err != nil {
				zap.S().Error(err)
				continue
			}
			fmt.Println(string(actorJson))

			out := dataJson
			fname := fmt.Sprintf("%s_%d.json", logType, height)
			if format != "" {
				out, err = compress(format, dataJson)
				if err != nil {
					zap.S().Error(err)
					return
				}
				fname = fmt.Sprintf("%s_%d.json.%s", logType, height, format)
			}

			if err := writeToFile(outPath, fname, out); err != nil {
				zap.S().Error(err)
				return
			}

		}
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

func DownloadActorJSON(height uint64, traces []byte) ([]byte, error) {
	var computeState *typesV2.ComputeStateOutputV2
	err := sonic.Unmarshal(traces, &computeState)
	if err != nil {
		return nil, err
	}
	if len(computeState.Trace) == 0 {
		return nil, fmt.Errorf("no traces found")
	}
	resp, err := account.AuthenticateMessage("mainnet", 10000, computeState.Trace[0].Msg.Params, computeState.Trace[0].MsgRct.Return)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	if string(data) == "{}" {
		return nil, fmt.Errorf("no data found")
	}
	return data, nil
}
