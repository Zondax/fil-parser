package multisig_test

import (
	"encoding/json"
	"reflect"
	"testing"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/v2/multisig"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

type testFn func(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser multisig.ParseFn) (map[string]interface{}, error)
type valueTestFn func(network string, height int64, txMetadata string) (interface{}, error)

var expectedData []byte
var expected map[string]any
var network string

func TestMain(m *testing.M) {
	network = "mainnet"
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		panic(err)
	}
	var err error
	expectedData, err = tools.ReadActorSnapshot()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestMsig(t *testing.T) {
	multisig := &multisig.Msig{}
	testFns := map[string]testFn{
		"Approve":                     multisig.Approve,
		"Cancel":                      multisig.Cancel,
		"RemoveSigner":                multisig.RemoveSigner,
		"ChangeNumApprovalsThreshold": multisig.ChangeNumApprovalsThreshold,
		"LockBalance":                 multisig.LockBalance,
		"UniversalReceiverHook":       multisig.UniversalReceiverHook,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runTest(t, fn, tests)
		})
	}
}

func TestValues(t *testing.T) {
	multisig := &multisig.Msig{}
	testFns := map[string]valueTestFn{
		"ChangeOwnerAddressValue":          multisig.ChangeOwnerAddressValue,
		"ParseWithdrawBalanceValue":        multisig.ParseWithdrawBalanceValue,
		"ParseInvokeContractValue":         multisig.ParseInvokeContractValue,
		"ParseAddSignerValue":              multisig.ParseAddSignerValue,
		"ParseApproveValue":                multisig.ParseApproveValue,
		"ParseCancelValue":                 multisig.ParseCancelValue,
		"ChangeNumApprovalsThresholdValue": multisig.ChangeNumApprovalsThresholdValue,
		"ParseConstructorValue":            multisig.ParseConstructorValue,
		"ParseLockBalanceValue":            multisig.ParseLockBalanceValue,
		"ParseRemoveSignerValue":           multisig.ParseRemoveSignerValue,
		"ParseSendValue":                   multisig.ParseSendValue,
		"ParseSwapSignerValue":             multisig.ParseSwapSignerValue,
		"ParseUniversalReceiverHookValue":  multisig.ParseUniversalReceiverHookValue,
		"AddVerifierValue":                 multisig.AddVerifierValue,
	}

	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runValueTest(t, fn, tests)
		})
	}

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
				multisig := &multisig.Msig{}
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
				multisig := &multisig.Msig{}
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

func runValueTest(t *testing.T, fn valueTestFn, tests []tools.TestCase[map[string]any]) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				// TODO: parse whole multisig tx first before calling this function

				result, err := fn(tt.Network, tt.Height, string(trace.Msg.Params))
				require.NoError(t, err)
				require.True(t, reflect.DeepEqual(result, tt.Expected))
			}
		})
	}
}
