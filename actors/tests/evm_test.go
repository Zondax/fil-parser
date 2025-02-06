package actortest

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/parser"
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

			got, err := p.ParseEvm(tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: nil,
			})
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
	actor, err := p.GetActor(manifest.EvmKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range evmWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.EvmKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			got, _, err := actor.Parse(network, height, tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: nil,
			}, cid.Undef, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV2_EvmWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.EvmKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range evmWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.EvmKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, _, err := actor.Parse(network, height, tt.txType, &parser.LotusMessage{
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
	actor, err := p.GetActor(manifest.EvmKey)
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

	got, _, err := actor.Parse(network, height, parser.MethodInvokeContract, &parser.LotusMessage{
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
	actor, err := p.GetActor(manifest.EvmKey)
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

	got, _, err := actor.Parse(network, height, parser.MethodInvokeContractReadOnly, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: rawReturn,
	}, cid.Undef, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.NotNil(t, got)
}

func TestActorParserV2_EVMInvokeContract_whenCborUnmarshalFail(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.EvmKey)
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
	got, _, err := actor.Parse(network, height, parser.MethodInvokeContract, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: rawReturn,
	}, cid.Undef, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, got["Params"], "0x70a082310000000000000000000000001a5ef7ef64e3fb12be3b43edd77819dc7f034b1f")
	require.Equal(t, got["Return"], "0x00000000000000000000000000000000000000000000000698b81208dfe49012")
}
