package cron

import (
	cronv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/cron"
	cronv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/cron"
	cronv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/cron"
	cronv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/cron"
	cronv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/cron"
	cronv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/cron"
	cronv8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/cron"
	"github.com/zondax/fil-parser/tools"
)

// TODO: update to correct height ranges
func CronConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return cronConstructorGeneric[*cronv8.ConstructorParams](raw, &cronv8.ConstructorParams{})
	case tools.V7.IsSupported(height):
		return cronConstructorGeneric[*cronv7.ConstructorParams](raw, &cronv7.ConstructorParams{})
	case tools.V6.IsSupported(height):
		return cronConstructorGeneric[*cronv6.ConstructorParams](raw, &cronv6.ConstructorParams{})
	case tools.V5.IsSupported(height):
		return cronConstructorGeneric[*cronv5.ConstructorParams](raw, &cronv5.ConstructorParams{})
	case tools.V4.IsSupported(height):
		return cronConstructorGeneric[*cronv4.ConstructorParams](raw, &cronv4.ConstructorParams{})
	case tools.V3.IsSupported(height):
		return cronConstructorGeneric[*cronv3.ConstructorParams](raw, &cronv3.ConstructorParams{})
	case tools.V2.IsSupported(height):
		return cronConstructorGeneric[*cronv2.ConstructorParams](raw, &cronv2.ConstructorParams{})
	}
	return nil, nil
}
