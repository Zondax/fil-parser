package tools

import "testing"

// TestVersionFromHeightResolvesEachVersion guards the most commonly-missed step of an NV bump:
// when LatestCalibrationVersion / LatestMainnetVersion is promoted to a new version, the
// PREVIOUS latest stops being the fall-through in VersionFromHeight and must gain an explicit
// `case`. Omitting it still compiles and still passes the actor-coverage tests, but silently
// mis-tags every height in the previous version's range as the new version.
//
// Add a row here for each version as it is promoted.
//
// Note: while LatestMainnetVersion is still V28 the two mainnet rows below pass with or without
// the fix, because the fall-through already returns V28 on mainnet. The calibration V28 row is
// the one carrying the guarantee today; the mainnet rows start biting when mainnet is promoted.
func TestVersionFromHeightResolvesEachVersion(t *testing.T) {
	tests := []struct {
		name    string
		network string
		height  int64
		want    uint
	}{
		{"calibration V27 range", CalibrationNetwork, 3007294 + 10, 27},
		{"calibration V28 range", CalibrationNetwork, 3694534 + 10, 28},
		{"mainnet V27 range", MainnetNetwork, 5348280 + 10, 27},
		{"mainnet V28 range", MainnetNetwork, 6052800 + 10, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VersionFromHeight(tt.network, tt.height).nodeVersion; got != tt.want {
				t.Errorf("VersionFromHeight(%s, %d) = V%d, want V%d", tt.network, tt.height, got, tt.want)
			}
		})
	}
}

// TestVersionsAfterIncludesCalibrationOnlyVersions guards a trap that silently breaks every
// network upgrade. VersionsAfter used to walk a VersionIterator, which is bounded by
// LatestVersion(currentNetwork); the package-level version values carry an empty currentNetwork,
// so that resolved to LatestMainnetVersion and quietly dropped any version already live on
// calibration but not yet promoted on mainnet. Callers such as Power.OnConsensusFault feed the
// result to AnyIsSupported and fall through to "unsupported height" for the missing version.
func TestVersionsAfterIncludesCalibrationOnlyVersions(t *testing.T) {
	got := VersionsAfter(V16)
	found := false
	for _, v := range got {
		if v.nodeVersion == LatestCalibrationVersion.nodeVersion {
			found = true
		}
	}
	if !found {
		t.Fatalf("VersionsAfter(V16) omitted the latest calibration version V%d; got %v",
			LatestCalibrationVersion.nodeVersion, got)
	}
}
