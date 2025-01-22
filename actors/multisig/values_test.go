package multisig_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/multisig"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

type valueTestFn func(height int64, txMetadata string) (interface{}, error)

func TestChangeOwnerAddressValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ChangeOwnerAddressValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ChangeOwnerAddressValue, tests)
}

func TestParseWithdrawBalanceValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseWithdrawBalanceValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseWithdrawBalanceValue, tests)
}

func TestParseInvokeContractValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseInvokeContractValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseInvokeContractValue, tests)
}

func TestParseAddSignerValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseAddSignerValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseAddSignerValue, tests)
}

func TestParseApproveValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseApproveValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseApproveValue, tests)
}

func TestParseCancelValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseCancelValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseCancelValue, tests)
}
func TestChangeNumApprovalsThresholdValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ChangeNumApprovalsThresholdValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ChangeNumApprovalsThresholdValue, tests)
}

func TestParseConstructorValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseConstructorValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseConstructorValue, tests)
}

func TestParseLockBalanceValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseLockBalanceValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseLockBalanceValue, tests)
}

func TestParseRemoveSignerValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseRemoveSignerValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseRemoveSignerValue, tests)
}

func TestParseSendValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseSendValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseSendValue, tests)
}

func TestParseSwapSignerValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseSwapSignerValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseSwapSignerValue, tests)
}

func TestParseUniversalReceiverHookValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseUniversalReceiverHookValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.ParseUniversalReceiverHookValue, tests)
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

				result, err := fn(tt.Height, string(trace.Msg.Params))
				require.NoError(t, err)
				require.True(t, reflect.DeepEqual(result, tt.Expected))
			}
		})
	}
}
