package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"testing"
)

func TestActorParser_evmWithParamsOrReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name     string
		txType   string
		f        func([]byte) (map[string]interface{}, error)
		fileName string
		key      string
	}{
		{
			name:     "Constructor",
			txType:   parser.MethodConstructor,
			f:        p.evmConstructor,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Get Byte Code",
			txType:   parser.MethodGetBytecode,
			f:        p.getByteCode,
			fileName: "return",
			key:      parser.ReturnKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.EvmKey, tt.txType, tt.fileName)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParmasAndReturn(manifest.EvmKey, tt.txType)
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
