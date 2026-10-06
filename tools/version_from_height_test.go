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
// With LatestMainnetVersion at V29, the mainnet V28 rows rely on the explicit V28 case (the
// fall-through now returns V29 on mainnet), like the calibration V28 row does.
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
		{"calibration V29 range", CalibrationNetwork, 4109133 + 10, 29},
		// NV29 (Solstice) mainnet: lotus v1.37.0 sets UpgradeSolsticeHeight = 6470279 (2026-10-19T12:59:30Z).
		{"mainnet last V28 height", MainnetNetwork, 6470279 - 1, 28},
		{"mainnet V29 at upgrade height", MainnetNetwork, 6470279, 29},
		{"mainnet V29 range", MainnetNetwork, 6470279 + 10, 29},
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

// V29's mainnet height comes from lotus buildconstants: lotus v1.37.0 schedules it at 6470279.
// A lotus pin still parking Solstice at UpgradeHeightUnscheduled would silently keep every
// post-fork mainnet height on V28.
func TestV29MainnetScheduled(t *testing.T) {
	if V29.mainnet != 6470279 {
		t.Fatalf("V29 mainnet height = %d, want 6470279 (lotus v1.37.0 UpgradeSolsticeHeight)", V29.mainnet)
	}
	if LatestMainnetVersion.nodeVersion != 29 {
		t.Fatalf("LatestMainnetVersion = V%d, want V29", LatestMainnetVersion.nodeVersion)
	}
}
