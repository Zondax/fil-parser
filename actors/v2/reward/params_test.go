package reward_test

import (
	"testing"

	"github.com/filecoin-project/go-address"
	rewardv10 "github.com/filecoin-project/go-state-types/builtin/v10/reward"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv12 "github.com/filecoin-project/go-state-types/builtin/v12/reward"
	rewardv13 "github.com/filecoin-project/go-state-types/builtin/v13/reward"
	rewardv14 "github.com/filecoin-project/go-state-types/builtin/v14/reward"
	rewardv15 "github.com/filecoin-project/go-state-types/builtin/v15/reward"
	rewardv16 "github.com/filecoin-project/go-state-types/builtin/v16/reward"
	rewardv17 "github.com/filecoin-project/go-state-types/builtin/v17/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	rewardv9 "github.com/filecoin-project/go-state-types/builtin/v9/reward"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/reward"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/reward"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/reward"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/reward"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/reward"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/reward"
	"github.com/zondax/fil-parser/actors/v2/reward"
)

func TestGetMinerFromAwardBlockRewardParams(t *testing.T) {
	addr, err := address.NewIDAddress(1)
	if err != nil {
		t.Fatal(err)
	}
	addrStr := addr.String()
	tests := []struct {
		name   string
		params any
		want   string
	}{
		{name: "legacyv1", params: legacyv1.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "legacyv2", params: legacyv2.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "legacyv3", params: legacyv3.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "legacyv4", params: legacyv4.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "legacyv5", params: legacyv5.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "legacyv6", params: legacyv6.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "legacyv7", params: legacyv7.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "v8", params: rewardv8.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "v9", params: rewardv9.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "v10", params: rewardv10.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "v11", params: rewardv11.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "v12", params: rewardv12.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "v13", params: rewardv13.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "v14", params: rewardv14.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "v15", params: rewardv15.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "v16", params: rewardv16.AwardBlockRewardParams{Miner: addr}, want: addrStr},

		{name: "*legacyv1", params: &legacyv1.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*legacyv2", params: &legacyv2.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*legacyv3", params: &legacyv3.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*legacyv4", params: &legacyv4.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*legacyv5", params: &legacyv5.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*legacyv6", params: &legacyv6.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*legacyv7", params: &legacyv7.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v8", params: &rewardv8.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v9", params: &rewardv9.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v10", params: &rewardv10.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v11", params: &rewardv11.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v12", params: &rewardv12.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v13", params: &rewardv13.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v14", params: &rewardv14.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v15", params: &rewardv15.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v16", params: &rewardv16.AwardBlockRewardParams{Miner: addr}, want: addrStr},
		{name: "*v17", params: &rewardv17.AwardBlockRewardParams{Miner: addr}, want: addrStr},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := reward.GetMinerFromAwardBlockRewardParams(test.params)
			if got != test.want {
				t.Errorf("GetMinerFromAwardBlockRewardParams(%v) = %s; want %s", test.params, got, test.want)
			}
		})
	}
}
