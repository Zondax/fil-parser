package actors

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
)

func TestActorParser_evmWithParamsOrReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte) (map[string]interface{}, error)
		key    string
	}{
		{
			name:   "Constructor",
			txType: parser.MethodConstructor,
			f:      p.evmConstructor,
			key:    parser.ParamsKey,
		},
		{
			name:   "Get Byte Code",
			txType: parser.MethodGetBytecode,
			f:      p.getByteCode,
			key:    parser.ReturnKey,
		},
		{
			name:   "Resurrect",
			txType: parser.MethodResurrect,
			f:      p.resurrect,
			key:    parser.ParamsKey,
		},
		{
			name:   "Get ByteCode Hash",
			txType: parser.MethodGetBytecodeHash,
			f:      p.getByteCodeHash,
			key:    parser.ReturnKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.EvmKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			got, err := tt.f(rawParams)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParser_evmWithParamsAndReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte, []byte) (map[string]interface{}, error)
	}{
		{
			name:   "Invoke Contract Delegate",
			txType: parser.MethodInvokeContractDelegate,
			f:      p.invokeContractDelegate,
		},
		{
			name:   "Get Storage At",
			txType: parser.MethodGetStorageAt,
			f:      p.getStorageAt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.EvmKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := tt.f(rawParams, rawReturn)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParser_invokeContract(t *testing.T) {
	p := getActorParser()
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

	got, err := p.invokeContract(rawParams, rawReturn)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, got["Params"], "0x8381e182ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000008b21c7d96a349834dcfaddf871accda700b843e1")
	require.Equal(t, got["Return"], "0x00000000000000000000000000000000000000000000000698b81208dfe49012")
}

func TestActorParser_invokeContractReadOnly(t *testing.T) {
	p := getActorParser()
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

	got, err := p.invokeContract(rawParams, rawReturn)
	require.NoError(t, err)
	require.NotNil(t, got)
}

func TestActorParser_invokeContract_whenCborUnmarshalFail(t *testing.T) {
	p := getActorParser()
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
	got, err := p.invokeContract(rawParams, rawReturn)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, got["Params"], "0x70a082310000000000000000000000001a5ef7ef64e3fb12be3b43edd77819dc7f034b1f")
	require.Equal(t, got["Return"], "0x00000000000000000000000000000000000000000000000698b81208dfe49012")
}
