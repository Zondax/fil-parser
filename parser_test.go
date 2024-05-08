package fil_parser

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	v1 "github.com/zondax/fil-parser/parser/v1"
	v2 "github.com/zondax/fil-parser/parser/v2"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/golem/pkg/zcache"

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
	nodeUrl           = "https://node-fil-mainnet-next.zondax.ch/rpc/v1"
	calibNextNodeUrl  = "https://node-fil-calibration-next.zondax.ch/rpc/v1"
	feeType           = "fee"
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

func getLotusClient(t *testing.T, url string) api.FullNode {
	lotusClient, _, err := client.NewFullNodeRPCV1(context.Background(), url, http.Header{})
	require.NoError(t, err)
	require.NotNil(t, lotusClient, "Lotus client should not be nil")

	return lotusClient
}

func getLib(t *testing.T, nodeURL string) *rosettaFilecoinLib.RosettaConstructionFilecoin {
	lotusClient := getLotusClient(t, nodeURL)

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
	require.NotNil(t, lib, "Rosetta lib should not be nil")
	return lib
}

func getCacheDataSource(t *testing.T, nodeURL string) common.DataSource {
	return common.DataSource{
		Node: getLotusClient(t, nodeURL),
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
		{
			name:    "parser with traces from v1",
			version: v1.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "2907480",
			results: expectedResults{
				totalTraces:  650,
				totalAddress: 98,
				totalTxCids:  99,
			},
		},
		{
			name:    "parser with traces from v1 and the corner case of duplicated fees with level 0",
			version: v1.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "845259",
			results: expectedResults{
				totalTraces:  31,
				totalAddress: 3,
				totalTxCids:  0,
			},
		},
		{
			name:    "parser with traces from v2",
			version: v2.NodeVersionsSupported[0],
			url:     nodeUrl,
			height:  "2907520",
			results: expectedResults{
				totalTraces:  907,
				totalAddress: 88,
				totalTxCids:  147,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3573062",
			results: expectedResults{
				totalTraces:  773,
				totalAddress: 70,
				totalTxCids:  118,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3573064",
			results: expectedResults{
				totalTraces:  734,
				totalAddress: 75,
				totalTxCids:  97,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.25",
			version: v2.NodeVersionsSupported[2],
			url:     nodeUrl,
			height:  "3573066",
			results: expectedResults{
				totalTraces:  1118,
				totalAddress: 102,
				totalTxCids:  177,
			},
		},
		{
			name:    "parser with traces from v2 and lotus 1.26 (calib)",
			version: v2.NodeVersionsSupported[2],
			url:     calibNextNodeUrl,
			height:  "1419335",
			results: expectedResults{
				totalTraces:  37,
				totalAddress: 11,
				totalTxCids:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lib := getLib(t, tt.url)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			logger, err := zap.NewDevelopment()
			require.NoError(t, err)

			p, err := NewFilecoinParser(lib, getCacheDataSource(t, tt.url), logger)
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
			lib := getLib(t, tt.url)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			logger, err := zap.NewDevelopment()
			require.NoError(t, err)

			p, err := NewFilecoinParser(lib, getCacheDataSource(t, tt.url), logger)
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
			lib := getLib(t, tt.url)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			logger, err := zap.NewDevelopment()
			require.NoError(t, err)

			p, err := NewFilecoinParser(lib, getCacheDataSource(t, tt.url), logger)
			require.NoError(t, err)

			txsData := types.TxsData{
				EthLogs:  ethlogs,
				Tipset:   tipset,
				Traces:   traces,
				Metadata: types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: "v1.22"}},
			}
			parsedResultV1, err := p.ParseTransactions(context.Background(), txsData)
			require.NoError(t, err)
			require.NotNil(t, parsedResultV1.Txs)
			require.NotNil(t, parsedResultV1.Addresses)

			txsData.Metadata = types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: "v1.23"}}
			parsedResultV2, err := p.ParseTransactions(context.Background(), txsData)
			require.NoError(t, err)
			require.NotNil(t, parsedResultV2.Txs)
			require.NotNil(t, parsedResultV2.Addresses)

			require.Equal(t, len(parsedResultV1.Txs), len(parsedResultV2.Txs))
			require.Equal(t, parsedResultV1.Addresses.Len(), parsedResultV2.Addresses.Len())
			require.Equal(t, len(parsedResultV1.TxCids), len(parsedResultV2.TxCids))

			for i := range parsedResultV1.Txs {
				require.True(t, parsedResultV1.Txs[i].Equal(*parsedResultV2.Txs[i]))
			}

			for i := range parsedResultV1.TxCids {
				require.True(t, reflect.DeepEqual(parsedResultV1.TxCids[i], parsedResultV2.TxCids[i]))
			}

			parsedResultV1.Addresses.Range(func(key string, value *types.AddressInfo) bool {
				v2Value, ok := parsedResultV2.Addresses.Get(key)
				require.True(t, ok)
				require.Equal(t, value, v2Value)
				return true
			})
		})
	}
}

func TestParser_ParseNativeEvents_FVM(t *testing.T) {
	height := uint64(8)
	// we need any random number for the test
	//nolint:gosec
	filAddress, err := address.NewIDAddress(uint64(rand.Int()))
	assert.NoError(t, err)
	tipsetCID := uuid.NewString()

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	parser, err := NewFilecoinParser(nil, getCacheDataSource(t, calibNextNodeUrl), logger)
	require.NoError(t, err)

	ipldNodeBuilder := basicnode.Prototype.String.NewBuilder()
	err = ipldNodeBuilder.AssignString("market_deals_event")
	assert.NoError(t, err)
	eventType, err := ipld.Encode(ipldNodeBuilder.Build(), dagcbor.Encode)
	assert.NoError(t, err)

	ipldNodeBuilder = basicnode.Prototype.Bytes.NewBuilder()
	err = ipldNodeBuilder.AssignBytes([]byte("test data"))
	assert.NoError(t, err)
	eventData, err := ipld.Encode(ipldNodeBuilder.Build(), dagcbor.Encode)
	assert.NoError(t, err)

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
					"value": "dGVzdCBkYXRh",
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
				Height:    height,
				TipsetCID: tipsetCID,
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
			assert.NoError(t, err)
			assert.NotNil(t, events)
			assert.NotEmpty(t, events.ParsedEvents)

			gotMetadata := map[int]map[string]any{}
			err = json.Unmarshal([]byte(events.ParsedEvents[0].Metadata), &gotMetadata)
			assert.NoError(t, err)

			for idx, v := range tt.wantMetadata {
				for entryKey, entryValue := range v {
					assert.EqualValues(t, entryValue, gotMetadata[idx][entryKey])
				}
			}
			assert.EqualValues(t, tipsetCID, events.ParsedEvents[0].TipsetCid)
			assert.EqualValues(t, tt.emitter.String(), events.ParsedEvents[0].Emitter)
			if len(tt.entries) > 0 { // only check for the selector_id if we have entries in the test case
				assert.EqualValues(t, "market_deals_event", events.ParsedEvents[0].SelectorID)
			}
			assert.EqualValues(t, tools.BuildId(tipsetCID, cid.Cid{}.String(), fmt.Sprint(0), types.EventTypeNative), events.ParsedEvents[0].ID)
		})
	}

}
func TestParser_ParseNativeEvents_EVM(t *testing.T) {
	height := uint64(8)
	ethAddress, err := address.NewDelegatedAddress(32, []byte{})
	assert.NoError(t, err)

	tipsetCID := uuid.NewString()

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

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	parser, err := NewFilecoinParser(nil, getCacheDataSource(t, calibNextNodeUrl), logger)
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
				Height:    height,
				TipsetCID: tipsetCID,
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
			assert.NoError(t, err)
			assert.NotNil(t, events)
			assert.NotEmpty(t, events.ParsedEvents)

			gotMetadata := map[string]any{}
			err = json.Unmarshal([]byte(events.ParsedEvents[0].Metadata), &gotMetadata)
			assert.NoError(t, err)

			assert.EqualValues(t, tt.wantMetadata["data"], gotMetadata["data"])
			assert.ElementsMatch(t, tt.wantMetadata["topics"], gotMetadata["topics"])
			assert.EqualValues(t, tools.BuildId(tipsetCID, cid.Cid{}.String(), fmt.Sprint(0), types.EventTypeEVM), events.ParsedEvents[0].ID)
			assert.EqualValues(t, tt.emitter.String(), events.ParsedEvents[0].Emitter)
			if len(tt.entries) > 0 { // only check the selector_id if there are entries in the test case
				assert.EqualValues(t, "0x013dbb9442ca9667baccc6230fcd5c1c4b2d4d2870f4bd20681d4d47cfd15184", events.ParsedEvents[0].SelectorID)
			}
		})
	}

}

func TestParser_ParseEthLogs(t *testing.T) {
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	cache, err := zcache.NewLocalCache(&zcache.LocalConfig{
		Logger: logger,
	})
	assert.NoError(t, err)

	var emitter ethtypes.EthAddress
	err = emitter.UnmarshalJSON([]byte(`"0xd4c5fb16488Aa48081296299d54b0c648C9333dA"`))
	assert.NoError(t, err)

	tipsetCID := cid.Cid{}.String()
	txCID := cid.Cid{}.String()
	height := uint64(8)

	eventData, err := json.Marshal(map[string]any{
		"x": "y",
		"a": "b",
	})
	assert.NoError(t, err)
	eventDataHex := hex.EncodeToString(eventData)

	parser, err := NewFilecoinParser(nil, getCacheDataSource(t, calibNextNodeUrl), logger)
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
						Address: emitter,
						Data:    eventData,
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
				Height:    height,
				TipsetCID: tipsetCID,
				EthLogs:   tt.ethLogs,
			}

			events, err := parser.ParseEthLogs(ctx, cache, eventsData)
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

			assert.EqualValues(t, tools.BuildId(tipsetCID, txCID, fmt.Sprint(0), types.EventTypeEVM), events.ParsedEvents[0].ID)
			assert.EqualValues(t, emitter.String(), events.ParsedEvents[0].Emitter)
		})
	}
}

func TestParseGenesis(t *testing.T) {
	network := "mainnet"
	genesisBalances, genesisTipset, err := getStoredGenesisData(network)
	if err != nil {
		t.Fatalf("Error getting genesis data: %s", err)
	}

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	lib := getLib(t, nodeUrl)
	p, err := NewFilecoinParser(lib, getCacheDataSource(t, nodeUrl), logger)
	assert.NoError(t, err)
	actualTxs, _ := p.ParseGenesis(genesisBalances, genesisTipset)

	assert.Equal(t, len(actualTxs), 21)
	assert.Equal(t, actualTxs[0].BlockCid, "bafy2bzacecnamqgqmifpluoeldx7zzglxcljo6oja4vrmtj7432rphldpdmm2")
	assert.Equal(t, actualTxs[0].TipsetCid, "bafy2bzacea3l7hchfijz5fvswab36fxepf6oagecp5hrstmol7zpm2l4tedf6")
}

func getStoredGenesisData(network string) (*types.GenesisBalances, *types.ExtendedTipSet, error) {
	balancesFilePath := filepath.Join("./data/genesis", fmt.Sprintf("%s_genesis_balances.json", network))
	tipsetFilePath := filepath.Join("./data/genesis", fmt.Sprintf("%s_genesis_tipset.json", network))

	var balances types.GenesisBalances
	var tipset types.ExtendedTipSet

	balancesFileContent, err := os.ReadFile(balancesFilePath)
	if err != nil {
		zap.S().Errorf("Error reading file '%s': %s", balancesFilePath, err.Error())
		return nil, nil, err
	}

	err = json.Unmarshal(balancesFileContent, &balances)
	if err != nil {
		zap.S().Errorf("Error unmarshalling genesis balances: %s", err.Error())
		return nil, nil, err
	}

	tipsetFileContent, err := os.ReadFile(tipsetFilePath)
	if err != nil {
		zap.S().Errorf("Error reading file '%s': %s", tipsetFilePath, err.Error())
		return nil, nil, err
	}

	err = json.Unmarshal(tipsetFileContent, &tipset)
	if err != nil {
		zap.S().Errorf("Error unmarshalling genesis tipset: %s", err.Error())
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
