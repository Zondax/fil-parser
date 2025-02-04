package verifiedregistry_test

import (
	"testing"

	verifregv10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifregv11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifregv12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifregv13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifregv14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifregv15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifregv8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifregv9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/verifreg"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/verifreg"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/verifreg"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/verifreg"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/verifreg"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/verifreg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v2 "github.com/zondax/fil-parser/actors/v2"
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

func TestMethodCoverage(t *testing.T) {
	vr := &verifiedregistry.VerifiedRegistry{}

	actorVersions := []any{
		legacyv2.Actor{},
		legacyv3.Actor{},
		legacyv4.Actor{},
		legacyv5.Actor{},
		legacyv6.Actor{},
		legacyv7.Actor{},
		verifregv8.Methods,
		verifregv9.Methods,
		verifregv10.Methods,
		verifregv11.Methods,
		verifregv12.Methods,
		verifregv13.Methods,
		verifregv14.Methods,
		verifregv15.Methods,
	}

	missingMethods := v2.MissingMethods(vr, actorVersions)
	assert.Empty(t, missingMethods, "missing methods: %v", missingMethods)
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
