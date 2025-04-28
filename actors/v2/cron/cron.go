package cron

import (
	"fmt"

	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"
	"github.com/zondax/fil-parser/tools"
)

func (c *Cron) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := cronConstructorParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}

	return cronConstructorLegacy(raw, params())
}
