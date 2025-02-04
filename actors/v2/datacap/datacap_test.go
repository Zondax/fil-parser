package datacap_test

import (
	_ "embed"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/actors/v2/datacap"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
)

type testFn func(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error)

var expectedData []byte
var expected map[string]any

var network string

func TestMain(m *testing.M) {
	network = "mainnet"
	// if err := json.Unmarshal(expectedData, &expected); err != nil {
	// 	panic(err)
	// }
	// var err error
	// expectedData, err = tools.ReadActorSnapshot()
	// if err != nil {
	// 	panic(err)
	// }
	os.Exit(m.Run())
}

func TestDatacap(t *testing.T) {
	datacap := &datacap.Datacap{}
	testFns := map[string]func(rawReturn []byte) (map[string]interface{}, error){
		"NameExported":        datacap.NameExported,
		"SymbolExported":      datacap.SymbolExported,
		"TotalSupplyExported": datacap.TotalSupplyExported,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runDatacapTest(t, fn, tests)
		})
	}
}

func TestBurn(t *testing.T) {
	datacap := &datacap.Datacap{}
	testFns := map[string]testFn{
		"BurnExported":      datacap.BurnExported,
		"BurnFromExported":  datacap.BurnFromExported,
		"DestroyExported":   datacap.DestroyExported,
		"MintExported":      datacap.MintExported,
		"TransferExported":  datacap.TransferExported,
		"IncreaseAllowance": datacap.IncreaseAllowanceExported,
		"DecreaseAllowance": datacap.DecreaseAllowanceExported,
		"RevokeAllowance":   datacap.RevokeAllowanceExported,
		"AllowanceExported": datacap.AllowanceExported,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runTest(t, fn, tests)
		})
	}
}

func TestGranularityExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GranularityExported", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				datacap := &datacap.Datacap{}
				result, err := datacap.GranularityExported(tt.Network, tt.Height, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestMethodCoverage(t *testing.T) {
	d := &datacap.Datacap{}

	actorVersions := []any{
		datacapv10.Methods,
		datacapv11.Methods,
		datacapv12.Methods,
		datacapv13.Methods,
		datacapv14.Methods,
		datacapv15.Methods,
	}

	missingMethods := v2.MissingMethods(d, actorVersions)
	assert.Empty(t, missingMethods, "missing methods: %v", missingMethods)
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

func runTest(t *testing.T, fn testFn, tests []tools.TestCase[map[string]any]) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := fn(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}
