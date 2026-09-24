package impl

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/zondax/fil-parser/actors/cache/impl/common"
	cacheMetrics "github.com/zondax/fil-parser/actors/cache/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	"github.com/zondax/golem/pkg/zcache"
)

// A DataSource without a cache config (e.g. only a node) must not panic and
// must fall back to local-only (in-memory) caches for both canonical and latest.
func TestZCacheBlockConfirmation_NilCacheConfig(t *testing.T) {
	var m ZCacheBlockConfirmation
	metrics := cacheMetrics.NewClient(metrics2.NewNoopMetricsClient(), "test")

	require.NotPanics(t, func() {
		require.NoError(t, m.NewImpl(common.DataSource{}, logger.NewLogger(), metrics, nil))
	})
	require.NotNil(t, m.offChainCanonical)
	require.NotNil(t, m.offChainLatest)
	assert.Equal(t, ZCacheLocalOnly, m.offChainCanonical.cacheType)
	assert.Equal(t, ZCacheLocalOnly, m.offChainLatest.cacheType)

	ctx := context.Background()
	canonicalShort, canonicalRobust := "f01234", "f1abjxfbp274xpdqcpuaykwkfb43omjotacm2p3za"
	latestShort, latestRobust := "f05678", "f17uoq6tp427uzv7fztkbsnn64iwotfrristwpryy"

	m.StoreAddressInfo(types.AddressInfo{Short: canonicalShort, Robust: canonicalRobust, IsCanonical: true})
	m.StoreAddressInfo(types.AddressInfo{Short: latestShort, Robust: latestRobust, IsCanonical: false})

	canonicalAddr, err := address.NewFromString(canonicalRobust)
	require.NoError(t, err)
	latestAddr, err := address.NewFromString(latestRobust)
	require.NoError(t, err)

	// canonical entries are visible to canonical and non-canonical lookups
	short, err := m.GetShortAddress(ctx, canonicalAddr, true)
	require.NoError(t, err)
	assert.Equal(t, canonicalShort, short)

	// latest entries are only visible to non-canonical lookups
	short, err = m.GetShortAddress(ctx, latestAddr, false)
	require.NoError(t, err)
	assert.Equal(t, latestShort, short)
	_, err = m.GetShortAddress(ctx, latestAddr, true)
	assert.Error(t, err)
}

func TestCacheConfigCopy_Nil(t *testing.T) {
	var c *common.CacheConfig
	require.NotPanics(t, func() {
		assert.Nil(t, c.Copy())
	})

	// A config without CombinedConfig/Local/Remote must copy without panicking.
	partial := &common.CacheConfig{Ttl: 5, LatestCacheTTL: 7}
	cp := partial.Copy()
	require.NotNil(t, cp)
	assert.Nil(t, cp.CombinedConfig)
	assert.EqualValues(t, 5, cp.Ttl)
	assert.EqualValues(t, 7, cp.LatestCacheTTL)
}

func TestCacheConfigCopy_DeepCopy(t *testing.T) {
	orig := &common.CacheConfig{
		CombinedConfig: &zcache.CombinedConfig{
			GlobalPrefix: "prefix",
			Local:        &zcache.LocalConfig{Prefix: "local"},
			Remote:       &zcache.RemoteConfig{Addr: "localhost:6379", TLSEnabled: true},
		},
	}
	cp := orig.Copy()
	cp.GlobalPrefix = "changed"
	cp.Local.Prefix = "changed"
	cp.Remote.Addr = "changed"

	assert.Equal(t, "prefix", orig.GlobalPrefix)
	assert.Equal(t, "local", orig.Local.Prefix)
	assert.Equal(t, "localhost:6379", orig.Remote.Addr)
	assert.True(t, cp.Remote.TLSEnabled)
}
