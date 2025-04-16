package cron

import (
	"fmt"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/cron"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/cron"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/cron"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/cron"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/cron"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/cron"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/cron"

	cronv10 "github.com/filecoin-project/go-state-types/builtin/v10/cron"
	cronv11 "github.com/filecoin-project/go-state-types/builtin/v11/cron"
	cronv12 "github.com/filecoin-project/go-state-types/builtin/v12/cron"
	cronv13 "github.com/filecoin-project/go-state-types/builtin/v13/cron"
	cronv14 "github.com/filecoin-project/go-state-types/builtin/v14/cron"
	cronv15 "github.com/filecoin-project/go-state-types/builtin/v15/cron"
	cronv16 "github.com/filecoin-project/go-state-types/builtin/v16/cron"
	cronv8 "github.com/filecoin-project/go-state-types/builtin/v8/cron"
	cronv9 "github.com/filecoin-project/go-state-types/builtin/v9/cron"

	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

func getCronConstructorParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V1.String(): &legacyv1.ConstructorParams{},

		tools.V8.String(): &legacyv2.ConstructorParams{},
		tools.V9.String(): &legacyv2.ConstructorParams{},

		tools.V10.String(): &legacyv3.ConstructorParams{},
		tools.V11.String(): &legacyv3.ConstructorParams{},

		tools.V12.String(): &legacyv4.ConstructorParams{},
		tools.V13.String(): &legacyv5.ConstructorParams{},
		tools.V14.String(): &legacyv6.ConstructorParams{},
		tools.V15.String(): &legacyv7.ConstructorParams{},
		tools.V16.String(): &cronv8.State{},
		tools.V17.String(): &cronv9.State{},
		tools.V18.String(): &cronv10.State{},

		tools.V19.String(): &cronv11.State{},
		tools.V20.String(): &cronv11.State{},

		tools.V21.String(): &cronv12.State{},
		tools.V22.String(): &cronv13.State{},
		tools.V23.String(): &cronv14.State{},
		tools.V24.String(): &cronv15.State{},
		tools.V25.String(): &cronv16.State{},
	}
}

func (c *Cron) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	cronConstructor, ok := getCronConstructorParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return cronConstructorLegacy(raw, cronConstructor)

}
