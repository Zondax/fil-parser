package tools

import "testing"

// upgradeHeights returns every height > 0 at which VersionFromHeight changes on network, with the
// version it returns from there on.
func upgradeHeights(network string) map[int64]version {
	out := map[int64]version{}
	for _, v := range supportedVersions {
		h := int64(v.mainnet)
		if network == CalibrationNetwork {
			h = int64(v.calibration)
		}
		if h <= 1 {
			continue
		}
		if VersionFromHeight(network, h).nodeVersion != VersionFromHeight(network, h-1).nodeVersion {
			out[h] = VersionFromHeight(network, h)
		}
	}
	return out
}

// TestExecutedVersionFromHeightAtEveryUpgrade checks, for every upgrade on both networks, that the
// tipset at the upgrade height U is executed by the old version and the tipset at U+1 by the new
// one, while VersionFromHeight keeps labelling U with the new version.
func TestExecutedVersionFromHeightAtEveryUpgrade(t *testing.T) {
	for _, network := range []string{MainnetNetwork, CalibrationNetwork} {
		ups := upgradeHeights(network)
		if len(ups) < 10 {
			t.Fatalf("%s: found only %d upgrades", network, len(ups))
		}
		for u, newV := range ups {
			oldV := VersionFromHeight(network, u-1)
			if got := ExecutedVersionFromHeight(network, u); got.nodeVersion != oldV.nodeVersion {
				t.Errorf("%s: ExecutedVersionFromHeight(%d) = %s, want %s (old)", network, u, got, oldV)
			}
			if got := ExecutedVersionFromHeight(network, u+1); got.nodeVersion != newV.nodeVersion {
				t.Errorf("%s: ExecutedVersionFromHeight(%d) = %s, want %s (new)", network, u+1, got, newV)
			}
			if got := VersionFromHeight(network, u); got.nodeVersion != newV.nodeVersion {
				t.Errorf("%s: VersionFromHeight(%d) = %s, want %s (unchanged)", network, u, got, newV)
			}
			if !oldV.IsExecutedAt(network, u) || newV.IsExecutedAt(network, u) {
				t.Errorf("%s: IsExecutedAt(%d) must hold for %s only", network, u, oldV)
			}
			if !newV.IsExecutedAt(network, u+1) || oldV.IsExecutedAt(network, u+1) {
				t.Errorf("%s: IsExecutedAt(%d) must hold for %s only", network, u+1, newV)
			}
			nv := newV
			nv.currentNetwork = network
			if got := nv.FirstExecutedHeight(); got != u+1 {
				t.Errorf("%s: %s.FirstExecutedHeight() = %d, want %d", network, newV, got, u+1)
			}
			if !AnyIsExecutedAt(network, u, VersionsBefore(oldV)...) || AnyIsExecutedAt(network, u, VersionsAfter(newV)...) {
				t.Errorf("%s: AnyIsExecutedAt(%d) disagrees with ExecutedVersionFromHeight", network, u)
			}
		}
	}
}

// TestExecutedVersionFromHeightMatchesChain pins ExecutedVersionFromHeight to StateNetworkVersion
// as reported by the Zondax archival (stable) nodes at U and U+1 (2026-11-20). StateNetworkVersion
// of a tipset is the version of the VM that runs its messages (GetNetworkVersion(ts.Height())).
//
// Not listed, because the version table itself disagrees with the chain over a whole range (not an
// off-by-one; tracked separately): mainnet V4 (the table uses Refuel 130800, which kept NV3; NV4
// starts after 138720), calibration V21 (the table uses WatermelonFix2 1108174; NV21 starts after
// 1013134) and calibration V24 (the table uses 2081674; NV24 starts after 2078794).
func TestExecutedVersionFromHeightMatchesChain(t *testing.T) {
	chain := []struct {
		network string
		upgrade int64
		before  uint // StateNetworkVersion at upgrade
		after   uint // StateNetworkVersion at upgrade+1
	}{
		{MainnetNetwork, 41280, 0, 1},
		{MainnetNetwork, 51000, 1, 2},
		{MainnetNetwork, 94000, 2, 3},
		{MainnetNetwork, 140760, 4, 5},
		{MainnetNetwork, 170000, 5, 6},
		{MainnetNetwork, 265200, 6, 7},
		{MainnetNetwork, 272400, 7, 8},
		{MainnetNetwork, 336458, 8, 9},
		{MainnetNetwork, 550321, 9, 10},
		{MainnetNetwork, 665280, 10, 11},
		{MainnetNetwork, 712320, 11, 12},
		{MainnetNetwork, 892800, 12, 13},
		{MainnetNetwork, 1231620, 13, 14},
		{MainnetNetwork, 1594680, 14, 15},
		{MainnetNetwork, 1960320, 15, 16},
		// 2383681 is a null round; the next tipset (2383682) runs NV17.
		{MainnetNetwork, 2383680, 16, 17},
		{MainnetNetwork, 2683348, 17, 18},
		{MainnetNetwork, 2809800, 18, 19},
		{MainnetNetwork, 2870280, 19, 20},
		{MainnetNetwork, 3469380, 20, 21},
		{MainnetNetwork, 3855360, 21, 22},
		{MainnetNetwork, 4154640, 22, 23},
		{MainnetNetwork, 4461240, 23, 24},
		{MainnetNetwork, 4878840, 24, 25},
		{MainnetNetwork, 5138040, 25, 26},
		{MainnetNetwork, 5348280, 26, 27},
		{MainnetNetwork, 6052800, 27, 28},
		{MainnetNetwork, 6470279, 28, 29},
		{CalibrationNetwork, 16800, 16, 17},
		{CalibrationNetwork, 322354, 17, 18},
		{CalibrationNetwork, 489094, 18, 19},
		{CalibrationNetwork, 492214, 19, 20},
		{CalibrationNetwork, 1427974, 21, 22},
		{CalibrationNetwork, 1779094, 22, 23},
		{CalibrationNetwork, 2523454, 24, 25},
		{CalibrationNetwork, 2543614, 25, 26},
		{CalibrationNetwork, 3007294, 26, 27},
		{CalibrationNetwork, 3694534, 27, 28},
		{CalibrationNetwork, 4109133, 28, 29},
	}
	for _, c := range chain {
		if got := ExecutedVersionFromHeight(c.network, c.upgrade).nodeVersion; got != c.before {
			t.Errorf("%s: ExecutedVersionFromHeight(%d) = V%d, chain runs NV%d", c.network, c.upgrade, got, c.before)
		}
		// for a null round at upgrade+1, the next tipset is the first one with the new version
		if got := ExecutedVersionFromHeight(c.network, c.upgrade+2).nodeVersion; got != c.after {
			t.Errorf("%s: ExecutedVersionFromHeight(%d) = V%d, chain runs NV%d", c.network, c.upgrade+2, got, c.after)
		}
		if got := ExecutedVersionFromHeight(c.network, c.upgrade+1).nodeVersion; got != c.after {
			t.Errorf("%s: ExecutedVersionFromHeight(%d) = V%d, chain runs NV%d", c.network, c.upgrade+1, got, c.after)
		}
	}
}

func TestExecutedVersionFromHeightGenesis(t *testing.T) {
	for _, network := range []string{MainnetNetwork, CalibrationNetwork} {
		for _, h := range []int64{0, 1} {
			if got, want := ExecutedVersionFromHeight(network, h), VersionFromHeight(network, h); got.nodeVersion != want.nodeVersion {
				t.Errorf("%s: ExecutedVersionFromHeight(%d) = %s, want %s", network, h, got, want)
			}
		}
	}
	if got := ExecutedVersionFromHeight(CalibrationNetwork, 2).nodeVersion; got != V16.nodeVersion {
		t.Errorf("calibration: ExecutedVersionFromHeight(2) = V%d, want V16", got)
	}
}
