package datacap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/datacap"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

func TestNameExported(t *testing.T) {
	tests := []test{}

	runDatacapTest(t, datacap.NameExported, tests)
}

func TestSymbolExported(t *testing.T) {
	tests := []test{}

	runDatacapTest(t, datacap.SymbolExported, tests)
}

func TestTotalSupplyExported(t *testing.T) {
	tests := []test{}

	runDatacapTest(t, datacap.TotalSupplyExported, tests)
}

func runDatacapTest(t *testing.T, fn func(rawReturn []byte) (map[string]interface{}, error), tests []test) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := fn(trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}
