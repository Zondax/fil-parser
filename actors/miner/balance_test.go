package miner_test

import (
	"testing"

	"github.com/zondax/fil-parser/actors/miner"
)

func TestGetAvailableBalance(t *testing.T) {
	tests := []test{}
	runTest(t, miner.GetAvailableBalance, tests)
}

func TestGetVestingFunds(t *testing.T) {
	tests := []test{}
	runTest(t, miner.GetVestingFunds, tests)
}

func TestParseWithdrawBalance(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ParseWithdrawBalance, tests)
}
