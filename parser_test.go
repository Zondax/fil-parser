package fil_parser

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"github.com/zondax/fil-parser/actors/database"
	"net/http"
	"os"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/stretchr/testify/require"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	"github.com/zondax/fil-parser/types"
)

const (
	dataPath          = "data/heights"
	fileDataExtension = "json.gz"
	tracesPrefix      = "traces"
	tipsetPrefix      = "tipset"
	ethLogPrefix      = "ethlog"
	nodeUrl           = "https://api.zondax.ch/fil/node/mainnet/rpc/v1"
)

func getFilename(prefix, height string) string {
	return fmt.Sprintf(`%s/%s_%s.%s`, dataPath, prefix, height, fileDataExtension)
}

func tracesFilename(height string) string {
	return getFilename(tracesPrefix, height)
}

func ehtlogFilename(height string) string {
	return getFilename(ethLogPrefix, height)
}

func tipsetFilename(height string) string {
	return getFilename(tipsetPrefix, height)
}

func readGzFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(gzipReader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func readTipset(height string) (*types.ExtendedTipSet, error) {
	rawTipset, err := readGzFile(tipsetFilename(height))
	if err != nil {
		return nil, err
	}
	var tipset *types.ExtendedTipSet
	err = sonic.Unmarshal(rawTipset, &tipset)
	if err != nil {
		return nil, err
	}

	return tipset, nil
}

func readEthLogs(height string) ([]types.EthLog, error) {
	rawLogs, err := readGzFile(ehtlogFilename(height))
	if err != nil {
		return nil, err
	}
	var logs []types.EthLog
	err = sonic.Unmarshal(rawLogs, &logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func TestParser_ParseTransactions(t *testing.T) {
	// expectedResults are from previous runs. This assures backward compatibility. (Worst case would be fewer traces
	// or address than previous versions)
	type expectedResults struct {
		totalTraces  int
		totalAddress int
	}
	tests := []struct {
		name    string
		version string
		url     string
		height  string
		results expectedResults
	}{
		{
			name:    "v22 parser with traces from v22",
			version: "v22",
			url:     nodeUrl,
			height:  "2907480",
			results: expectedResults{
				totalTraces:  652,
				totalAddress: 95,
			},
		},
		{
			name:    "v23 parser with traces from v22",
			version: "v23",
			url:     nodeUrl,
			height:  "2907480",
			results: expectedResults{
				totalTraces:  652,
				totalAddress: 95,
			},
		},
		{
			name:    "v22 parser with traces from v23",
			version: "v22",
			url:     nodeUrl,
			height:  "2907520",
			results: expectedResults{
				totalTraces:  907,
				totalAddress: 85,
			},
		},
		{
			name:    "v23 parser with traces from v23",
			version: "v23",
			url:     nodeUrl,
			height:  "2907520",
			results: expectedResults{
				totalTraces:  907,
				totalAddress: 85,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lotusClient, _, err := client.NewFullNodeRPCV1(context.Background(), tt.url, http.Header{})
			require.NoError(t, err)
			require.NotNil(t, lotusClient)

			database.SetupActorsDatabase(&lotusClient)

			lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
			require.NotNil(t, lib)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			p, err := NewParser(lib, tt.version)
			require.NoError(t, err)
			txs, adds, err := p.ParseTransactions(traces, &tipset.TipSet, ethlogs)
			require.NoError(t, err)
			require.NotNil(t, txs)
			require.NotNil(t, adds)
			require.Equal(t, tt.results.totalTraces, len(txs))
			require.Equal(t, tt.results.totalAddress, len(adds))
		})
	}
}

func TestParser_InDepthCompare(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		height string
	}{
		{
			name:   "height downloaded with v22",
			url:    nodeUrl,
			height: "2907480",
		},
		{
			name:   "height downloaded with v23",
			url:    nodeUrl,
			height: "2907520",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lotusClient, _, err := client.NewFullNodeRPCV1(context.Background(), tt.url, http.Header{})
			require.NoError(t, err)
			require.NotNil(t, lotusClient)

			database.SetupActorsDatabase(&lotusClient)

			lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
			require.NotNil(t, lib)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			v22P, err := NewParser(lib, "v22")
			require.NoError(t, err)
			v22Txs, v22Adds, err := v22P.ParseTransactions(traces, &tipset.TipSet, ethlogs)
			require.NoError(t, err)
			require.NotNil(t, v22Txs)
			require.NotNil(t, v22Adds)

			v23P, err := NewParser(lib, "v23")
			require.NoError(t, err)
			v23Txs, v23Adds, err := v23P.ParseTransactions(traces, &tipset.TipSet, ethlogs)
			require.NoError(t, err)
			require.NotNil(t, v23Txs)
			require.NotNil(t, v23Adds)

			require.Equal(t, len(v22Txs), len(v23Txs))
			require.Equal(t, len(v22Adds), len(v23Adds))

			for i := range v22Txs {
				require.Equal(t, v22Txs[i], v23Txs[i])
			}
			for k := range v22Adds {
				require.Equal(t, v22Adds[k], v23Adds[k])
			}
			// todo
			require.Equal(t, v22Txs, v23Txs)
		})
	}
}
