package fil_parser

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/stretchr/testify/assert"

	"github.com/bytedance/sonic"
	"github.com/filecoin-project/go-address"
	filBig "github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	cidLink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/parser"
	v1 "github.com/zondax/fil-parser/parser/v1"
	v2 "github.com/zondax/fil-parser/parser/v2"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/golem/pkg/logger"
	"github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/zcache"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	"github.com/zondax/fil-parser/types"
)

const (
	dataPath          = "data/heights"
	fileDataExtension = "json.gz"
	tracesPrefix      = "traces"
	tipsetPrefix      = "tipset"
	ethLogPrefix      = "ethlog"
	nativeLogPrefix   = "nativelog"
	nodeUrl           = "https://hel1-node-fil-mainnet-light.zondax.ch/rpc/v1"
	calibNextNodeUrl  = "https://node-fil-calibration-stable.zondax.ch/rpc/v1"
	feeType           = "fee"
)

var gLogger = logger.NewLogger(logger.Config{
	Level:    "error",
	Encoding: "text",
})

// var nodeUrlParserV1 *FilecoinParser
// var calibNextNodeParserV1 *FilecoinParser
// var nodeUrlParserV2 *FilecoinParser
// var calibNextNodeParserV2 *FilecoinParser

var mainnetCacheDataSource = getCacheDataSource("mainnet", nodeUrl)
var calibNextNodeCacheDataSource = getCacheDataSource("calibration", calibNextNodeUrl)

func getFilename(prefix, height string) string {
	return fmt.Sprintf(`%s/%s_%s.%s`, dataPath, prefix, height, fileDataExtension)
}

func tracesFilename(height string) string {
	return getFilename(tracesPrefix, height)
}

func ehtlogFilename(height string) string {
	return getFilename(ethLogPrefix, height)
}

func nativeLogFilename(height string) string {
	return getFilename(nativeLogPrefix, height)
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

func readNativeLogs(height string) ([]*filTypes.ActorEvent, error) {
	rawLogs, err := readGzFile(nativeLogFilename(height))
	if err != nil {
		return nil, err
	}
	var logs []*filTypes.ActorEvent
	err = sonic.Unmarshal(rawLogs, &logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func getLotusClient(url string) api.FullNode {
	lotusClient, _, err := client.NewFullNodeRPCV1(context.Background(), url, http.Header{})
	if err != nil {
		panic(err)
	}

	return lotusClient
}

func getLib(nodeURL string) *rosettaFilecoinLib.RosettaConstructionFilecoin {
	lotusClient := getLotusClient(nodeURL)

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
	if lib == nil {
		panic("Rosetta lib should not be nil")
	}
	return lib
}

func getCacheDataSource(networkName string, nodeURL string) common.DataSource {
	cacheTTL := -1

	return common.DataSource{
		Node: getLotusClient(nodeURL),
		Config: common.DataSourceConfig{
			NetworkName: networkName,
			Cache: &common.CacheConfig{
				CombinedConfig: &zcache.CombinedConfig{
					Local: &zcache.LocalConfig{},
					Remote: &zcache.RemoteConfig{
						Addr:     os.Getenv("REDIS_ADDR"),
						Password: os.Getenv("REDIS_SECRET"),
					},
					IsRemoteBestEffort: false,
					GlobalPrefix:       "ci",
					GlobalMetricServer: metrics.NewNoopMetrics(),
				},
				Ttl: time.Duration(cacheTTL) * time.Second,
			},
		},
	}
}

func TestParser_ParseTransactions(t *testing.T) {
	// expectedResults are from previous runs. This assures backward compatibility. (Worst case would be fewer traces
	// or address than previous versions)
	type expectedResults struct {
		totalTraces  int
		totalAddress int
		totalTxCids  int
	}
	tests := []struct {
		name    string
		version string
		url     string
		height  string
		results expectedResults
	}{
		// {
		// 	name:    "parser with traces from v1",
		// 	version: v1.NodeVersionsSupported[0],
		// 	url:     nodeUrl,
		// 	height:  "2907480",
		// 	results: expectedResults{
		// 		totalTraces:  650,
		// 		totalAddress: 232,
		// 		totalTxCids:  102,
		// 	},
		// },
		{
			name:    "parser with traces from v1 and the corner case of duplicated fees with level 0",
			version: v1.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "845259",
			results: expectedResults{
				totalTraces:  31,
				totalAddress: 12,
				totalTxCids:  10,
			},
		},
		// {
		// 	name:    "parser with traces from v2",
		// 	version: v2.NodeVersionsSupported[0],
		// 	url:     nodeUrl,
		// 	height:  "2907520",
		// 	results: expectedResults{
		// 		totalTraces:  910,
		// 		totalAddress: 234,
		// 		totalTxCids:  151,
		// 	},
		// },
		// {
		// 	name:    "parser with traces from v2 and lotus 1.25",
		// 	version: v2.NodeVersionsSupported[2],
		// 	url:     nodeUrl,
		// 	height:  "3573062",
		// 	results: expectedResults{
		// 		totalTraces:  774,
		// 		totalAddress: 209,
		// 		totalTxCids:  121,
		// 	},
		// },
		// {
		// 	name:    "parser with traces from v2 and lotus 1.25",
		// 	version: v2.NodeVersionsSupported[2],
		// 	url:     nodeUrl,
		// 	height:  "3573064",
		// 	results: expectedResults{
		// 		totalTraces:  734,
		// 		totalAddress: 206,
		// 		totalTxCids:  101,
		// 	},
		// },
		// {
		// 	name:    "parser with traces from v2 and lotus 1.25",
		// 	version: v2.NodeVersionsSupported[2],
		// 	url:     nodeUrl,
		// 	height:  "3573066",
		// 	results: expectedResults{
		// 		totalTraces:  1121,
		// 		totalAddress: 274,
		// 		totalTxCids:  187,
		// 	},
		// },
		// {
		// 	name:    "parser with traces from v2 and lotus 1.26 (calib)",
		// 	version: v2.NodeVersionsSupported[2],
		// 	url:     calibNextNodeUrl,
		// 	height:  "1419335",
		// 	results: expectedResults{
		// 		totalTraces:  37,
		// 		totalAddress: 18,
		// 		totalTxCids:  5,
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := getLib(tt.url)

			var p *FilecoinParser
			var err error

			if tt.url == nodeUrl {
				p, err = NewFilecoinParser(l, mainnetCacheDataSource, gLogger)
			} else {
				p, err = NewFilecoinParser(l, calibNextNodeCacheDataSource, gLogger)
			}
			require.NoError(t, err)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			txsData := types.TxsData{
				EthLogs:  ethlogs,
				Tipset:   tipset,
				Traces:   traces,
				Metadata: types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: tt.version}},
			}

			parsedResult, err := p.ParseTransactions(context.Background(), txsData)
			require.NoError(t, err)
			require.NotNil(t, parsedResult.Txs)
			require.NotNil(t, parsedResult.Addresses)
			require.Equal(t, tt.results.totalTraces, len(parsedResult.Txs))
			require.Equal(t, tt.results.totalAddress, parsedResult.Addresses.Len())
			require.Equal(t, tt.results.totalTxCids, len(parsedResult.TxCids))
		})
	}
}

func TestParser_GetBaseFee(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		url      string
		height   string
		baseFee  *big.Int
		fallback bool
	}{
		{
			name:    "parser with getBaseFee",
			version: v2.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "2907480",
			baseFee: big.NewInt(96036633),
		},
		{
			name:     "parser with getBaseFee fallback",
			version:  v2.NodeVersionsSupported[0],
			url:      nodeUrl,
			height:   "3450305",
			baseFee:  big.NewInt(100),
			fallback: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p *FilecoinParser
			var err error
			if tt.url == nodeUrl {
				p, err = NewFilecoinParser(getLib(tt.url), mainnetCacheDataSource, gLogger)
			} else {
				p, err = NewFilecoinParser(getLib(tt.url), calibNextNodeCacheDataSource, gLogger)
			}
			require.NoError(t, err)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			baseFee, err := p.GetBaseFee(traces, types.BlockMetadata{}, tipset)
			require.NoError(t, err)
			require.Equal(t, baseFee, tt.baseFee.Uint64())
			if tt.fallback {
				require.Equal(t, baseFee, tipset.Blocks()[0].ParentBaseFee.Uint64())
			}
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
			name:   "height downloaded with v1",
			url:    nodeUrl,
			height: "2907480",
		},
		{
			name:   "height downloaded with v2",
			url:    nodeUrl,
			height: "2907520",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p1 *FilecoinParser
			var p2 *FilecoinParser
			var err1 error
			var err2 error
			if tt.url == nodeUrl {
				p1, err1 = NewFilecoinParser(getLib(tt.url), mainnetCacheDataSource, gLogger)
				p2, err2 = NewFilecoinParser(getLib(tt.url), mainnetCacheDataSource, gLogger)
			} else {
				p1, err1 = NewFilecoinParser(getLib(tt.url), calibNextNodeCacheDataSource, gLogger)
				p2, err2 = NewFilecoinParser(getLib(tt.url), calibNextNodeCacheDataSource, gLogger)
			}
			require.NoError(t, err1)
			require.NoError(t, err2)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			wg := sync.WaitGroup{}
			wg.Add(2)
			var parsedResultV1 *types.TxsParsedResult
			var parsedResultV2 *types.TxsParsedResult
			go func() {
				txsData1 := types.TxsData{
					EthLogs:  ethlogs,
					Tipset:   tipset,
					Traces:   traces,
					Metadata: types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: "v1.22"}},
				}
				defer wg.Done()
				parsedResultV1, err1 = p1.ParseTransactions(context.Background(), txsData1)
			}()
			go func() {
				defer wg.Done()
				txsData2 := types.TxsData{
					EthLogs:  ethlogs,
					Tipset:   tipset,
					Traces:   traces,
					Metadata: types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: "v1.23"}},
				}
				parsedResultV2, err2 = p2.ParseTransactions(context.Background(), txsData2)
			}()

			wg.Wait()

			require.NoError(t, err1)
			require.NoError(t, err2)

			require.NotNil(t, parsedResultV1.Txs)
			require.NotNil(t, parsedResultV1.Addresses)
			require.NotNil(t, parsedResultV2.Txs)
			require.NotNil(t, parsedResultV2.Addresses)

			require.Equal(t, len(parsedResultV1.Txs), len(parsedResultV2.Txs))
			require.Equal(t, parsedResultV1.Addresses.Len(), parsedResultV2.Addresses.Len())
			require.Equal(t, len(parsedResultV1.TxCids), len(parsedResultV2.TxCids))

			for i := range parsedResultV1.Txs {
				tmp1, _ := json.Marshal(parsedResultV1.Txs[i])
				tmp2, _ := json.Marshal(parsedResultV2.Txs[i])
				if parsedResultV1.Txs[i].TxType == parser.TotalFeeOp {
					parsedResultV1.Txs[i].TxTo = parser.BurnAddress
					parsedResultV2.Txs[i].TxTo = parser.BurnAddress
				}

				if strings.EqualFold(parsedResultV1.Txs[i].TxType, parser.MethodUnknown) && !strings.EqualFold(parsedResultV2.Txs[i].TxType, parser.MethodUnknown) {
					// v2 fixed the unknown method
					continue
				}

				require.Truef(t, parsedResultV1.Txs[i].Equal(*parsedResultV2.Txs[i]), "tx %d is not equal\n%s\n%s", i, tmp1, tmp2)
			}

			for i := range parsedResultV1.TxCids {
				tmp1, _ := json.Marshal(parsedResultV1.TxCids[i])
				tmp2, _ := json.Marshal(parsedResultV2.TxCids[i])
				require.Truef(t, reflect.DeepEqual(parsedResultV1.TxCids[i], parsedResultV2.TxCids[i]), "tx cid %d is not equal\n%s\n%s", i, tmp1, tmp2)
			}

			parsedResultV1.Addresses.Range(func(key string, value *types.AddressInfo) bool {
				v2Value, ok := parsedResultV2.Addresses.Get(key)
				require.Truef(t, ok, "address %s is not equal", key)
				require.Equal(t, value, v2Value)
				return true
			})
		})
	}
}

func TestParser_ParseEvents_EVM_FromTraceFile(t *testing.T) {
	// expectedResults are from previous runs. This assures backward compatibility. (Worst case would be fewer traces
	// or address than previous versions)
	type expectedResults struct {
		totalTraces       int
		totalNativeEvents int
		totalEVMEvents    int
	}
	tests := []struct {
		name    string
		version string
		url     string
		height  string
		results expectedResults
	}{
		{
			name:    "parser with traces from v1",
			version: v1.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "2907480",
			results: expectedResults{
				totalTraces:       104,
				totalNativeEvents: 0,
				totalEVMEvents:    104,
			},
		},
		{
			name:    "parser with traces from v2",
			version: v2.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "2907520",
			results: expectedResults{
				totalTraces:       11,
				totalNativeEvents: 0,
				totalEVMEvents:    11,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3573062",
			results: expectedResults{
				totalTraces:       1,
				totalNativeEvents: 0,
				totalEVMEvents:    1,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3573064",
			results: expectedResults{
				totalTraces:       4,
				totalNativeEvents: 0,
				totalEVMEvents:    4,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3573066",
			results: expectedResults{
				totalTraces:       1,
				totalNativeEvents: 0,
				totalEVMEvents:    1,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.26 (calib)",
			version: v2.NodeVersionsSupported[2],
			url:     calibNextNodeUrl,
			height:  "1419335",
			results: expectedResults{
				totalTraces:       2,
				totalNativeEvents: 0,
				totalEVMEvents:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p *FilecoinParser
			var err error
			if tt.url == nodeUrl {
				p, err = NewFilecoinParser(getLib(tt.url), mainnetCacheDataSource, gLogger)
			} else {
				p, err = NewFilecoinParser(getLib(tt.url), calibNextNodeCacheDataSource, gLogger)
			}
			require.NoError(t, err)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)

			eventsData := types.EventsData{
				EthLogs:  ethlogs,
				Tipset:   tipset,
				Metadata: types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: tt.version}},
			}

			parsedResult, err := p.ParseEthLogs(context.Background(), eventsData)
			require.NoError(t, err)
			require.NotNil(t, parsedResult.ParsedEvents)
			require.Equal(t, tt.results.totalTraces, len(parsedResult.ParsedEvents))
			require.Equal(t, tt.results.totalNativeEvents, parsedResult.NativeEvents)
			require.Equal(t, tt.results.totalEVMEvents, parsedResult.EVMEvents)
		})
	}
}
func TestParser_ParseEvents_FVM_FromTraceFile(t *testing.T) {
	// expectedResults are from previous runs. This assures backward compatibility. (Worst case would be fewer traces
	// or address than previous versions)
	type expectedResults struct {
		totalTraces       int
		totalNativeEvents int
		totalEVMEvents    int
	}
	tests := []struct {
		name    string
		version string
		url     string
		height  string
		results expectedResults
	}{
		{
			name:    "parser with traces from v1",
			version: v1.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "3897964",
			results: expectedResults{
				totalTraces:       808,
				totalNativeEvents: 808,
				totalEVMEvents:    0,
			},
		},
		{
			name:    "parser with traces from v2",
			version: v2.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "3897964",
			results: expectedResults{
				totalTraces:       808,
				totalNativeEvents: 808,
				totalEVMEvents:    0,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3897964",
			results: expectedResults{
				totalTraces:       808,
				totalNativeEvents: 808,
				totalEVMEvents:    0,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3897964",
			results: expectedResults{
				totalTraces:       808,
				totalNativeEvents: 808,
				totalEVMEvents:    0,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3897964",
			results: expectedResults{
				totalTraces:       808,
				totalNativeEvents: 808,
				totalEVMEvents:    0,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.26 (calib)",
			version: v2.NodeVersionsSupported[2],
			url:     calibNextNodeUrl,
			height:  "3897964",
			results: expectedResults{
				totalTraces:       808,
				totalNativeEvents: 808,
				totalEVMEvents:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p *FilecoinParser
			var err error
			if tt.url == nodeUrl {
				p, err = NewFilecoinParser(getLib(tt.url), mainnetCacheDataSource, gLogger)
			} else {
				p, err = NewFilecoinParser(getLib(tt.url), calibNextNodeCacheDataSource, gLogger)
			}
			require.NoError(t, err)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			nativeLogs, err := readNativeLogs(tt.height)
			require.NoError(t, err)

			eventsData := types.EventsData{
				NativeLog: nativeLogs,
				Tipset:    tipset,
				Metadata:  types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: tt.version}},
			}

			parsedResult, err := p.ParseNativeEvents(context.Background(), eventsData)
			require.NoError(t, err)
			require.NotNil(t, parsedResult.ParsedEvents)
			require.Equal(t, tt.results.totalTraces, len(parsedResult.ParsedEvents))
			require.Equal(t, tt.results.totalNativeEvents, parsedResult.NativeEvents)
			require.Equal(t, tt.results.totalEVMEvents, parsedResult.EVMEvents)
		})
	}
}

func buildCidLink(cid cid.Cid) datamodel.Link {
	return cidLink.Link{Cid: cid}
}

func ipldEncode(t *testing.T, builder datamodel.NodeBuilder, data any) []byte {
	var err error

	switch x := data.(type) {
	case string:
		err = builder.AssignString(x)
	case []byte:
		err = builder.AssignBytes(x)
	case datamodel.Link:
		err = builder.AssignLink(x)
	case int64:
		err = builder.AssignInt(x)
	}

	require.NoError(t, err)
	encoded, err := ipld.Encode(builder.Build(), dagcbor.Encode)
	require.NoError(t, err)
	return encoded
}

func TestParser_ParseNativeEvents_FVM(t *testing.T) {
	// we need any random number for the test
	//nolint:gosec
	filAddress, err := address.NewIDAddress(uint64(rand.Int()))
	assert.NoError(t, err)

	tipset := &types.ExtendedTipSet{
		TipSet:        filTypes.TipSet{},
		BlockMessages: nil,
	}

	parser, err := NewFilecoinParser(getLib(calibNextNodeUrl), calibNextNodeCacheDataSource, gLogger)
	require.NoError(t, err)

	eventType := ipldEncode(t, basicnode.Prototype.String.NewBuilder(), "market_deals_event")
	eventData := ipldEncode(t, basicnode.Prototype.Bytes.NewBuilder(), []byte("test_data"))

	// cid event data
	eventCid, err := cid.Decode("baga6ea4seaqeyz6zikyr2bqbhy6mrocoqwagx45vlbpsbem7euqv5mf3hrvn2fy")
	require.NoError(t, err)
	link := buildCidLink(eventCid)
	cidEventType := ipldEncode(t, basicnode.Prototype.String.NewBuilder(), "sector_activated")
	cidEventData := ipldEncode(t, basicnode.Prototype.Link.NewBuilder(), link)

	// nullable cid event data
	nullableCidEventType := ipldEncode(t, basicnode.Prototype.String.NewBuilder(), "sector_activated")
	b := basicnode.Prototype__Any{}.NewBuilder()
	err = b.AssignNull()
	require.NoError(t, err)
	nullableCidEventData, err := ipld.Encode(b.Build(), dagcbor.Encode)
	require.NoError(t, err)

	// bigInt event data
	bigInt, err := filBig.FromString("12345678901234567891234567890123456789012345678901234567890")
	require.NoError(t, err)
	bigIntEventType := ipldEncode(t, basicnode.Prototype.String.NewBuilder(), "verifier_balance")
	tmp, err := bigInt.Bytes()
	require.NoError(t, err)
	bigIntEventData := ipldEncode(t, basicnode.Prototype.Bytes.NewBuilder(), tmp)

	largeInt := math.MaxInt64
	largeIntEventData := ipldEncode(t, basicnode.Prototype.Int.NewBuilder(), int64(largeInt))

	smallInt := 10
	smallIntEventData := ipldEncode(t, basicnode.Prototype.Int.NewBuilder(), int64(smallInt))

	negativeInt := -10
	negativeIntEventData := ipldEncode(t, basicnode.Prototype.Int.NewBuilder(), int64(negativeInt))

	veryNegativeInt := math.MinInt64
	veryNegativeIntEventData := ipldEncode(t, basicnode.Prototype.Int.NewBuilder(), int64(veryNegativeInt))

	tb := []struct {
		name         string
		entries      []filTypes.EventEntry
		emitter      address.Address
		wantMetadata map[int]map[string]any
		wantErr      bool
	}{
		{
			name:    "error ipld decode",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: []byte("invalid"), // invalid format for provided codec
				},
				{
					Flags: 0x03,
					Key:   "data",
					Codec: 0x51,
					Value: eventData,
				},
			},
			wantErr: true,
		},
		{
			name:    "error retreiving $type from entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x52, // wrong codec
					Value: []byte("invalid"),
				},
				{
					Flags: 0x03,
					Key:   "data",
					Codec: 0x51,
					Value: eventData,
				},
			},
			wantErr: true,
		},
		{
			name:    "success no entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{},
		},
		{
			name:    "succes native fvm events",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: eventType,
				},
				{
					Flags: 0x03,
					Key:   "data",
					Codec: 0x51,
					Value: eventData,
				},
			},
			wantMetadata: map[int]map[string]any{
				0: {
					"flags": 3,
					"key":   "$type",
					"value": "market_deals_event",
				},
				1: {
					"flags": 3,
					"key":   "data",
					"value": "dGVzdF9kYXRh",
				},
			},
		},
		{
			name:    "success native negative int event entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: eventType,
				},
				{
					Flags: 0x03,
					Key:   "expiry",
					Codec: 0x51,
					Value: negativeIntEventData,
				},
			},
			wantMetadata: map[int]map[string]any{
				0: {
					"flags": 3,
					"key":   "$type",
					"value": "market_deals_event",
				},
				1: {
					"flags": 3,
					"key":   "expiry",
					"value": negativeInt,
				},
			},
		},
		{
			name:    "success native very negative int event entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: eventType,
				},
				{
					Flags: 0x03,
					Key:   "expiry",
					Codec: 0x51,
					Value: veryNegativeIntEventData,
				},
			},
			wantMetadata: map[int]map[string]any{
				0: {
					"flags": 3,
					"key":   "$type",
					"value": "market_deals_event",
				},
				1: {
					"flags": 3,
					"key":   "expiry",
					"value": fmt.Sprint(veryNegativeInt),
				},
			},
		},
		{
			name:    "success native small int event entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: eventType,
				},
				{
					Flags: 0x03,
					Key:   "expiry",
					Codec: 0x51,
					Value: smallIntEventData,
				},
			},
			wantMetadata: map[int]map[string]any{
				0: {
					"flags": 3,
					"key":   "$type",
					"value": "market_deals_event",
				},
				1: {
					"flags": 3,
					"key":   "expiry",
					"value": smallInt,
				},
			},
		},
		{
			name:    "success native large int event entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: eventType,
				},
				{
					Flags: 0x03,
					Key:   "expiry",
					Codec: 0x51,
					Value: largeIntEventData,
				},
			},
			wantMetadata: map[int]map[string]any{
				0: {
					"flags": 3,
					"key":   "$type",
					"value": "market_deals_event",
				},
				1: {
					"flags": 3,
					"key":   "expiry",
					"value": fmt.Sprint(largeInt),
				},
			},
		},
		{
			name:    "success native bigInt event entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: bigIntEventType,
				},
				{
					Flags: 0x03,
					Key:   "balance",
					Codec: 0x51,
					Value: bigIntEventData,
				},
			},
			wantMetadata: map[int]map[string]any{
				0: {
					"flags": 3,
					"key":   "$type",
					"value": "verifier_balance",
				},
				1: {
					"flags": 3,
					"key":   "balance",
					"value": "12345678901234567891234567890123456789012345678901234567890",
				},
			},
		},
		{
			name:    "succes native cid event entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: cidEventType,
				},
				{
					Flags: 0x03,
					Key:   "piece_cid",
					Codec: 0x51,
					Value: cidEventData,
				},
			},
			wantMetadata: map[int]map[string]any{
				0: {
					"flags": 3,
					"key":   "$type",
					"value": "sector_activated",
				},
				1: {
					"flags": 3,
					"key":   "piece_cid",
					"value": map[string]any{
						"/": "baga6ea4seaqeyz6zikyr2bqbhy6mrocoqwagx45vlbpsbem7euqv5mf3hrvn2fy",
					},
				},
			},
		},
		{
			name:    "succes native nullable cid event entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: nullableCidEventType,
				},
				{
					Flags: 0x03,
					Key:   "unsealed_cid",
					Codec: 0x51,
					Value: nullableCidEventData,
				},
			},
			wantMetadata: map[int]map[string]any{
				0: {
					"flags": 3,
					"key":   "$type",
					"value": "sector_activated",
				},
				1: {
					"flags": 3,
					"key":   "unsealed_cid",
					"value": nil,
				},
			},
		},
		{
			name:    "succes native nullable cid and valid cid event entries",
			emitter: filAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "$type",
					Codec: 0x51,
					Value: nullableCidEventType,
				},
				{
					Flags: 0x03,
					Key:   "unsealed_cid",
					Codec: 0x51,
					Value: nullableCidEventData,
				},
				{
					Flags: 0x03,
					Key:   "piece_cid",
					Codec: 0x51,
					Value: cidEventData,
				},
			},
			wantMetadata: map[int]map[string]any{
				0: {
					"flags": 3,
					"key":   "$type",
					"value": "sector_activated",
				},
				1: {
					"flags": 3,
					"key":   "unsealed_cid",
					"value": nil,
				},
				2: {
					"flags": 3,
					"key":   "piece_cid",
					"value": map[string]any{
						"/": "baga6ea4seaqeyz6zikyr2bqbhy6mrocoqwagx45vlbpsbem7euqv5mf3hrvn2fy",
					},
				},
			},
		},
	}

	for i := range tb {
		tt := tb[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			eventsData := types.EventsData{
				Tipset: tipset,
				NativeLog: []*filTypes.ActorEvent{
					{
						Emitter: tt.emitter,
						Entries: tt.entries,
					},
				},
			}
			events, err := parser.ParseNativeEvents(ctx, eventsData)
			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, events)
			require.NotEmpty(t, events.ParsedEvents)

			gotMetadata := map[int]map[string]any{}
			err = json.Unmarshal([]byte(events.ParsedEvents[0].Metadata), &gotMetadata)
			assert.NoError(t, err)

			for idx, v := range tt.wantMetadata {
				for entryKey, entryValue := range v {
					assert.EqualValues(t, entryValue, gotMetadata[idx][entryKey])
				}
			}
			assert.EqualValues(t, tipset.GetCidString(), events.ParsedEvents[0].TipsetCid)
			assert.EqualValues(t, tt.emitter.String(), events.ParsedEvents[0].Emitter)
			if len(tt.entries) > 0 { // only check for the selector_id if we have entries in the test case\
				assert.Regexp(t, "market_deals_event|sector_activated|verifier_balance", events.ParsedEvents[0].SelectorID)
			}

			// check if IDs are unique for all events
			foundIDs := map[string]bool{}
			for idx, evt := range events.ParsedEvents {
				wantID := tools.BuildId(tipset.GetCidString(), cid.Cid{}.String(), fmt.Sprint(idx), types.EventTypeNative)
				gotID := evt.ID

				assert.EqualValues(t, wantID, gotID)
				assert.NotContains(t, foundIDs, gotID)
				foundIDs[gotID] = true
			}
		})
	}

}
func TestParser_ParseNativeEvents_EVM(t *testing.T) {
	ethAddress, err := address.NewDelegatedAddress(32, []byte{})
	assert.NoError(t, err)

	tipset := &types.ExtendedTipSet{
		TipSet:        filTypes.TipSet{},
		BlockMessages: nil,
	}

	var topic ethtypes.EthHash
	err = topic.UnmarshalJSON([]byte(`"0x013dbb9442ca9667baccc6230fcd5c1c4b2d4d2870f4bd20681d4d47cfd15184"`))
	assert.NoError(t, err)

	topicBytes := make([]byte, ethtypes.EthHashLength)
	n := copy(topicBytes, topic[:ethtypes.EthHashLength])
	assert.Equal(t, ethtypes.EthHashLength, n)

	eventData, err := json.Marshal(map[string]any{
		"x": "y",
		"a": "b",
	})
	assert.NoError(t, err)
	eventDataHex := hex.EncodeToString(eventData)

	parser, err := NewFilecoinParser(getLib(calibNextNodeUrl), calibNextNodeCacheDataSource, gLogger)
	require.NoError(t, err)

	tb := []struct {
		name         string
		entries      []filTypes.EventEntry
		emitter      address.Address
		wantMetadata map[string]any
		wantErr      bool
	}{
		{
			name:    "error retrieving topic from entry",
			emitter: ethAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "t1",
					Codec: 0x52, // wrong codec
					Value: []byte{},
				},
			},
			wantErr: true,
		},
		{
			name:    "error parsing ethHash",
			emitter: ethAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "t1",
					Codec: 0x55,
					Value: []byte{}, // empty hash
				},
			},
			wantErr: true,
		},
		{
			name:    "succes native evm events no entries",
			emitter: ethAddress,
			entries: []filTypes.EventEntry{},
			wantMetadata: map[string]any{
				"topics": []string{},
				"data":   "",
			},
		},
		{
			name:    "succes native evm events",
			emitter: ethAddress,
			entries: []filTypes.EventEntry{
				{
					Flags: 0x03,
					Key:   "t1",
					Codec: 0x55,
					Value: createTopic(t, "0x013dbb9442ca9667baccc6230fcd5c1c4b2d4d2870f4bd20681d4d47cfd15184"),
				},
				{
					Flags: 0x03,
					Key:   "t2",
					Codec: 0x55,
					Value: createTopic(t, "0xab8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
				},
				{
					Flags: 0x03,
					Key:   "t3",
					Codec: 0x55,
					Value: createTopic(t, "0xbb8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
				},
				{
					Flags: 0x03,
					Key:   "d",
					Codec: 0x55,
					Value: eventData,
				},
			},
			wantMetadata: map[string]any{
				"data": eventDataHex,
				"topics": []string{
					"0x013dbb9442ca9667baccc6230fcd5c1c4b2d4d2870f4bd20681d4d47cfd15184",
					"0xab8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738",
					"0xbb8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738",
				},
			},
		},
	}

	for i := range tb {
		tt := tb[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			eventsData := types.EventsData{
				Tipset: tipset,
				NativeLog: []*filTypes.ActorEvent{
					{
						Emitter: tt.emitter,
						Entries: tt.entries,
					},
					{
						Emitter: tt.emitter,
						Entries: tt.entries,
					},
				},
			}

			events, err := parser.ParseNativeEvents(ctx, eventsData)
			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, events)
			require.NotEmpty(t, events.ParsedEvents)

			gotMetadata := map[string]any{}
			err = json.Unmarshal([]byte(events.ParsedEvents[0].Metadata), &gotMetadata)
			assert.NoError(t, err)

			assert.EqualValues(t, tt.wantMetadata["data"], gotMetadata["data"])
			assert.ElementsMatch(t, tt.wantMetadata["topics"], gotMetadata["topics"])

			// check if IDs are unique for all events
			foundIDs := map[string]bool{}
			for idx, evt := range events.ParsedEvents {
				wantID := tools.BuildId(tipset.GetCidString(), cid.Cid{}.String(), fmt.Sprint(idx), types.EventTypeEVM)
				gotID := evt.ID

				assert.EqualValues(t, wantID, gotID)
				assert.NotContains(t, foundIDs, gotID)
				foundIDs[gotID] = true
			}

			assert.EqualValues(t, tt.emitter.String(), events.ParsedEvents[0].Emitter)
			if len(tt.entries) > 0 { // only check the selector_id if there are entries in the test case
				assert.EqualValues(t, "0x013dbb9442ca9667baccc6230fcd5c1c4b2d4d2870f4bd20681d4d47cfd15184", events.ParsedEvents[0].SelectorID)
			}
		})
	}

}

func TestParser_ParseEthLogs(t *testing.T) {
	var emitter ethtypes.EthAddress
	err := emitter.UnmarshalJSON([]byte(`"0xd4c5fb16488Aa48081296299d54b0c648C9333dA"`))
	assert.NoError(t, err)

	txCID := cid.Cid{}.String()

	tipset := &types.ExtendedTipSet{
		TipSet:        filTypes.TipSet{},
		BlockMessages: nil,
	}

	eventData, err := json.Marshal(map[string]any{
		"x": "y",
		"a": "b",
	})
	assert.NoError(t, err)
	eventDataHex := hex.EncodeToString(eventData)

	parser, err := NewFilecoinParser(getLib(calibNextNodeUrl), calibNextNodeCacheDataSource, gLogger)
	require.NoError(t, err)

	tb := []struct {
		name         string
		ethLogs      []types.EthLog
		wantMetadata map[string]any
		wantErr      bool
		wantSig      bool
	}{
		{
			name:    "success when signature not found",
			wantSig: false,
			ethLogs: []types.EthLog{
				{
					TransactionCid: txCID,
					EthLog: ethtypes.EthLog{
						Address: emitter,
						Data:    eventData,
						Topics: []ethtypes.EthHash{
							createEthHash(t, "0x013dbb9442ca9667baccc6230fcd5c1c4b2d4d2870f4bd20681d4d47cfd15184"),
							createEthHash(t, "0xab8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
							createEthHash(t, "0xbb8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
						},
					},
				},
			},
			wantMetadata: map[string]any{
				"data": eventDataHex,
				"topics": []string{
					"0x013dbb9442ca9667baccc6230fcd5c1c4b2d4d2870f4bd20681d4d47cfd15184",
					"0xab8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738",
					"0xbb8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738",
				},
			},
		},
		{
			name: "success no topics",
			ethLogs: []types.EthLog{
				{
					TransactionCid: txCID,
					EthLog: ethtypes.EthLog{
						Address: emitter,
						Data:    eventData,
						Topics:  []ethtypes.EthHash{},
					},
				},
			},
			wantMetadata: map[string]any{
				"topics": []string{},
				"data":   eventDataHex,
			},
		},
		{
			name:    "success",
			wantSig: true,
			ethLogs: []types.EthLog{
				{
					TransactionCid: txCID,
					EthLog: ethtypes.EthLog{
						TransactionIndex: 0,
						LogIndex:         0,
						Address:          emitter,
						Data:             eventData,
						Topics: []ethtypes.EthHash{
							createEthHash(t, "0x25eaabaf991947ec22f473a02c14ffbcc08ffe2cef8d81ac12b6db2c14ce23a0"),
							createEthHash(t, "0xab8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
							createEthHash(t, "0xbb8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
						},
					},
				},
				{
					TransactionCid: txCID,
					EthLog: ethtypes.EthLog{
						TransactionIndex: 0,
						LogIndex:         1,
						Address:          emitter,
						Data:             eventData,
						Topics: []ethtypes.EthHash{
							createEthHash(t, "0x25eaabaf991947ec22f473a02c14ffbcc08ffe2cef8d81ac12b6db2c14ce23a0"),
							createEthHash(t, "0xab8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
							createEthHash(t, "0xbb8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
						},
					},
				},
				{
					TransactionCid: txCID,
					EthLog: ethtypes.EthLog{
						TransactionIndex: 1,
						LogIndex:         0,
						Address:          emitter,
						Data:             eventData,
						Topics: []ethtypes.EthHash{
							createEthHash(t, "0x25eaabaf991947ec22f473a02c14ffbcc08ffe2cef8d81ac12b6db2c14ce23a0"),
							createEthHash(t, "0xab8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
							createEthHash(t, "0xbb8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
						},
					},
				},
				{
					TransactionCid: txCID,
					EthLog: ethtypes.EthLog{
						TransactionIndex: 1,
						LogIndex:         1,
						Address:          emitter,
						Data:             eventData,
						Topics: []ethtypes.EthHash{
							createEthHash(t, "0x25eaabaf991947ec22f473a02c14ffbcc08ffe2cef8d81ac12b6db2c14ce23a0"),
							createEthHash(t, "0xab8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
							createEthHash(t, "0xbb8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738"),
						},
					},
				},
			},
			wantMetadata: map[string]any{
				"data": eventDataHex,
				"topics": []string{
					"0x25eaabaf991947ec22f473a02c14ffbcc08ffe2cef8d81ac12b6db2c14ce23a0",
					"0xab8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738",
					"0xbb8653edf9f51785664a643b47605a7ba3d917b5339a0724e7642c114d0e4738",
				},
			},
		},
	}

	for i := range tb {
		tt := tb[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			eventsData := types.EventsData{
				Tipset:  tipset,
				EthLogs: tt.ethLogs,
			}

			events, err := parser.ParseEthLogs(ctx, eventsData)
			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, events)
			assert.NotEmpty(t, events.ParsedEvents)

			gotMetadata := map[string]any{}
			err = json.Unmarshal([]byte(events.ParsedEvents[0].Metadata), &gotMetadata)
			assert.NoError(t, err)

			if len(tt.ethLogs[0].Topics) > 0 && tt.wantSig {
				// manually chosen signature of hash: 0x25eaabaf991947ec22f473a02c14ffbcc08ffe2cef8d81ac12b6db2c14ce23a0
				assert.EqualValues(t, "l1InfoLeafMap(uint256)", events.ParsedEvents[0].SelectorSig)
			}

			assert.EqualValues(t, tt.wantMetadata["data"], gotMetadata["data"])
			assert.ElementsMatch(t, tt.wantMetadata["topics"], gotMetadata["topics"])

			// check if IDs are unique for all events
			foundIDs := map[string]bool{}
			for idx, evt := range events.ParsedEvents {
				wantID := tools.BuildId(tipset.GetCidString(), txCID, fmt.Sprint(idx), types.EventTypeEVM)
				gotID := evt.ID

				assert.EqualValues(t, wantID, gotID)
				assert.NotContains(t, foundIDs, gotID)
				foundIDs[gotID] = true
			}

			assert.EqualValues(t, emitter.String(), events.ParsedEvents[0].Emitter)
		})
	}
}

func TestParser_MultisigEventsFromTxs(t *testing.T) {
	type expectedResults struct {
		proposals    []types.MultisigProposal
		multisigInfo []types.MultisigInfo
	}
	tests := []struct {
		name    string
		version string
		url     string
		height  string
		results expectedResults
	}{
		{
			name:    "multisig events height 14107",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "14107",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 14107, MultisigAddress: "f080", ProposalID: 0, Signer: "f0103", ActionType: "Propose", TxTypeToExecute: "AddVerifier", Value: `{"Address":"f1zo7ub42i3s5cutljzjuqwnltt4xxm4y4f7l5s2i","Allowance":"100000000000000"}`},
				},
				multisigInfo: []types.MultisigInfo{},
			},
		},
		{
			name:    "multisig events height 1467665",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1467665",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 1467665, MultisigAddress: "f080", ProposalID: 11, Signer: "f018896", ActionType: "Approve", TxTypeToExecute: "", Value: `{"ID":11,"ProposalHash":"/jgVZzOjfHFnrI5K514wyJ+WSVNtLQhthbCrDsX+Dmg="}`},
				},
				multisigInfo: []types.MultisigInfo{},
			},
		},
		{
			name:    "multisig events height 197673",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "197673",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 197673, MultisigAddress: "f03735", ProposalID: 1, Signer: "f01014", ActionType: "Approve", TxTypeToExecute: "", Value: "{\"ID\":1,\"ProposalHash\":null}"},
				},
				multisigInfo: []types.MultisigInfo{
					{Height: 197673, MultisigAddress: "f03735", TxCid: "bafy2bzacedr3hke3xt2jvtret2yalvhkpctefwsgddqyziggcfmgurd7igqaq", Signer: "f03735", ActionType: "AddSigner", Value: "{\"Signer\":\"f15xwdubazj7aft6ylmiw54fa27zyyl3rpc6olgcy\",\"Increase\":false}"},
				},
			},
		},
		{
			name:    "multisig events height 78689",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "78689",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 78689, MultisigAddress: "f02412", ProposalID: 0, Signer: "f02252", ActionType: "Propose", TxTypeToExecute: "UniversalReceiverHook", Value: `{"Value":"gUkALcv0hA7KAAA="}`},
				},
				multisigInfo: []types.MultisigInfo{},
			},
		},

		{
			name:    "multisig events height 47645",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "47645",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 47645, MultisigAddress: "f22ny34zaozvfsffk445tazmohsvygits3763xpuy", ProposalID: 1, Signer: "f01148", ActionType: "Propose", TxTypeToExecute: "ChangeNumApprovalsThreshold", Value: "{\"NewThreshold\":3}"},
				},
				multisigInfo: []types.MultisigInfo{},
			},
		},

		{
			name:    "multisig events height 39035",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "39035",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 39035, MultisigAddress: "f23pa4gt4jgkl55drdyzb7dscjzdfh725u45xzwsy", ProposalID: 1, Signer: "f01717", ActionType: "Propose", TxTypeToExecute: "ChangeOwnerAddress", Value: `{"Value":"f01816"}`},
				},
				multisigInfo: []types.MultisigInfo{},
			},
		},

		{
			name:    "multisig events height 47635",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "47635",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 47635, MultisigAddress: "f22ny34zaozvfsffk445tazmohsvygits3763xpuy", ProposalID: 0, Signer: "f01148", ActionType: "Propose", TxTypeToExecute: "AddSigner", Value: "{\"Increase\":false,\"Signer\":\"f3vyx6j6jwrpw4dfspselowh6p4sg6cewgykfvnyomtma5eh4exgkkj4my6ki2sax7zdiavi2wbt3dbet3svxq\"}"},
				},
				multisigInfo: []types.MultisigInfo{},
			},
		},
		{
			name:    "multisig events height 38940",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "38940",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 38940, MultisigAddress: "f23pa4gt4jgkl55drdyzb7dscjzdfh725u45xzwsy", ProposalID: 0, Signer: "f01410", ActionType: "Propose", TxTypeToExecute: "Send", Value: "{\"Value\":\"30000000000000000\"}"},
				},
				multisigInfo: []types.MultisigInfo{},
			},
		},
		{
			name:    "multisig events height 1698055",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1698055",
			results: expectedResults{
				proposals: []types.MultisigProposal{},
				multisigInfo: []types.MultisigInfo{
					{Height: 1698055, MultisigAddress: "f0123724", Signer: "f01", ActionType: "Constructor", TxCid: "bafy2bzacednj5rv7pbdgyvq2ztknz5xj3eazfqanxiy5za6dpb4kgv3vtd72w", Value: "{\"Signers\":[\"f1fik4crqpv33laa6gvf23vz3sjpioka4go47e2hi\",\"f1hjvq6aays3ohfzxo7sw353esyscnbzabblxs2pq\",\"f3vwrfmyc6wnomefflu2rcejrbvi4zmczq2pcibf2rl7udbwbf4eyjkpftwyhdigevfprnj5g32h5liaxf56qq\"],\"NumApprovalsThreshold\":2,\"UnlockDuration\":0,\"StartEpoch\":0}"},
				},
			},
		},
		{
			name:    "multisig events height 1576593",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1576593",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 1576593, MultisigAddress: "f2b3v3bp55krpaqz24fxmlgggbz3gaik6fv5f7ryy", ProposalID: 138, Signer: "f091402", ActionType: "Cancel", TxTypeToExecute: "", Value: "{\"ID\":138,\"ProposalHash\":\"vXH0+s6OtR7wEs0aVsxBgB1/bgOqCSoZ/ImHyBlDVcw=\"}"},
				},
				multisigInfo: []types.MultisigInfo{},
			},
		},
		{
			name:    "multisig events height 1572087",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1572087",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 1572087, MultisigAddress: "f2b3v3bp55krpaqz24fxmlgggbz3gaik6fv5f7ryy", ProposalID: 105, Signer: "f091402", ActionType: "Propose", TxTypeToExecute: "ChangeNumApprovalsThreshold", Value: "{\"NewThreshold\":2}"},
				},
				multisigInfo: []types.MultisigInfo{
					{Height: 1572087, MultisigAddress: "f2b3v3bp55krpaqz24fxmlgggbz3gaik6fv5f7ryy", TxCid: "bafy2bzacec74jgx36mdxmggmoxbjhub3cfnzvfu7dujagdk3il7ttz7emu4q4", Signer: "f0110268", ActionType: "ChangeNumApprovalsThreshold", Value: "{\"NewThreshold\":2}"},
				},
			},
		},
		{
			name:    "multisig events height 1552242",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1552242",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 1552242, MultisigAddress: "f2t7urdjxp5jf3su5qyf4i25encrozjws6k2uxg2i", ProposalID: 3, Signer: "f06067", ActionType: "Approve", TxTypeToExecute: "", Value: "{\"ID\":3,\"ProposalHash\":\"lMtCwZTOT/0X+G4aplxpQN9xyPWvLLwpUObAVRTqUSI=\"}"},
				},
				multisigInfo: []types.MultisigInfo{
					{Height: 1552242, MultisigAddress: "f2t7urdjxp5jf3su5qyf4i25encrozjws6k2uxg2i", TxCid: "bafy2bzacecutjnons7vvzlamg6sekcnmo5hbkfobdo52p2minokt2rz6vgsqy", Signer: "f059513", ActionType: "SwapSigner", Value: "{\"From\":\"f1dywbadna5yyf546mloeoc7gxrzj7n5uog6llv5y\",\"To\":\"f16sfr4wmxu7ouxayxqqmacmgdfqfbasm4qr472fq\"}"},
				},
			},
		},
		{
			name:    "multisig events height 1352134",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1352134",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 1352134, MultisigAddress: "f2kpwyxvbr547eaikwjavx6bs4otae3cqbn5u2t2y", ProposalID: 0, Signer: "f019764", ActionType: "Propose", TxTypeToExecute: "LockBalance", Value: "{\"Amount\":\"1000000000000000000\",\"StartEpoch\":1352039,\"UnlockDuration\":876000}"},
				},
				multisigInfo: []types.MultisigInfo{
					{Height: 1352134, MultisigAddress: "f2kpwyxvbr547eaikwjavx6bs4otae3cqbn5u2t2y", Signer: "f066958", ActionType: "LockBalance", TxCid: "bafy2bzacecfvtxvsfkjrj6odvuii7m5bf52vqni66nwk4clwp5j5x6ovothco", Value: "{\"StartEpoch\":1352039,\"UnlockDuration\":876000,\"Amount\":\"1000000000000000000\"}"},
				},
			},
		},
		{
			name:    "multisig events height 1334035",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1334035",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 1334035, MultisigAddress: "f2h4xqc7krcpfulaqch6hxphsp6ze5fwobfrpur2i", ProposalID: 2, Signer: "f06068", ActionType: "Approve", TxTypeToExecute: "", Value: "{\"ID\":2,\"ProposalHash\":\"V5xMhdHMYFd7uwnILnd1BeH6SoksOKTI4KtUyZXOS4k=\"}"},
				},
				multisigInfo: []types.MultisigInfo{
					{Height: 1334035, TxCid: "bafy2bzaceasprlgdy4dbb2cxzwo4opofxk26vkpw3fe3qf5oxj2yjsb7scjoq", MultisigAddress: "f2h4xqc7krcpfulaqch6hxphsp6ze5fwobfrpur2i", Signer: "f063654", ActionType: "RemoveSigner", Value: "{\"Signer\":\"f1fbagfbmhk52hhbih2yt2jixkbisoqtrg4k2kn7a\",\"Decrease\":true}"},
				},
			},
		},
		{
			name:    "multisig events height 1289201",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1289201",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 1289201, MultisigAddress: "f064773", ProposalID: 0, Signer: "f064766", ActionType: "Propose", TxTypeToExecute: "SwapSigner", Value: "{\"From\":\"f3vwq5mw6sagjzqap73q56xayzmnrlqpvlecgcduwqmpsr33cngoszviq4eeet7gc5j3he2kf34hmskecjvqva\",\"To\":\"f3sg5mydbqdszt6wld3sjofhotutji5r2vbi5nvraybulexajcqg2fdas6sq7oiihdeqmw7ii3xdzlx723oeja\"}"},
				},
				multisigInfo: []types.MultisigInfo{},
			},
		},
		{
			name:    "multisig events height 1258459",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1258459",
			results: expectedResults{
				proposals: []types.MultisigProposal{},
				multisigInfo: []types.MultisigInfo{
					{Height: 1258459, MultisigAddress: "f063814", TxCid: "bafy2bzaced6uosdwea2ztyao56umwjza5qzpveg73afohzh3wgoclgtfhtpek", Signer: "f01", ActionType: "Constructor", Value: "{\"Signers\":[\"f16xlkjp3dcfrsb257duoqfgj7glo2uvvgxyy4gmy\",\"f1dywbadna5yyf546mloeoc7gxrzj7n5uog6llv5y\",\"f1fbagfbmhk52hhbih2yt2jixkbisoqtrg4k2kn7a\"],\"NumApprovalsThreshold\":2,\"UnlockDuration\":0,\"StartEpoch\":0}"},
				},
			},
		},
		{
			name:    "multisig events height 1256171",
			version: v2.NodeVersionsSupported[0],
			url:     calibNextNodeUrl,
			height:  "1256171",
			results: expectedResults{
				proposals: []types.MultisigProposal{
					{Height: 1256171, MultisigAddress: "f063719", ProposalID: 1, Signer: "f063720", ActionType: "Propose", TxTypeToExecute: "RemoveSigner", Value: "{\"Decrease\":false,\"Signer\":\"f1bsqp2nixftm5kacppzrsjkv62ot3kckucthu7ca\"}"},
				},
				multisigInfo: []types.MultisigInfo{
					{Height: 1256171, MultisigAddress: "f063719", TxCid: "bafy2bzacecuhvthgttyv7q3q53p4lqhfkkdh2wktaxywtfarofehqvtsgifnw", Signer: "f063719", ActionType: "RemoveSigner", Value: "{\"Signer\":\"f1bsqp2nixftm5kacppzrsjkv62ot3kckucthu7ca\",\"Decrease\":false}"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p *FilecoinParser
			var err error
			if tt.url == nodeUrl {
				p, err = NewFilecoinParserWithActorV2(getLib(tt.url), mainnetCacheDataSource, gLogger)
			} else {
				p, err = NewFilecoinParserWithActorV2(getLib(tt.url), calibNextNodeCacheDataSource, gLogger)
			}
			require.NoError(t, err)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			txsData := types.TxsData{
				EthLogs:  ethlogs,
				Tipset:   tipset,
				Traces:   traces,
				Metadata: types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: tt.version}},
			}

			parsedResult, err := p.ParseTransactions(context.Background(), txsData)
			require.NoError(t, err)
			require.NotNil(t, parsedResult.Txs)

			tipsetCid := txsData.Tipset.GetCidString()
			tipsetKey := txsData.Tipset.Key()
			events, err := p.ParseMultisigEvents(context.Background(), parsedResult.Txs, tipsetCid, tipsetKey)
			require.NoError(t, err)
			require.NotNil(t, events)

			require.Len(t, events.Proposals, len(tt.results.proposals), fmt.Sprintf("Expected %d proposals, but got %d", len(tt.results.proposals), len(events.Proposals)))
			for i, expected := range tt.results.proposals {
				assert.Equal(t, expected.MultisigAddress, events.Proposals[i].MultisigAddress, fmt.Sprintf("Mismatch in MultisigAddress at proposal index %d: expected %s, got %s", i, expected.MultisigAddress, events.Proposals[i].MultisigAddress))
				assert.Equal(t, expected.ProposalID, events.Proposals[i].ProposalID, fmt.Sprintf("Mismatch in ProposalID at proposal index %d: expected %d, got %d", i, expected.ProposalID, events.Proposals[i].ProposalID))
				assert.Equal(t, expected.Signer, events.Proposals[i].Signer, fmt.Sprintf("Mismatch in Signer at proposal index %d: expected %s, got %s", i, expected.Signer, events.Proposals[i].Signer))
				assert.Equal(t, expected.ActionType, events.Proposals[i].ActionType, fmt.Sprintf("Mismatch in ActionType at proposal index %d: expected %s, got %s", i, expected.ActionType, events.Proposals[i].ActionType))
				assert.Equal(t, expected.TxTypeToExecute, events.Proposals[i].TxTypeToExecute, fmt.Sprintf("Mismatch in TxTypeToExecute at proposal index %d: expected %s, got %s", i, expected.TxTypeToExecute, events.Proposals[i].TxTypeToExecute))
				compareJSONKeys(t, expected.Value, events.Proposals[i].Value)
				// assert.EqualValuesf(t, expected.Value, events.Proposals[i].Value, fmt.Sprintf("Mismatch in Value at proposal index %d: expected %s, got %s", i, expected.Value, events.Proposals[i].Value))
			}

			require.Len(t, events.MultisigInfo, len(tt.results.multisigInfo), fmt.Sprintf("Expected %d multisig info entries, but got %d", len(tt.results.multisigInfo), len(events.MultisigInfo)))
			for i, expected := range tt.results.multisigInfo {
				assert.Equal(t, expected.MultisigAddress, events.MultisigInfo[i].MultisigAddress, fmt.Sprintf("Mismatch in MultisigAddress at multisig info index %d: expected %s, got %s", i, expected.MultisigAddress, events.MultisigInfo[i].MultisigAddress))
				assert.Equal(t, expected.TxCid, events.MultisigInfo[i].TxCid, fmt.Sprintf("Mismatch in TxCid at multisig info index %d: expected %s, got %s", i, expected.TxCid, events.MultisigInfo[i].TxCid))
				assert.Equal(t, expected.Signer, events.MultisigInfo[i].Signer, fmt.Sprintf("Mismatch in Signer at multisig info index %d: expected %s, got %s", i, expected.Signer, events.MultisigInfo[i].Signer))
				assert.Equal(t, expected.ActionType, events.MultisigInfo[i].ActionType, fmt.Sprintf("Mismatch in ActionType at multisig info index %d: expected %s, got %s", i, expected.ActionType, events.MultisigInfo[i].ActionType))
				compareJSONKeys(t, expected.Value, events.MultisigInfo[i].Value)
				// assert.EqualValuesf(t, expected.Value, events.MultisigInfo[i].Value, fmt.Sprintf("Mismatch in Value at multisig info index %d: expected %s, got %s", i, expected.Value, events.MultisigInfo[i].Value))
			}
		})
	}
}

func compareJSONKeys(t *testing.T, expected, actual string) {
	expectedMap := make(map[string]any)
	actualMap := make(map[string]any)

	err := json.Unmarshal([]byte(expected), &expectedMap)
	require.NoError(t, err)
	err = json.Unmarshal([]byte(actual), &actualMap)
	require.NoError(t, err)

	for k, expectedValue := range expectedMap {
		actualValue, ok := actualMap[k]
		if !ok {
			assert.Failf(t, "Key %s not found in actual map", k)
		}
		assert.EqualValuesf(t, expectedValue, actualValue, fmt.Sprintf("Mismatch in Value at key %s: expected %v, got %v", k, expectedValue, actualValue))
	}
}

func TestParseGenesis(t *testing.T) {
	tests := []struct {
		name              string
		network           string
		nodeUrl           string
		cacheDataSource   common.DataSource
		expectedTxs       int
		expectedBlockCid  string
		expectedTipsetCid string
	}{
		{
			name:              "mainnet",
			network:           "mainnet",
			nodeUrl:           nodeUrl,
			cacheDataSource:   mainnetCacheDataSource,
			expectedTxs:       21,
			expectedBlockCid:  "bafy2bzacecnamqgqmifpluoeldx7zzglxcljo6oja4vrmtj7432rphldpdmm2",
			expectedTipsetCid: "bafy2bzacea3l7hchfijz5fvswab36fxepf6oagecp5hrstmol7zpm2l4tedf6",
		},
		{
			name:              "calibration",
			network:           "calibration",
			nodeUrl:           calibNextNodeUrl,
			cacheDataSource:   calibNextNodeCacheDataSource,
			expectedTxs:       8,
			expectedBlockCid:  "bafy2bzacecyaggy24wol5ruvs6qm73gjibs2l2iyhcqmvi7r7a4ph7zx3yqd4",
			expectedTipsetCid: "bafy2bzacebbqulnhstepn4hdbgaxf2grjqxgu6itf53ml7tvyps2z7f726s32",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := tt.network
			genesisBalances, genesisTipset, err := getStoredGenesisData(network)
			require.NoError(t, err)

			p, err := NewFilecoinParser(getLib(tt.nodeUrl), tt.cacheDataSource, gLogger)
			require.NoError(t, err)
			actualTxs, _ := p.ParseGenesis(genesisBalances, genesisTipset)

			assert.Equal(t, len(actualTxs), tt.expectedTxs)
			assert.Equal(t, actualTxs[0].BlockCid, tt.expectedBlockCid)
			assert.Equal(t, actualTxs[0].TipsetCid, tt.expectedTipsetCid)
		})
	}

}

func TestParseGenesisMultisig(t *testing.T) {
	tests := []struct {
		name            string
		network         string
		nodeUrl         string
		cacheDataSource common.DataSource
	}{
		{
			name:            "mainnet",
			network:         "mainnet",
			nodeUrl:         nodeUrl,
			cacheDataSource: mainnetCacheDataSource,
		},
		{
			name:            "calibration",
			network:         "calibration",
			nodeUrl:         calibNextNodeUrl,
			cacheDataSource: calibNextNodeCacheDataSource,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := tt.network
			genesisFilePath := filepath.Join("./data/genesis", fmt.Sprintf("%s_genesis_multisig_info.json", network))
			content, err := os.ReadFile(genesisFilePath)
			require.NoError(t, err)

			var expectedMultisigInfo []*types.MultisigInfo
			err = json.Unmarshal(content, &expectedMultisigInfo)
			require.NoError(t, err)

			genesisBalances, genesisTipset, err := getStoredGenesisData(network)
			require.NoError(t, err)

			p, err := NewFilecoinParser(getLib(tt.nodeUrl), tt.cacheDataSource, gLogger)
			require.NoError(t, err)

			ctx := context.Background()
			gotMultiSigInfo, err := p.ParseGenesisMultisig(ctx, genesisBalances, genesisTipset)
			require.NoError(t, err)
			require.NotNil(t, gotMultiSigInfo)
			require.Equal(t, len(expectedMultisigInfo), len(gotMultiSigInfo))
			require.ElementsMatch(t, expectedMultisigInfo, gotMultiSigInfo)
		})
	}
}

func TestParser_ActorVersionComparison(t *testing.T) {
	type expectedResults struct {
		totalTraces  int
		totalAddress int
		totalTxCids  int
	}
	tests := []struct {
		name       string
		version    string
		url        string
		height     string
		results    expectedResults
		shouldFail bool
	}{
		{
			name:    "parser with traces from v1",
			version: v1.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "2907480",
			results: expectedResults{
				totalTraces:  650,
				totalAddress: 232,
				totalTxCids:  102,
			},
		},
		{
			name:    "parser with traces from v2",
			version: v2.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "2907520",
			results: expectedResults{
				totalTraces:  910,
				totalAddress: 234,
				totalTxCids:  151,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3573062",
			results: expectedResults{
				totalTraces:  774,
				totalAddress: 209,
				totalTxCids:  121,
			},
		},
		{
			name:    "should fail (actorsV2 fixed eth_address parsing in exec): parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3573064",
			results: expectedResults{
				totalTraces:  734,
				totalAddress: 206,
				totalTxCids:  101,
			},
			shouldFail: true,
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3573066",
			results: expectedResults{
				totalTraces:  1121,
				totalAddress: 274,
				totalTxCids:  187,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.26 (calib)",
			version: v2.NodeVersionsSupported[2],
			url:     calibNextNodeUrl,
			height:  "1419335",
			results: expectedResults{
				totalTraces:  37,
				totalAddress: 18,
				totalTxCids:  5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pv1 *FilecoinParser
			var pv2 *FilecoinParser
			var err1 error
			var err2 error
			if tt.url == nodeUrl {
				pv1, err1 = NewFilecoinParser(getLib(tt.url), mainnetCacheDataSource, gLogger)
				pv2, err2 = NewFilecoinParserWithActorV2(getLib(tt.url), mainnetCacheDataSource, gLogger)
			} else {
				pv1, err1 = NewFilecoinParser(getLib(tt.url), calibNextNodeCacheDataSource, gLogger)
				pv2, err2 = NewFilecoinParserWithActorV2(getLib(tt.url), calibNextNodeCacheDataSource, gLogger)
			}
			require.NoError(t, err1)
			require.NoError(t, err2)
			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			fmt.Println("tipset", tipset.Height())
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			txsData := types.TxsData{
				EthLogs:  ethlogs,
				Tipset:   tipset,
				Traces:   traces,
				Metadata: types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: tt.version}},
			}

			wg := sync.WaitGroup{}
			wg.Add(2)
			var parsedResultActorV1 *types.TxsParsedResult
			var parsedResultActorV2 *types.TxsParsedResult
			go func() {
				defer wg.Done()
				parsedResultActorV1, err1 = pv1.ParseTransactions(context.Background(), txsData)
			}()
			go func() {
				defer wg.Done()
				parsedResultActorV2, err2 = pv2.ParseTransactions(context.Background(), txsData)
			}()

			wg.Wait()

			require.NoErrorf(t, err1, "error parsing v1: %s", err1)
			require.NoErrorf(t, err2, "error parsing v2: %s", err2)

			require.NotNil(t, parsedResultActorV1.Txs)
			require.NotNil(t, parsedResultActorV2.Txs)

			require.Equal(t, len(parsedResultActorV1.Txs), len(parsedResultActorV2.Txs))

			assert.Equal(t, len(parsedResultActorV1.TxCids), len(parsedResultActorV2.TxCids))
			assert.Equal(t, tt.results.totalTraces, len(parsedResultActorV1.Txs))
			assert.Equal(t, tt.results.totalAddress, parsedResultActorV1.Addresses.Len())
			assert.Equal(t, tt.results.totalTxCids, len(parsedResultActorV1.TxCids))

			require.Equalf(t, parsedResultActorV1.Addresses.Len(), parsedResultActorV2.Addresses.Len(), "v1 %d , v2 %d", parsedResultActorV1.Addresses.Len(), parsedResultActorV2.Addresses.Len())
			// compare tx metadata
			failedTxType := map[string]string{}
			for i, tx := range parsedResultActorV1.Txs {
				var metadataV1 map[string]interface{}
				err := json.Unmarshal([]byte(tx.TxMetadata), &metadataV1)
				if err != nil {
					t.Fatalf("Error unmarshalling v1 tx metadata: %s", err.Error())
				}
				var metadataV2 map[string]interface{}
				err = json.Unmarshal([]byte(parsedResultActorV2.Txs[i].TxMetadata), &metadataV2)
				if err != nil {
					t.Fatalf("Error unmarshalling v2 tx metadata: %s", err.Error())
				}

				if tt.shouldFail {
					continue
				}

				if metadataV1[parser.ParamsKey] != nil {
					// multisig propose correctly parses return params in v2
					if tx.TxType != parser.MethodPropose {
						require.EqualValuesf(t, metadataV1[parser.ParamsKey], metadataV2[parser.ParamsKey], fmt.Sprintf("tx_type: %s \n V1: %s \n V2: %s", tx.TxType, tx.TxMetadata, parsedResultActorV2.Txs[i].TxMetadata))
					}
				}
				if metadataV1[parser.ReturnKey] != nil {
					// ClaimAllocations return struct changed to support slices.
					// ActivateDeals metadata was fixed to parse correctly in v2.
					if tx.TxType != parser.MethodClaimAllocations && tx.TxType != parser.MethodActivateDeals {
						require.EqualValuesf(t, metadataV1[parser.ReturnKey], metadataV2[parser.ReturnKey], fmt.Sprintf("tx_type: %s \n V1: %s \n V2: %s", tx.TxType, tx.TxMetadata, parsedResultActorV2.Txs[i].TxMetadata))
					}
				}

			}
			assert.Equal(t, 0, len(failedTxType), "Tx metadata mismatch for tx_type: %v", failedTxType)
		})
	}
}

func getStoredGenesisData(network string) (*types.GenesisBalances, *types.ExtendedTipSet, error) {
	balancesFilePath := filepath.Join("./data/genesis", fmt.Sprintf("%s_genesis_balances.json", network))
	tipsetFilePath := filepath.Join("./data/genesis", fmt.Sprintf("%s_genesis_tipset.json", network))

	var balances types.GenesisBalances
	var tipset types.ExtendedTipSet

	balancesFileContent, err := os.ReadFile(balancesFilePath)
	if err != nil {
		gLogger.Errorf("Error reading file '%s': %s", balancesFilePath, err.Error())
		return nil, nil, err
	}

	err = json.Unmarshal(balancesFileContent, &balances)
	if err != nil {
		gLogger.Errorf("Error unmarshalling genesis balances: %s", err.Error())
		return nil, nil, err
	}

	tipsetFileContent, err := os.ReadFile(tipsetFilePath)
	if err != nil {
		gLogger.Errorf("Error reading file '%s': %s", tipsetFilePath, err.Error())
		return nil, nil, err
	}

	err = json.Unmarshal(tipsetFileContent, &tipset)
	if err != nil {
		gLogger.Errorf("Error unmarshalling genesis tipset: %s", err.Error())
		return nil, nil, err
	}

	return &balances, &tipset, nil
}

func createEthHash(t *testing.T, hash string) ethtypes.EthHash {
	var ethHash ethtypes.EthHash
	err := ethHash.UnmarshalJSON([]byte(fmt.Sprintf(`"%s"`, hash)))
	assert.NoError(t, err)
	return ethHash
}

func createTopic(t *testing.T, hash string) []byte {
	topic := createEthHash(t, hash)

	topicBytes := make([]byte, ethtypes.EthHashLength)
	n := copy(topicBytes, topic[:ethtypes.EthHashLength])
	assert.Equal(t, ethtypes.EthHashLength, n)

	return topicBytes
}
