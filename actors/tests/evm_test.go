package actortest

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/metrics"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var evmWithParamsOrReturnTests = []struct {
	name   string
	txType string
	key    string
}{
	{
		name:   "Constructor",
		txType: parser.MethodConstructor,
		key:    parser.ParamsKey,
	},
	{
		name:   "Get Byte Code",
		txType: parser.MethodGetBytecode,
		key:    parser.ReturnKey,
	},
	{
		name:   "Resurrect",
		txType: parser.MethodResurrect,
		key:    parser.ParamsKey,
	},
	{
		name:   "Get ByteCode Hash",
		txType: parser.MethodGetBytecodeHash,
		key:    parser.ReturnKey,
	},
}

var evmWithParamsAndReturnTests = []struct {
	name   string
	txType string
}{
	{
		name:   "Invoke Contract Delegate",
		txType: parser.MethodInvokeContractDelegate,
	},
	{
		name:   "Get Storage At",
		txType: parser.MethodGetStorageAt,
	},
}

func TestActorParserV1_EvmWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range evmWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.EvmKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, err := p.ParseEvm(tt.txType, msg, msgRct)

			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV1_EvmWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range evmWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.EvmKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := p.ParseEvm(tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			})
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParserV1_EVMInvokeContract(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)
	rawParams, rawReturn, err := getParamsAndReturn(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, rawParams)
	require.NotNil(t, rawReturn)

	msg, err := deserializeMessage(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, msg)

	ethLogs, err := getEthLogs(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, ethLogs)

	got, err := p.ParseEvm(parser.MethodInvokeContract, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: rawReturn,
	})
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, got["Params"], "0x8381e182ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000008b21c7d96a349834dcfaddf871accda700b843e1")
	require.Equal(t, got["Return"], "0x00000000000000000000000000000000000000000000000698b81208dfe49012")
}

func TestActorParserV1_EVMInvokeContractReadOnly(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)
	rawParams, rawReturn, err := getParamsAndReturn(manifest.EvmKey, parser.MethodInvokeContractReadOnly)
	require.NoError(t, err)
	require.NotNil(t, rawParams)
	require.NotNil(t, rawReturn)

	msg, err := deserializeMessage(manifest.EvmKey, parser.MethodInvokeContractReadOnly)
	require.NoError(t, err)
	require.NotNil(t, msg)

	ethLogs, err := getEthLogs(manifest.EvmKey, parser.MethodInvokeContractReadOnly)
	require.NoError(t, err)
	require.NotNil(t, ethLogs)

	got, err := p.ParseEvm(parser.MethodInvokeContractReadOnly, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: rawReturn,
	})
	require.NoError(t, err)
	require.NotNil(t, got)
}

func TestActorParserV1_EVMInvokeContract_whenCborUnmarshalFail(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)
	_, rawReturn, err := getParamsAndReturn(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, rawReturn)

	msg, err := deserializeMessage(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, msg)

	ethLogs, err := getEthLogs(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, ethLogs)

	hexParamsString := "70a082310000000000000000000000001a5ef7ef64e3fb12be3b43edd77819dc7f034b1f"
	rawParams, _ := hex.DecodeString(hexParamsString)
	got, err := p.ParseEvm(parser.MethodInvokeContract, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: rawReturn,
	})
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, got["Params"], "0x70a082310000000000000000000000001a5ef7ef64e3fb12be3b43edd77819dc7f034b1f")
	require.Equal(t, got["Return"], "0x00000000000000000000000000000000000000000000000698b81208dfe49012")
}

func TestActorParserV2_EvmWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.EvmKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range evmWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.EvmKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), tt.txType, msg, msgRct, cid.Undef, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV2_EvmWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.EvmKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range evmWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.EvmKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, cid.Undef, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParserV2_EVMInvokeContract(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.EvmKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	rawParams, rawReturn, err := getParamsAndReturn(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, rawParams)
	require.NotNil(t, rawReturn)

	msg, err := deserializeMessage(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, msg)

	ethLogs, err := getEthLogs(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, ethLogs)

	got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), parser.MethodInvokeContract, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: rawReturn,
	}, cid.Undef, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, got["Params"], "0x8381e182ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000008b21c7d96a349834dcfaddf871accda700b843e1")
	require.Equal(t, got["Return"], "0x00000000000000000000000000000000000000000000000698b81208dfe49012")
}

func TestActorParserV2_EVMInvokeContractReadOnly(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.EvmKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	rawParams, rawReturn, err := getParamsAndReturn(manifest.EvmKey, parser.MethodInvokeContractReadOnly)
	require.NoError(t, err)
	require.NotNil(t, rawParams)
	require.NotNil(t, rawReturn)

	msg, err := deserializeMessage(manifest.EvmKey, parser.MethodInvokeContractReadOnly)
	require.NoError(t, err)
	require.NotNil(t, msg)

	ethLogs, err := getEthLogs(manifest.EvmKey, parser.MethodInvokeContractReadOnly)
	require.NoError(t, err)
	require.NotNil(t, ethLogs)

	got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), parser.MethodInvokeContractReadOnly, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: rawReturn,
	}, cid.Undef, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.NotNil(t, got)
}

func TestActorParserV2_EVMInvokeContract_whenCborUnmarshalFail(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.EvmKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	_, rawReturn, err := getParamsAndReturn(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, rawReturn)

	msg, err := deserializeMessage(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, msg)

	ethLogs, err := getEthLogs(manifest.EvmKey, parser.MethodInvokeContract)
	require.NoError(t, err)
	require.NotNil(t, ethLogs)

	hexParamsString := "70a082310000000000000000000000001a5ef7ef64e3fb12be3b43edd77819dc7f034b1f"
	rawParams, _ := hex.DecodeString(hexParamsString)
	got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), parser.MethodInvokeContract, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: rawReturn,
	}, cid.Undef, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, got["Params"], "0x70a082310000000000000000000000001a5ef7ef64e3fb12be3b43edd77819dc7f034b1f")
	require.Equal(t, got["Return"], "0x00000000000000000000000000000000000000000000000698b81208dfe49012")
}
