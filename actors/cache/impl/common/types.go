package common

import (
	"github.com/filecoin-project/lotus/api"
	"github.com/zondax/golem/pkg/zcache"
	"time"
)

type DataSourceConfig struct {
	Cache       *CacheConfig
	NetworkName string
}

type DataSource struct {
	Node   api.FullNode
	Config DataSourceConfig
}

type CacheConfig struct {
	*zcache.CombinedConfig
	Ttl time.Duration
}
