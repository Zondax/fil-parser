package v2_test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"

	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/golem/pkg/logger"
)

// TestMethodNameHeightSemantics pins the two meanings of a height in the method-name API, using
// reward ClaimExported, added by builtin-actors v19 (NV29, calibration upgrade at 4109133).
//   - GetMethodName / ActorMethods: the version in force from height on (tools.VersionFromHeight).
//     External callers pass tools.LatestVersion(network).Height() to mean the latest version
//     (e.g. fil-indexing's MempoolWatcher), so the upgrade height resolves the new methods.
//   - GetMethodNameForTipset / ActorMethodsForTipset: the version that executed the tipset at
//     height (tools.ExecutedVersionFromHeight); the upgrade tipset still ran the old actors.
func TestMethodNameHeightSemantics(t *testing.T) {
	ctx := context.Background()
	const (
		net     = tools.CalibrationNetwork
		upgrade = int64(4109133)
	)
	l := logger.NewDevelopmentLogger()
	claim := builtin.MustGenerateFRCMethodNum("Claim")

	name, err := v2.GetMethodName(ctx, claim, manifest.RewardKey, upgrade, net, nil, l)
	require.NoError(t, err)
	require.Equal(t, "ClaimExported", name, "in-force semantics: the upgrade height resolves V29")

	name, err = v2.GetMethodNameForTipset(ctx, claim, manifest.RewardKey, upgrade, net, nil, l)
	require.NoError(t, err)
	require.Equal(t, parser.UnknownStr, name, "the tipset at the upgrade height was executed by V28")

	name, err = v2.GetMethodNameForTipset(ctx, claim, manifest.RewardKey, upgrade+1, net, nil, l)
	require.NoError(t, err)
	require.Equal(t, "ClaimExported", name)

	name, err = v2.GetMethodName(ctx, claim, manifest.RewardKey, upgrade-1, net, nil, l)
	require.NoError(t, err)
	require.Equal(t, parser.UnknownStr, name)

	inForce, err := v2.ActorMethods(ctx, manifest.RewardKey, upgrade, net, nil, l)
	require.NoError(t, err)
	executed, err := v2.ActorMethodsForTipset(ctx, manifest.RewardKey, upgrade, net, nil, l)
	require.NoError(t, err)
	require.Contains(t, inForce, claim)
	require.NotContains(t, executed, claim)

	// the latest-version sentinel used by external callers keeps resolving the latest version
	latest, err := v2.ActorMethods(ctx, manifest.RewardKey, tools.LatestVersion(net).Height(), net, nil, l)
	require.NoError(t, err)
	require.Contains(t, latest, claim)
}
