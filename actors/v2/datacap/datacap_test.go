package datacap_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/v2/datacap"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

type testFn func(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error)

var expectedData []byte
var expected map[string]any

var network string

func TestMain(m *testing.M) {
	network = "mainnet"
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		panic(err)
	}
	var err error
	expectedData, err = tools.ReadActorSnapshot()
	if err != nil {
		panic(err)
	}
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
