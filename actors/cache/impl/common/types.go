package common

import (
	"time"

	"github.com/filecoin-project/lotus/api"
	"github.com/zondax/golem/pkg/zcache"
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
	Ttl            time.Duration
	LatestCacheTTL time.Duration
}

// Copy returns a deep copy of the config. It is nil-safe: a nil config (or nil
// embedded/sub configs) are preserved as nil instead of panicking.
func (c *CacheConfig) Copy() *CacheConfig {
	if c == nil {
		return nil
	}

	cp := &CacheConfig{
		Ttl:            c.Ttl,
		LatestCacheTTL: c.LatestCacheTTL,
	}
	if c.CombinedConfig == nil {
		return cp
	}

	combined := *c.CombinedConfig
	if c.Local != nil {
		local := *c.Local
		combined.Local = &local
	}
	if c.Remote != nil {
		remote := *c.Remote
		combined.Remote = &remote
	}
	cp.CombinedConfig = &combined

	return cp
}
