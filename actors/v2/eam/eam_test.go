package eam_test

import (
	"os"
	"testing"

	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/actors/v2/eam"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

var network string

var expectedData []byte
var expected map[string]any

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

type testFn func(network string, height int64, raw, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error)

func TestEam(t *testing.T) {
	eam := &eam.Eam{}
	testFns := map[string]testFn{
		"CreateExternal": eam.CreateExternal,
		"Create":         eam.Create,
		"Create2":        eam.Create2,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runTest(t, fn, tests)
		})
	}
}

func TestMethodCoverage(t *testing.T) {
	e := &eam.Eam{}

	actorVersions := []any{
		eamv10.Methods,
		eamv11.Methods,
		eamv12.Methods,
		eamv13.Methods,
		eamv14.Methods,
		eamv15.Methods,
	}

	missingMethods := v2.MissingMethods(e, actorVersions)
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
				result, addrInfo, err := fn(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return, trace.Msg.Cid())
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
				require.Equal(t, addrInfo, tt.Address)
			}
		})
	}
}
