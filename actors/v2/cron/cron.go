package cron

import (
	"fmt"

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
	cronv8 "github.com/filecoin-project/go-state-types/builtin/v8/cron"
	cronv9 "github.com/filecoin-project/go-state-types/builtin/v9/cron"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

func (c *Cron) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {

	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V10)...):
		return cronConstructorLegacy(raw, &legacyv2.ConstructorParams{})

	case tools.V11.IsSupported(network, height):
		return cronConstructorLegacy(raw, &legacyv3.ConstructorParams{})
	case tools.V12.IsSupported(network, height):
		return cronConstructorLegacy(raw, &legacyv4.ConstructorParams{})
	case tools.V13.IsSupported(network, height):
		return cronConstructorLegacy(raw, &legacyv5.ConstructorParams{})
	case tools.V14.IsSupported(network, height):
		return cronConstructorLegacy(raw, &legacyv6.ConstructorParams{})
	case tools.V15.IsSupported(network, height):
		return cronConstructorLegacy(raw, &legacyv7.ConstructorParams{})

	case tools.V16.IsSupported(network, height):
		return cronConstructorLegacy(raw, &cronv8.State{})
	case tools.V17.IsSupported(network, height):
		return cronConstructorLegacy(raw, &cronv9.State{})
	case tools.V18.IsSupported(network, height):
		return cronConstructorLegacy(raw, &cronv10.State{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return cronConstructorLegacy(raw, &cronv11.State{})
	case tools.V21.IsSupported(network, height):
		return cronConstructorLegacy(raw, &cronv12.State{})
	case tools.V22.IsSupported(network, height):
		return cronConstructorLegacy(raw, &cronv13.State{})
	case tools.V23.IsSupported(network, height):
		return cronConstructorLegacy(raw, &cronv14.State{})
	case tools.V24.IsSupported(network, height):
		return cronConstructorLegacy(raw, &cronv15.State{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
