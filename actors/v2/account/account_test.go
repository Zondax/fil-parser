package account_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/actors/v2/account"
	"github.com/zondax/fil-parser/tools"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	accountv10 "github.com/filecoin-project/go-state-types/builtin/v10/account"
	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv12 "github.com/filecoin-project/go-state-types/builtin/v12/account"
	accountv13 "github.com/filecoin-project/go-state-types/builtin/v13/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	accountv15 "github.com/filecoin-project/go-state-types/builtin/v15/account"
	accountv8 "github.com/filecoin-project/go-state-types/builtin/v8/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v9/account"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/account"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/account"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/account"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

var expectedData []byte
var expected map[string]any
var lib *rosettaFilecoinLib.RosettaConstructionFilecoin

var network string

func TestMain(m *testing.M) {
	// var err error
	network = "mainnet"
	// lib, err = tools.GetLib(tools.NodeUrl)
	// if err != nil {
	// 	panic(err)
	// }

	// err = json.Unmarshal(expectedData, &expected)
	// if err != nil {
	// 	panic(err)
	// }
	// expectedData, err = tools.ReadActorSnapshot()
	// if err != nil {
	// 	panic(err)
	// }
	os.Exit(m.Run())
}

func TestAuthenticateMessage(t *testing.T) {

	tests, err := tools.LoadTestData[map[string]any](network, "AuthenticateMessage", expected)
	require.NoError(t, err)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				account := &account.Account{}
				result, err := account.AuthenticateMessage(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				data, err := json.Marshal(result)
				require.NoError(t, err)
				fmt.Println(string(data))
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestPubkeyAddress(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "PubkeyAddress", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				account := &account.Account{}
				result, err := account.PubkeyAddress(tt.Network, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}

		})
	}
}

func TestMethodCoverage(t *testing.T) {
	a := &account.Account{}

	actorVersions := []any{
		legacyv2.Actor{},
		legacyv3.Actor{},
		legacyv4.Actor{},
		legacyv5.Actor{},
		legacyv6.Actor{},
		legacyv7.Actor{},
		accountv8.Methods,
		accountv9.Methods,
		accountv10.Methods,
		accountv11.Methods,
		accountv12.Methods,
		accountv13.Methods,
		accountv14.Methods,
		accountv15.Methods,
	}

	missingMethods := v2.MissingMethods(a, actorVersions)
	assert.Empty(t, missingMethods, "missing methods: %v", missingMethods)
}
