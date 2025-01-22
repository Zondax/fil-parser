package verifiedregistry_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	verifiedregistry "github.com/zondax/fil-parser/actors/verifiedRegistry"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

//go:embed expected.json
var expectedData []byte
var expected map[string]any

var network string

func TestMain(m *testing.M) {
	network = "mainnet"
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		panic(err)
	}
	m.Run()
}

type testFn func(network string, height int64, raw []byte) (map[string]interface{}, error)
type testFn2 func(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error)

func TestAddVerifier(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "AddVerifier", expected)
	require.NoError(t, err)
	runTest(t, verifiedregistry.AddVerifier, tests)
}

func TestRemoveVerifier(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "RemoveVerifier", expected)
	require.NoError(t, err)
	runTest(t, verifiedregistry.RemoveVerifier, tests)
}

func TestAddVerifiedClient(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "AddVerifiedClient", expected)
	require.NoError(t, err)
	runTest(t, verifiedregistry.AddVerifiedClient, tests)
}

func TestUseBytes(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "UseBytes", expected)
	require.NoError(t, err)
	runTest(t, verifiedregistry.UseBytes, tests)
}

func TestRestoreBytes(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "RestoreBytes", expected)
	require.NoError(t, err)
	runTest(t, verifiedregistry.RestoreBytes, tests)
}

func TestRemoveVerifiedClientDataCap(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "RemoveVerifiedClientDataCap", expected)
	require.NoError(t, err)
	runTest(t, verifiedregistry.RemoveVerifiedClientDataCap, tests)
}

func TestDeprecated1(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "Deprecated1", expected)
	require.NoError(t, err)
	runTest(t, verifiedregistry.Deprecated1, tests)
}

func TestDeprecated2(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "Deprecated2", expected)
	require.NoError(t, err)
	runTest(t, verifiedregistry.Deprecated2, tests)
}

func TestClaimAllocations(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ClaimAllocations", expected)
	require.NoError(t, err)
	runTest2(t, verifiedregistry.ClaimAllocations, tests)
}

func TestGetClaims(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetClaims", expected)
	require.NoError(t, err)
	runTest2(t, verifiedregistry.GetClaims, tests)
}

func TestExtendClaimTerms(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ExtendClaimTerms", expected)
	require.NoError(t, err)
	runTest2(t, verifiedregistry.ExtendClaimTerms, tests)
}

func TestRemoveExpiredClaims(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "RemoveExpiredClaims", expected)
	require.NoError(t, err)
	runTest2(t, verifiedregistry.RemoveExpiredClaims, tests)
}

func TestVerifregUniversalReceiverHook(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "VerifregUniversalReceiverHook", expected)
	require.NoError(t, err)
	runTest2(t, verifiedregistry.VerifregUniversalReceiverHook, tests)
}
func TestRemoveExpiredAllocations(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "RemoveExpiredAllocations", expected)
	require.NoError(t, err)
	runTest2(t, verifiedregistry.RemoveExpiredAllocations, tests)
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

				result, err := fn(tt.Network, tt.Height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func runTest2(t *testing.T, fn testFn2, tests []tools.TestCase[map[string]any]) {
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
