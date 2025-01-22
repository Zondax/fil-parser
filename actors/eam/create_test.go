package eam_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/eam"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
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

type testFn func(height int64, raw, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error)

func TestParseCreateExternal(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("CreateExternalExported", expected)
	require.NoError(t, err)

	runTest(t, eam.ParseCreateExternal, tests)
}

func TestParseCreate(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("CreateExported", expected)
	require.NoError(t, err)

	runTest(t, eam.ParseCreate, tests)
}

func TestParseCreate2(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("Create2Exported", expected)
	require.NoError(t, err)

	runTest(t, eam.ParseCreate2, tests)
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
				result, addrInfo, err := fn(tt.Height, trace.Msg.Params, trace.MsgRct.Return, trace.Msg.Cid())
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
				require.Equal(t, addrInfo, tt.Address)
			}
		})
	}
}
