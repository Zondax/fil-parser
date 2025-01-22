package verifiedregistry_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	verifiedregistry "github.com/zondax/fil-parser/actors/verifiedRegistry"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

type testFn func(height int64, raw []byte) (map[string]interface{}, error)
type testFn2 func(height int64, raw, rawReturn []byte) (map[string]interface{}, error)
type test struct {
	name     string
	height   int64
	version  string
	expected map[string]interface{}
}

func TestAddVerifier(t *testing.T) {
	tests := []test{}
	runTest(t, verifiedregistry.AddVerifier, tests)
}

func TestRemoveVerifier(t *testing.T) {
	tests := []test{}
	runTest(t, verifiedregistry.RemoveVerifier, tests)
}

func TestAddVerifiedClient(t *testing.T) {
	tests := []test{}
	runTest(t, verifiedregistry.AddVerifiedClient, tests)
}

func TestUseBytes(t *testing.T) {
	tests := []test{}
	runTest(t, verifiedregistry.UseBytes, tests)
}

func TestRestoreBytes(t *testing.T) {
	tests := []test{}
	runTest(t, verifiedregistry.RestoreBytes, tests)
}

func TestRemoveVerifiedClientDataCap(t *testing.T) {
	tests := []test{}
	runTest(t, verifiedregistry.RemoveVerifiedClientDataCap, tests)
}

func TestDeprecated1(t *testing.T) {
	tests := []test{}
	runTest(t, verifiedregistry.Deprecated1, tests)
}

func TestDeprecated2(t *testing.T) {
	tests := []test{}
	runTest(t, verifiedregistry.Deprecated2, tests)
}

func TestClaimAllocations(t *testing.T) {
	tests := []test{}
	runTest2(t, verifiedregistry.ClaimAllocations, tests)
}

func TestGetClaims(t *testing.T) {
	tests := []test{}
	runTest2(t, verifiedregistry.GetClaims, tests)
}

func TestExtendClaimTerms(t *testing.T) {
	tests := []test{}
	runTest2(t, verifiedregistry.ExtendClaimTerms, tests)
}

func TestRemoveExpiredClaims(t *testing.T) {
	tests := []test{}
	runTest2(t, verifiedregistry.RemoveExpiredClaims, tests)
}

func TestVerifregUniversalReceiverHook(t *testing.T) {
	tests := []test{}
	runTest2(t, verifiedregistry.VerifregUniversalReceiverHook, tests)
}
func TestRemoveExpiredAllocations(t *testing.T) {
	tests := []test{}
	runTest2(t, verifiedregistry.RemoveExpiredAllocations, tests)
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

				result, err := fn(tt.height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}

func runTest2(t *testing.T, fn testFn2, tests []test) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := fn(tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}
