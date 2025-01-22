package init_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	initActor "github.com/zondax/fil-parser/actors/init"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

type testFn func(height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error)

type test struct {
	name     string
	version  string
	url      string
	height   int64
	expected map[string]any
	address  *types.AddressInfo
}

func TestParseExec(t *testing.T) {
	tests := []test{}

	runTest(t, initActor.ParseExec, tests)
}

func TestParseExec4(t *testing.T) {
	tests := []test{}

	runTest(t, initActor.ParseExec4, tests)
}

func TestInitConstructor(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := initActor.InitConstructor(tt.height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}

}

func runTest(t *testing.T, fn testFn, tests []test) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				lotusMsg := &parser.LotusMessage{
					To:     trace.Msg.To,
					From:   trace.Msg.From,
					Method: trace.Msg.Method,
				}
				result, address, err := fn(tt.height, lotusMsg, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
				require.Equal(t, address, tt.address)
			}
		})
	}
}
