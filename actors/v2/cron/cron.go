package cron

import (
	"fmt"

	cronv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/cron"
	cronv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/cron"
	cronv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/cron"
	cronv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/cron"
	cronv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/cron"
	cronv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/cron"
	cronv8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/cron"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

func (c *Cron) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return cronConstructorGeneric(raw, &cronv2.ConstructorParams{})
	case tools.V16.IsSupported(network, height):
		return cronConstructorGeneric(raw, &cronv2.ConstructorParams{})
	case tools.V17.IsSupported(network, height):
		return cronConstructorGeneric(raw, &cronv3.ConstructorParams{})
	case tools.V18.IsSupported(network, height):
		return cronConstructorGeneric(raw, &cronv4.ConstructorParams{})
	case tools.V19.IsSupported(network, height):
		return cronConstructorGeneric(raw, &cronv5.ConstructorParams{})
	case tools.V20.IsSupported(network, height):
		return cronConstructorGeneric(raw, &cronv6.ConstructorParams{})
	case tools.V21.IsSupported(network, height):
		return cronConstructorGeneric(raw, &cronv7.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsAfter(network, tools.V22)...):
		return cronConstructorGeneric(raw, &cronv8.ConstructorParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
