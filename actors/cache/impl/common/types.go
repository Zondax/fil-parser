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

func (c *CacheConfig) Copy() *CacheConfig {
	return &CacheConfig{
		CombinedConfig: &zcache.CombinedConfig{
			GlobalPrefix:       c.GlobalPrefix,
			IsRemoteBestEffort: c.IsRemoteBestEffort,
			Local: &zcache.LocalConfig{
				Prefix:       c.Local.Prefix,
				Logger:       c.Local.Logger,
				MetricServer: c.Local.MetricServer,
				StatsMetrics: c.Local.StatsMetrics,
				NumCounters:  c.Local.NumCounters,
				MaxCostMB:    c.Local.MaxCostMB,
				BufferItems:  c.Local.BufferItems,
			},
			Remote: &zcache.RemoteConfig{
				Network:            c.Remote.Network,
				Addr:               c.Remote.Addr,
				Password:           c.Remote.Password,
				DB:                 c.Remote.DB,
				DialTimeout:        c.Remote.DialTimeout,
				ReadTimeout:        c.Remote.ReadTimeout,
				WriteTimeout:       c.Remote.WriteTimeout,
				PoolSize:           c.Remote.PoolSize,
				MinIdleConns:       c.Remote.MinIdleConns,
				MaxConnAge:         c.Remote.MaxConnAge,
				PoolTimeout:        c.Remote.PoolTimeout,
				IdleTimeout:        c.Remote.IdleTimeout,
				IdleCheckFrequency: c.Remote.IdleCheckFrequency,
				Prefix:             c.Remote.Prefix,
				Logger:             c.Remote.Logger,
				MetricServer:       c.Remote.MetricServer,
				StatsMetrics:       c.Remote.StatsMetrics,
			},
			GlobalLogger:       c.GlobalLogger,
			GlobalMetricServer: c.GlobalMetricServer,
			GlobalStatsMetrics: c.GlobalStatsMetrics,
		},
		Ttl:            c.Ttl,
		LatestCacheTTL: c.LatestCacheTTL,
	}
}
