package power

import (
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	powerv10 "github.com/filecoin-project/go-state-types/builtin/v10/power"
	powerv11 "github.com/filecoin-project/go-state-types/builtin/v11/power"
	powerv12 "github.com/filecoin-project/go-state-types/builtin/v12/power"
	powerv13 "github.com/filecoin-project/go-state-types/builtin/v13/power"
	powerv14 "github.com/filecoin-project/go-state-types/builtin/v14/power"
	powerv15 "github.com/filecoin-project/go-state-types/builtin/v15/power"
	powerv16 "github.com/filecoin-project/go-state-types/builtin/v16/power"
	powerv17 "github.com/filecoin-project/go-state-types/builtin/v17/power"
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
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// All methods can be found in the Actor.Exports method in
// the correct version package for "github.com/filecoin-project/specs-actors/actors/builtin/power"

func v1Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	p := &Power{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		1: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		2: {
			Name:   parser.MethodCreateMiner,
			Method: p.CreateMinerExported,
		},
		3: {
			Name:   parser.MethodUpdateClaimedPower,
			Method: p.UpdateClaimedPower,
		},
		4: {
			Name:   parser.MethodEnrollCronEvent,
			Method: p.EnrollCronEvent,
		},
		5: {
			Name:   parser.MethodOnEpochTickEnd,
			Method: actors.ParseEmptyParamsAndReturn,
		},
		6: {
			Name:   parser.MethodUpdatePledgeTotal,
			Method: p.UpdatePledgeTotal,
		},
		7: {
			Name:   parser.MethodOnConsensusFault,
			Method: p.OnConsensusFault,
		},
		8: {
			Name:   parser.MethodSubmitPoRepForBulkVerify,
			Method: p.SubmitPoRepForBulkVerify,
		},
		9: {
			Name:   parser.MethodCurrentTotalPower,
			Method: p.CurrentTotalPower,
		},
	}
}
func v2Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	methods := v1Methods()
	return methods
}
func v3Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	methods := v2Methods()
	return methods
}

func v4Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	methods := v3Methods()
	return methods
}

func v5Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	methods := v4Methods()
	return methods
}

func v6Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	methods := v5Methods()
	return methods
}
func v7Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	methods := v6Methods()
	return methods
}

var currentTotalPowerReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CurrentTotalPowerReturn) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CurrentTotalPowerReturn) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CurrentTotalPowerReturn) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CurrentTotalPowerReturn) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CurrentTotalPowerReturn) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CurrentTotalPowerReturn) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CurrentTotalPowerReturn) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CurrentTotalPowerReturn) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CurrentTotalPowerReturn) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CurrentTotalPowerReturn) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CurrentTotalPowerReturn) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CurrentTotalPowerReturn) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CurrentTotalPowerReturn) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CurrentTotalPowerReturn) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.CurrentTotalPowerReturn) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.CurrentTotalPowerReturn) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(powerv8.CurrentTotalPowerReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(powerv9.CurrentTotalPowerReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.CurrentTotalPowerReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.CurrentTotalPowerReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.CurrentTotalPowerReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.CurrentTotalPowerReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.CurrentTotalPowerReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.CurrentTotalPowerReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.CurrentTotalPowerReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.CurrentTotalPowerReturn) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.CurrentTotalPowerReturn) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.CurrentTotalPowerReturn) },
}

var constructorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.MinerConstructorParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.MinerConstructorParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.MinerConstructorParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.MinerConstructorParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.MinerConstructorParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.MinerConstructorParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.MinerConstructorParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.MinerConstructorParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.MinerConstructorParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.MinerConstructorParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.MinerConstructorParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.MinerConstructorParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.MinerConstructorParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.MinerConstructorParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.MinerConstructorParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.MinerConstructorParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(powerv8.MinerConstructorParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(powerv9.MinerConstructorParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.MinerConstructorParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerConstructorParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerConstructorParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.MinerConstructorParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.MinerConstructorParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.MinerConstructorParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.MinerConstructorParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerConstructorParams) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerConstructorParams) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.MinerConstructorParams) },
}

var createMinerParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CreateMinerParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CreateMinerParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CreateMinerParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CreateMinerParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CreateMinerParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CreateMinerParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CreateMinerParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CreateMinerParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.CreateMinerParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.CreateMinerParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(powerv8.CreateMinerParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(powerv9.CreateMinerParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.CreateMinerParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.CreateMinerParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.CreateMinerParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.CreateMinerParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.CreateMinerParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.CreateMinerParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.CreateMinerParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.CreateMinerParams) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.CreateMinerParams) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.CreateMinerParams) },
}

var createMinerReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CreateMinerReturn) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CreateMinerReturn) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CreateMinerReturn) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CreateMinerReturn) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerReturn) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerReturn) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerReturn) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerReturn) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerReturn) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CreateMinerReturn) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CreateMinerReturn) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CreateMinerReturn) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CreateMinerReturn) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CreateMinerReturn) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.CreateMinerReturn) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.CreateMinerReturn) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(powerv8.CreateMinerReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(powerv9.CreateMinerReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.CreateMinerReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.CreateMinerReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.CreateMinerReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.CreateMinerReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.CreateMinerReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.CreateMinerReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.CreateMinerReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.CreateMinerReturn) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.CreateMinerReturn) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.CreateMinerReturn) },
}

var enrollCronEventParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.EnrollCronEventParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.EnrollCronEventParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.EnrollCronEventParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.EnrollCronEventParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.EnrollCronEventParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.EnrollCronEventParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.EnrollCronEventParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.EnrollCronEventParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.EnrollCronEventParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.EnrollCronEventParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.EnrollCronEventParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.EnrollCronEventParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.EnrollCronEventParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.EnrollCronEventParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.EnrollCronEventParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.EnrollCronEventParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(powerv8.EnrollCronEventParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(powerv9.EnrollCronEventParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.EnrollCronEventParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.EnrollCronEventParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.EnrollCronEventParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.EnrollCronEventParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.EnrollCronEventParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.EnrollCronEventParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.EnrollCronEventParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.EnrollCronEventParams) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.EnrollCronEventParams) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.EnrollCronEventParams) },
}

var updateClaimedPowerParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateClaimedPowerParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateClaimedPowerParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateClaimedPowerParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.UpdateClaimedPowerParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UpdateClaimedPowerParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UpdateClaimedPowerParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UpdateClaimedPowerParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UpdateClaimedPowerParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UpdateClaimedPowerParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.UpdateClaimedPowerParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.UpdateClaimedPowerParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.UpdateClaimedPowerParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.UpdateClaimedPowerParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.UpdateClaimedPowerParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.UpdateClaimedPowerParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.UpdateClaimedPowerParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(powerv8.UpdateClaimedPowerParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(powerv9.UpdateClaimedPowerParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.UpdateClaimedPowerParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.UpdateClaimedPowerParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.UpdateClaimedPowerParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.UpdateClaimedPowerParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.UpdateClaimedPowerParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.UpdateClaimedPowerParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.UpdateClaimedPowerParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.UpdateClaimedPowerParams) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.UpdateClaimedPowerParams) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.UpdateClaimedPowerParams) },
}

var networkRawPowerReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.NetworkRawPowerReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.NetworkRawPowerReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.NetworkRawPowerReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.NetworkRawPowerReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.NetworkRawPowerReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.NetworkRawPowerReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.NetworkRawPowerReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.NetworkRawPowerReturn) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.NetworkRawPowerReturn) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.NetworkRawPowerReturn) },
}

var minerRawPowerParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.MinerRawPowerParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerRawPowerParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerRawPowerParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.MinerRawPowerParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.MinerRawPowerParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.MinerRawPowerParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.MinerRawPowerParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerRawPowerParams) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerRawPowerParams) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.MinerRawPowerParams) },
}

var minerRawPowerReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.MinerRawPowerReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerRawPowerReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerRawPowerReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.MinerRawPowerReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.MinerRawPowerReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.MinerRawPowerReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.MinerRawPowerReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerRawPowerReturn) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerRawPowerReturn) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.MinerRawPowerReturn) },
}

var minerCountReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.MinerCountReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerCountReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerCountReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.MinerCountReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.MinerCountReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.MinerCountReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.MinerCountReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerCountReturn) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerCountReturn) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.MinerCountReturn) },
}

var minerConsensusCountReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(powerv10.MinerConsensusCountReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerConsensusCountReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(powerv11.MinerConsensusCountReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(powerv12.MinerConsensusCountReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(powerv13.MinerConsensusCountReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(powerv14.MinerConsensusCountReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(powerv15.MinerConsensusCountReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerConsensusCountReturn) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(powerv16.MinerConsensusCountReturn) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(powerv17.MinerConsensusCountReturn) },
}
