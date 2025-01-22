package multisig_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/multisig"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
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

type testFn func(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser multisig.ParseFn) (map[string]interface{}, error)

func TestApprove(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "Approve", expected)
	require.NoError(t, err)
	runTest(t, multisig.Approve, tests)
}

func TestCancel(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "Cancel", expected)
	require.NoError(t, err)
	runTest(t, multisig.Cancel, tests)
}

func TestRemoveSigner(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "RemoveSigner", expected)
	require.NoError(t, err)
	runTest(t, multisig.RemoveSigner, tests)
}

func TestChangeNumApprovalsThreshold(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ChangeNumApprovalsThreshold", expected)
	require.NoError(t, err)
	runTest(t, multisig.ChangeNumApprovalsThreshold, tests)
}

func TestLockBalance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "LockBalance", expected)
	require.NoError(t, err)
	runTest(t, multisig.LockBalance, tests)
}

func TestUniversalReceiverHook(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "UniversalReceiverHook", expected)
	require.NoError(t, err)
	runTest(t, multisig.UniversalReceiverHook, tests)
}

func TestMsigConstructor(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "MsigConstructor", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := multisig.MsigConstructor(tt.Network, tt.Height, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}
func TestMsigParams(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "MsigParams", expected)
	require.NoError(t, err)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
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
				result, err := multisig.MsigParams(tt.Network, lotusMsg, tt.Height, tt.TipsetKey, func(*parser.LotusMessage, int64, filTypes.TipSetKey) (string, error) {
					return "", nil
				})
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
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

				lotusMsg := &parser.LotusMessage{
					To:     trace.Msg.To,
					From:   trace.Msg.From,
					Method: trace.Msg.Method,
				}
				result, err := fn(tt.Network, lotusMsg, tt.Height, tt.TipsetKey, trace.MsgRct.Return, func(*parser.LotusMessage, int64, filTypes.TipSetKey) (string, error) {
					return "", nil
				})
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}
