package cron_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/actors/v2/cron"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"

	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/cron"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/cron"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/cron"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/cron"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/cron"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/cron"

	cronv10 "github.com/filecoin-project/go-state-types/builtin/v10/cron"
	cronv11 "github.com/filecoin-project/go-state-types/builtin/v11/cron"
	cronv12 "github.com/filecoin-project/go-state-types/builtin/v12/cron"
	cronv13 "github.com/filecoin-project/go-state-types/builtin/v13/cron"
	cronv14 "github.com/filecoin-project/go-state-types/builtin/v14/cron"
	cronv15 "github.com/filecoin-project/go-state-types/builtin/v15/cron"
	cronv8 "github.com/filecoin-project/go-state-types/builtin/v8/cron"
	cronv9 "github.com/filecoin-project/go-state-types/builtin/v9/cron"
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

func TestCron(t *testing.T) {
	cron := &cron.Cron{}
	testFns := map[string]testFn{
		"Constructor": cron.Constructor,
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
	c := &cron.Cron{}

	actorVersions := []any{
		legacyv2.Actor{},
		legacyv3.Actor{},
		legacyv4.Actor{},
		legacyv5.Actor{},
		legacyv6.Actor{},
		legacyv7.Actor{},
		cronv8.Methods,
		cronv9.Methods,
		cronv10.Methods,
		cronv11.Methods,
		cronv12.Methods,
		cronv13.Methods,
		cronv14.Methods,
		cronv15.Methods,
	}

	missingMethods := v2.MissingMethods(c, actorVersions)
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
