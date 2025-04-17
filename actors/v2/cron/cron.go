package cron

import (
	"fmt"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

func (c *Cron) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	cronConstructor, ok := getCronConstructorParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return cronConstructorLegacy(raw, cronConstructor)

}
