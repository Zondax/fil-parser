package power

import (
	powerv10 "github.com/filecoin-project/go-state-types/builtin/v10/power"
	powerv11 "github.com/filecoin-project/go-state-types/builtin/v11/power"
	powerv12 "github.com/filecoin-project/go-state-types/builtin/v12/power"
	powerv13 "github.com/filecoin-project/go-state-types/builtin/v13/power"
	powerv14 "github.com/filecoin-project/go-state-types/builtin/v14/power"
	powerv15 "github.com/filecoin-project/go-state-types/builtin/v15/power"
	powerv16 "github.com/filecoin-project/go-state-types/builtin/v16/power"
	powerv8 "github.com/filecoin-project/go-state-types/builtin/v8/power"
	powerv9 "github.com/filecoin-project/go-state-types/builtin/v9/power"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/power"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/power"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/power"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/power"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/power"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/power"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/power"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

func currentTotalPowerReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.CurrentTotalPowerReturn{},

		tools.V8.String(): &legacyv2.CurrentTotalPowerReturn{},
		tools.V9.String(): &legacyv2.CurrentTotalPowerReturn{},

		tools.V10.String(): &legacyv3.CurrentTotalPowerReturn{},
		tools.V11.String(): &legacyv3.CurrentTotalPowerReturn{},

		tools.V12.String(): &legacyv4.CurrentTotalPowerReturn{},
		tools.V13.String(): &legacyv5.CurrentTotalPowerReturn{},
		tools.V14.String(): &legacyv6.CurrentTotalPowerReturn{},
		tools.V15.String(): &legacyv7.CurrentTotalPowerReturn{},
		tools.V16.String(): &powerv8.CurrentTotalPowerReturn{},
		tools.V17.String(): &powerv9.CurrentTotalPowerReturn{},
		tools.V18.String(): &powerv10.CurrentTotalPowerReturn{},

		tools.V19.String(): &powerv11.CurrentTotalPowerReturn{},
		tools.V20.String(): &powerv11.CurrentTotalPowerReturn{},

		tools.V21.String(): &powerv12.CurrentTotalPowerReturn{},
		tools.V22.String(): &powerv13.CurrentTotalPowerReturn{},
		tools.V23.String(): &powerv14.CurrentTotalPowerReturn{},
		tools.V24.String(): &powerv15.CurrentTotalPowerReturn{},
		tools.V25.String(): &powerv16.CurrentTotalPowerReturn{},
	}
}

func constructorParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.MinerConstructorParams{},

		tools.V8.String(): &legacyv2.MinerConstructorParams{},
		tools.V9.String(): &legacyv2.MinerConstructorParams{},

		tools.V10.String(): &legacyv3.MinerConstructorParams{},
		tools.V11.String(): &legacyv3.MinerConstructorParams{},

		tools.V12.String(): &legacyv4.MinerConstructorParams{},
		tools.V13.String(): &legacyv5.MinerConstructorParams{},
		tools.V14.String(): &legacyv6.MinerConstructorParams{},
		tools.V15.String(): &legacyv7.MinerConstructorParams{},
		tools.V16.String(): &powerv8.MinerConstructorParams{},
		tools.V17.String(): &powerv9.MinerConstructorParams{},
		tools.V18.String(): &powerv10.MinerConstructorParams{},

		tools.V19.String(): &powerv11.MinerConstructorParams{},
		tools.V20.String(): &powerv11.MinerConstructorParams{},

		tools.V21.String(): &powerv12.MinerConstructorParams{},
		tools.V22.String(): &powerv13.MinerConstructorParams{},
		tools.V23.String(): &powerv14.MinerConstructorParams{},
		tools.V24.String(): &powerv15.MinerConstructorParams{},
		tools.V25.String(): &powerv16.MinerConstructorParams{},
	}
}

func createMinerParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.CreateMinerParams{},

		tools.V8.String(): &legacyv2.CreateMinerParams{},
		tools.V9.String(): &legacyv2.CreateMinerParams{},

		tools.V10.String(): &legacyv3.CreateMinerParams{},
		tools.V11.String(): &legacyv3.CreateMinerParams{},

		tools.V12.String(): &legacyv4.CreateMinerParams{},
		tools.V13.String(): &legacyv5.CreateMinerParams{},
		tools.V14.String(): &legacyv6.CreateMinerParams{},
		tools.V15.String(): &legacyv7.CreateMinerParams{},
		tools.V16.String(): &powerv8.CreateMinerParams{},
		tools.V17.String(): &powerv9.CreateMinerParams{},
		tools.V18.String(): &powerv10.CreateMinerParams{},

		tools.V19.String(): &powerv11.CreateMinerParams{},
		tools.V20.String(): &powerv11.CreateMinerParams{},

		tools.V21.String(): &powerv12.CreateMinerParams{},
		tools.V22.String(): &powerv13.CreateMinerParams{},
		tools.V23.String(): &powerv14.CreateMinerParams{},
		tools.V24.String(): &powerv15.CreateMinerParams{},
		tools.V25.String(): &powerv16.CreateMinerParams{},
	}
}

func createMinerReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.CreateMinerReturn{},

		tools.V8.String(): &legacyv2.CreateMinerReturn{},
		tools.V9.String(): &legacyv2.CreateMinerReturn{},

		tools.V10.String(): &legacyv3.CreateMinerReturn{},
		tools.V11.String(): &legacyv3.CreateMinerReturn{},

		tools.V12.String(): &legacyv4.CreateMinerReturn{},
		tools.V13.String(): &legacyv5.CreateMinerReturn{},
		tools.V14.String(): &legacyv6.CreateMinerReturn{},
		tools.V15.String(): &legacyv7.CreateMinerReturn{},
		tools.V16.String(): &powerv8.CreateMinerReturn{},
		tools.V17.String(): &powerv9.CreateMinerReturn{},
		tools.V18.String(): &powerv10.CreateMinerReturn{},

		tools.V19.String(): &powerv11.CreateMinerReturn{},
		tools.V20.String(): &powerv11.CreateMinerReturn{},

		tools.V21.String(): &powerv12.CreateMinerReturn{},
		tools.V22.String(): &powerv13.CreateMinerReturn{},
		tools.V23.String(): &powerv14.CreateMinerReturn{},
		tools.V24.String(): &powerv15.CreateMinerReturn{},
		tools.V25.String(): &powerv16.CreateMinerReturn{},
	}
}

func enrollCronEventParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.EnrollCronEventParams{},

		tools.V8.String(): &legacyv2.EnrollCronEventParams{},
		tools.V9.String(): &legacyv2.EnrollCronEventParams{},

		tools.V10.String(): &legacyv3.EnrollCronEventParams{},
		tools.V11.String(): &legacyv3.EnrollCronEventParams{},

		tools.V12.String(): &legacyv4.EnrollCronEventParams{},
		tools.V13.String(): &legacyv5.EnrollCronEventParams{},
		tools.V14.String(): &legacyv6.EnrollCronEventParams{},
		tools.V15.String(): &legacyv7.EnrollCronEventParams{},
		tools.V16.String(): &powerv8.EnrollCronEventParams{},
		tools.V17.String(): &powerv9.EnrollCronEventParams{},
		tools.V18.String(): &powerv10.EnrollCronEventParams{},

		tools.V19.String(): &powerv11.EnrollCronEventParams{},
		tools.V20.String(): &powerv11.EnrollCronEventParams{},

		tools.V21.String(): &powerv12.EnrollCronEventParams{},
		tools.V22.String(): &powerv13.EnrollCronEventParams{},
		tools.V23.String(): &powerv14.EnrollCronEventParams{},
		tools.V24.String(): &powerv15.EnrollCronEventParams{},
		tools.V25.String(): &powerv16.EnrollCronEventParams{},
	}
}

func updateClaimedPowerParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.UpdateClaimedPowerParams{},

		tools.V8.String(): &legacyv2.UpdateClaimedPowerParams{},
		tools.V9.String(): &legacyv2.UpdateClaimedPowerParams{},

		tools.V10.String(): &legacyv3.UpdateClaimedPowerParams{},
		tools.V11.String(): &legacyv3.UpdateClaimedPowerParams{},

		tools.V12.String(): &legacyv4.UpdateClaimedPowerParams{},
		tools.V13.String(): &legacyv5.UpdateClaimedPowerParams{},
		tools.V14.String(): &legacyv6.UpdateClaimedPowerParams{},
		tools.V15.String(): &legacyv7.UpdateClaimedPowerParams{},
		tools.V16.String(): &powerv8.UpdateClaimedPowerParams{},
		tools.V17.String(): &powerv9.UpdateClaimedPowerParams{},
		tools.V18.String(): &powerv10.UpdateClaimedPowerParams{},

		tools.V19.String(): &powerv11.UpdateClaimedPowerParams{},
		tools.V20.String(): &powerv11.UpdateClaimedPowerParams{},

		tools.V21.String(): &powerv12.UpdateClaimedPowerParams{},
		tools.V22.String(): &powerv13.UpdateClaimedPowerParams{},
		tools.V23.String(): &powerv14.UpdateClaimedPowerParams{},
		tools.V24.String(): &powerv15.UpdateClaimedPowerParams{},
		tools.V25.String(): &powerv16.UpdateClaimedPowerParams{},
	}
}

func networkRawPowerReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &powerv10.NetworkRawPowerReturn{},

		tools.V19.String(): &powerv11.NetworkRawPowerReturn{},
		tools.V20.String(): &powerv11.NetworkRawPowerReturn{},

		tools.V21.String(): &powerv12.NetworkRawPowerReturn{},
		tools.V22.String(): &powerv13.NetworkRawPowerReturn{},
		tools.V23.String(): &powerv14.NetworkRawPowerReturn{},
		tools.V24.String(): &powerv15.NetworkRawPowerReturn{},
		tools.V25.String(): &powerv16.NetworkRawPowerReturn{},
	}
}

func minerRawPowerParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(powerv10.MinerRawPowerParams),

		tools.V19.String(): new(powerv11.MinerRawPowerParams),
		tools.V20.String(): new(powerv11.MinerRawPowerParams),

		tools.V21.String(): new(powerv12.MinerRawPowerParams),
		tools.V22.String(): new(powerv13.MinerRawPowerParams),
		tools.V23.String(): new(powerv14.MinerRawPowerParams),
		tools.V24.String(): new(powerv15.MinerRawPowerParams),
		tools.V25.String(): new(powerv16.MinerRawPowerParams),
	}
}

func minerRawPowerReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &powerv10.MinerRawPowerReturn{},

		tools.V19.String(): &powerv11.MinerRawPowerReturn{},
		tools.V20.String(): &powerv11.MinerRawPowerReturn{},

		tools.V21.String(): &powerv12.MinerRawPowerReturn{},
		tools.V22.String(): &powerv13.MinerRawPowerReturn{},
		tools.V23.String(): &powerv14.MinerRawPowerReturn{},
		tools.V24.String(): &powerv15.MinerRawPowerReturn{},
		tools.V25.String(): &powerv16.MinerRawPowerReturn{},
	}
}

func minerCountReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(powerv10.MinerCountReturn),

		tools.V19.String(): new(powerv11.MinerCountReturn),
		tools.V20.String(): new(powerv11.MinerCountReturn),

		tools.V21.String(): new(powerv12.MinerCountReturn),
		tools.V22.String(): new(powerv13.MinerCountReturn),
		tools.V23.String(): new(powerv14.MinerCountReturn),
		tools.V24.String(): new(powerv15.MinerCountReturn),
		tools.V25.String(): new(powerv16.MinerCountReturn),
	}
}

func minerConsensusCountReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(powerv10.MinerConsensusCountReturn),

		tools.V19.String(): new(powerv11.MinerConsensusCountReturn),
		tools.V20.String(): new(powerv11.MinerConsensusCountReturn),

		tools.V21.String(): new(powerv12.MinerConsensusCountReturn),
		tools.V22.String(): new(powerv13.MinerConsensusCountReturn),
		tools.V23.String(): new(powerv14.MinerConsensusCountReturn),
		tools.V24.String(): new(powerv15.MinerConsensusCountReturn),
		tools.V25.String(): new(powerv16.MinerConsensusCountReturn),
	}
}
