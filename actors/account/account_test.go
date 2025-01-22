package account_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/account"
	"github.com/zondax/fil-parser/tools"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	// typesV1 "github.com/zondax/fil-parser/parser/v1/types"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

var lib *rosettaFilecoinLib.RosettaConstructionFilecoin

func TestMain(m *testing.M) {
	var err error
	lib, err = tools.GetLib(tools.NodeUrl)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestAuthenticateMessage(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		url      string
		height   int64
		expected map[string]any
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := account.AuthenticateMessage(tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}

		})
	}
}

func TestPubkeyAddress(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		url      string
		height   int64
		expected map[string]any
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := account.PubkeyAddress(trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}

		})
	}
}
