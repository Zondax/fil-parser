package verifiedregistry_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	verifiedregistry "github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

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
	m.Run()
}

type testFn func(network string, height int64, raw []byte) (map[string]interface{}, error)
type testFn2 func(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error)

func TestVerifiedRegistry(t *testing.T) {
	verifiedregistry := &verifiedregistry.VerifiedRegistry{}
	testFns := map[string]testFn{
		"AddVerifier":                 verifiedregistry.AddVerifier,
		"RemoveVerifier":              verifiedregistry.RemoveVerifier,
		"AddVerifiedClient":           verifiedregistry.AddVerifiedClientExported,
		"UseBytes":                    verifiedregistry.UseBytes,
		"RestoreBytes":                verifiedregistry.RestoreBytes,
		"RemoveVerifiedClientDataCap": verifiedregistry.RemoveVerifiedClientDataCap,
		"Deprecated1":                 verifiedregistry.Deprecated1,
		"Deprecated2":                 verifiedregistry.Deprecated2,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runTest(t, fn, tests)
		})
	}
}

func TestVerifiedRegistry2(t *testing.T) {
	verifiedregistry := &verifiedregistry.VerifiedRegistry{}
	testFns := map[string]testFn2{
		"ClaimAllocations":              verifiedregistry.ClaimAllocations,
		"GetClaims":                     verifiedregistry.GetClaimsExported,
		"ExtendClaimTerms":              verifiedregistry.ExtendClaimTermsExported,
		"RemoveExpiredClaims":           verifiedregistry.RemoveExpiredClaimsExported,
		"VerifregUniversalReceiverHook": verifiedregistry.UniversalReceiverHook,
		"RemoveExpiredAllocations":      verifiedregistry.RemoveExpiredAllocationsExported,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runTest2(t, fn, tests)
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
