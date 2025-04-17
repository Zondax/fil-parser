package reward

import (
	rewardv10 "github.com/filecoin-project/go-state-types/builtin/v10/reward"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv12 "github.com/filecoin-project/go-state-types/builtin/v12/reward"
	rewardv13 "github.com/filecoin-project/go-state-types/builtin/v13/reward"
	rewardv14 "github.com/filecoin-project/go-state-types/builtin/v14/reward"
	rewardv15 "github.com/filecoin-project/go-state-types/builtin/v15/reward"
	rewardv16 "github.com/filecoin-project/go-state-types/builtin/v16/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	rewardv9 "github.com/filecoin-project/go-state-types/builtin/v9/reward"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/reward"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/reward"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/reward"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/reward"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/reward"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/reward"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

func awardBlockRewardParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.AwardBlockRewardParams{},

		tools.V8.String(): &legacyv2.AwardBlockRewardParams{},
		tools.V9.String(): &legacyv2.AwardBlockRewardParams{},

		tools.V10.String(): &legacyv3.AwardBlockRewardParams{},
		tools.V11.String(): &legacyv3.AwardBlockRewardParams{},

		tools.V12.String(): &legacyv4.AwardBlockRewardParams{},
		tools.V13.String(): &legacyv5.AwardBlockRewardParams{},
		tools.V14.String(): &legacyv6.AwardBlockRewardParams{},
		tools.V15.String(): &legacyv7.AwardBlockRewardParams{},
		tools.V16.String(): &rewardv8.AwardBlockRewardParams{},
		tools.V17.String(): &rewardv9.AwardBlockRewardParams{},
		tools.V18.String(): &rewardv10.AwardBlockRewardParams{},

		tools.V19.String(): &rewardv11.AwardBlockRewardParams{},
		tools.V20.String(): &rewardv11.AwardBlockRewardParams{},

		tools.V21.String(): &rewardv12.AwardBlockRewardParams{},
		tools.V22.String(): &rewardv13.AwardBlockRewardParams{},
		tools.V23.String(): &rewardv14.AwardBlockRewardParams{},
		tools.V24.String(): &rewardv15.AwardBlockRewardParams{},
		tools.V25.String(): &rewardv16.AwardBlockRewardParams{},
	}
}

func thisEpochRewardReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ThisEpochRewardReturn{},

		tools.V8.String(): &legacyv2.ThisEpochRewardReturn{},
		tools.V9.String(): &legacyv2.ThisEpochRewardReturn{},

		tools.V10.String(): &legacyv3.ThisEpochRewardReturn{},
		tools.V11.String(): &legacyv3.ThisEpochRewardReturn{},

		tools.V12.String(): &legacyv4.ThisEpochRewardReturn{},
		tools.V13.String(): &legacyv5.ThisEpochRewardReturn{},
		tools.V14.String(): &legacyv6.ThisEpochRewardReturn{},
		tools.V15.String(): &legacyv7.ThisEpochRewardReturn{},
		tools.V16.String(): &rewardv8.ThisEpochRewardReturn{},
		tools.V17.String(): &rewardv9.ThisEpochRewardReturn{},
		tools.V18.String(): &rewardv10.ThisEpochRewardReturn{},

		tools.V19.String(): &rewardv11.ThisEpochRewardReturn{},
		tools.V20.String(): &rewardv11.ThisEpochRewardReturn{},

		tools.V21.String(): &rewardv12.ThisEpochRewardReturn{},
		tools.V22.String(): &rewardv13.ThisEpochRewardReturn{},
		tools.V23.String(): &rewardv14.ThisEpochRewardReturn{},
		tools.V24.String(): &rewardv15.ThisEpochRewardReturn{},
		tools.V25.String(): &rewardv16.ThisEpochRewardReturn{},
	}
}
