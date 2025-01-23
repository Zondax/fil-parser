package account_test

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/account"
	"github.com/zondax/fil-parser/tools"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	// typesV1 "github.com/zondax/fil-parser/parser/v1/types"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

//go:embed expected.json
var expectedData []byte
var expected map[string]any
var lib *rosettaFilecoinLib.RosettaConstructionFilecoin

var network string

func TestMain(m *testing.M) {
	var err error
	network = "mainnet"
	lib, err = tools.GetLib(tools.NodeUrl)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(expectedData, &expected)
	if err != nil {
		panic(err)
	}
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
