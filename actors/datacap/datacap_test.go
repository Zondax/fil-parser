package datacap_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/datacap"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

//go:embed expected.json
var expectedData []byte
var expected map[string]any

func TestMain(m *testing.M) {
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestNameExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("NameExported", expected)
	require.NoError(t, err)

	runDatacapTest(t, datacap.NameExported, tests)
}

func TestSymbolExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("SymbolExported", expected)
	require.NoError(t, err)

	runDatacapTest(t, datacap.SymbolExported, tests)
}

func TestTotalSupplyExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("TotalSupplyExported", expected)
	require.NoError(t, err)

	runDatacapTest(t, datacap.TotalSupplyExported, tests)
}

func runDatacapTest(t *testing.T, fn func(rawReturn []byte) (map[string]interface{}, error), tests []tools.TestCase[map[string]any]) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := fn(trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}
