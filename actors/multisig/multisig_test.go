package multisig_test

import (
	"testing"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/multisig"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

type testFn func(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser multisig.ParseFn) (map[string]interface{}, error)
type tests struct {
	name      string
	height    int64
	version   string
	expected  map[string]interface{}
	tipsetKey filTypes.TipSetKey
}

func TestApprove(t *testing.T) {
	tests := []tests{}
	runTest(t, multisig.Approve, tests)
}

func TestCancel(t *testing.T) {
	tests := []tests{}
	runTest(t, multisig.Cancel, tests)
}

func TestRemoveSigner(t *testing.T) {
	tests := []tests{}
	runTest(t, multisig.RemoveSigner, tests)
}

func TestChangeNumApprovalsThreshold(t *testing.T) {
	tests := []tests{}
	runTest(t, multisig.ChangeNumApprovalsThreshold, tests)
}

func TestLockBalance(t *testing.T) {
	tests := []tests{}
	runTest(t, multisig.LockBalance, tests)
}

func TestUniversalReceiverHook(t *testing.T) {
	tests := []tests{}
	runTest(t, multisig.UniversalReceiverHook, tests)
}

func TestMsigConstructor(t *testing.T) {
	tests := []tests{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := multisig.MsigConstructor(tt.height, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))

			}
		})
	}
}
func TestMsigParams(t *testing.T) {
	tests := []tests{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				lotusMsg := &parser.LotusMessage{
					To:     trace.Msg.To,
					From:   trace.Msg.From,
					Method: trace.Msg.Method,
				}
				result, err := multisig.MsigParams(lotusMsg, tt.height, tt.tipsetKey, func(*parser.LotusMessage, int64, filTypes.TipSetKey) (string, error) {
					return "", nil
				})
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))

			}
		})
	}
}

func runTest(t *testing.T, fn testFn, tests []tests) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				lotusMsg := &parser.LotusMessage{
					To:     trace.Msg.To,
					From:   trace.Msg.From,
					Method: trace.Msg.Method,
				}
				result, err := fn(lotusMsg, tt.height, tt.tipsetKey, trace.MsgRct.Return, func(*parser.LotusMessage, int64, filTypes.TipSetKey) (string, error) {
					return "", nil
				})
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))

			}
		})
	}
}
