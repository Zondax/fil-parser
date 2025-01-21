package cron

import (
	cronv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/cron"
	cronv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/cron"
	cronv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/cron"
	cronv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/cron"
	cronv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/cron"
	cronv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/cron"
	cronv8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/cron"
)

// TODO: update to correct height ranges
func CronConstructor(height uint64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return cronConstructorGeneric[*cronv8.ConstructorParams](raw, &cronv8.ConstructorParams{})
	case 7:
		return cronConstructorGeneric[*cronv7.ConstructorParams](raw, &cronv7.ConstructorParams{})
	case 6:
		return cronConstructorGeneric[*cronv6.ConstructorParams](raw, &cronv6.ConstructorParams{})
	case 5:
		return cronConstructorGeneric[*cronv5.ConstructorParams](raw, &cronv5.ConstructorParams{})
	case 4:
		return cronConstructorGeneric[*cronv4.ConstructorParams](raw, &cronv4.ConstructorParams{})
	case 3:
		return cronConstructorGeneric[*cronv3.ConstructorParams](raw, &cronv3.ConstructorParams{})
	case 2:
		return cronConstructorGeneric[*cronv2.ConstructorParams](raw, &cronv2.ConstructorParams{})
	}
	return nil, nil
}
