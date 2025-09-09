package cron

import (
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	cronv10 "github.com/filecoin-project/go-state-types/builtin/v10/cron"
	cronv11 "github.com/filecoin-project/go-state-types/builtin/v11/cron"
	cronv12 "github.com/filecoin-project/go-state-types/builtin/v12/cron"
	cronv13 "github.com/filecoin-project/go-state-types/builtin/v13/cron"
	cronv14 "github.com/filecoin-project/go-state-types/builtin/v14/cron"
	cronv15 "github.com/filecoin-project/go-state-types/builtin/v15/cron"
	cronv16 "github.com/filecoin-project/go-state-types/builtin/v16/cron"
	cronv17 "github.com/filecoin-project/go-state-types/builtin/v17/cron"
	cronv8 "github.com/filecoin-project/go-state-types/builtin/v8/cron"
	cronv9 "github.com/filecoin-project/go-state-types/builtin/v9/cron"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/cron"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/cron"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/cron"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/cron"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/cron"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/cron"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/cron"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// All methods can be found in the Actor.Exports method in
// the correct version package for "github.com/filecoin-project/specs-actors/actors/builtin/cron"

func v1Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsCron.Constructor: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		legacyBuiltin.MethodsCron.EpochTick: {
			Name:   parser.MethodEpochTick,
			Method: actors.ParseEmptyParamsAndReturn,
		},
	}
}

func v2Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

func v3Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

func v4Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

func v5Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

func v6Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

func v7Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

var cronConstructorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ConstructorParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ConstructorParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ConstructorParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ConstructorParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ConstructorParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ConstructorParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(cronv8.State) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(cronv9.State) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(cronv10.State) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(cronv11.State) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(cronv11.State) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(cronv12.State) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(cronv13.State) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(cronv14.State) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(cronv15.State) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(cronv16.State) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(cronv17.State) },
}
