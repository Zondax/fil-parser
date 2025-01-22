package init_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	initActor "github.com/zondax/fil-parser/actors/init"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
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

type testFn func(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error)

func TestParseExec(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseExec", expected)
	require.NoError(t, err)

	runTest(t, initActor.ParseExec, tests)
}

func TestParseExec4(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseExec4", expected)
	require.NoError(t, err)

	runTest(t, initActor.ParseExec4, tests)
}

func TestInitConstructor(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("InitConstructor", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := initActor.InitConstructor(tt.Network, tt.Height, trace.Msg.Params)
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
				lotusMsg := &parser.LotusMessage{
					To:     trace.Msg.To,
					From:   trace.Msg.From,
					Method: trace.Msg.Method,
				}
				result, address, err := fn(tt.Network, tt.Height, lotusMsg, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
				require.Equal(t, address, tt.Address)
			}
		})
	}
}
