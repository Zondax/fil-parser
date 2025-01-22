package miner_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/miner"
	"github.com/zondax/fil-parser/tools"
)

func TestGetAvailableBalance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetAvailableBalance", expectedData)
	require.NoError(t, err)
	runTest(t, miner.GetAvailableBalance, tests)
}

func TestGetVestingFunds(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetVestingFunds", expectedData)
	require.NoError(t, err)
	runTest(t, miner.GetVestingFunds, tests)
}

func TestParseWithdrawBalance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ParseWithdrawBalance", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ParseWithdrawBalance, tests)
}
