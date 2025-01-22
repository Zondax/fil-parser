package eam_test

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/eam"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type testFn func(height int64, raw, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error)

type test struct {
	name     string
	version  string
	url      string
	height   int64
	expected map[string]any
	addr     *types.AddressInfo
}

func TestParseCreateExternal(t *testing.T) {
	tests := []test{}

	runTest(t, eam.ParseCreateExternal, tests)
}

func TestParseCreate(t *testing.T) {
	tests := []test{}

	runTest(t, eam.ParseCreate, tests)
}

func TestParseCreate2(t *testing.T) {
	tests := []test{}

	runTest(t, eam.ParseCreate2, tests)
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
				result, addrInfo, err := fn(tt.height, trace.Msg.Params, trace.MsgRct.Return, trace.Msg.Cid())
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
				require.Equal(t, addrInfo, tt.addr)
			}
		})
	}
}
