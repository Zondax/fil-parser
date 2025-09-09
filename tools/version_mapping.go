package tools

import (
	"container/list"
	"fmt"
	"math"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/lotus/build/buildconstants"
)

// All versions and actor versions are listed here:
// https://github.com/filecoin-project/go-state-types/blob/master/network/version.go

// The minimimum calibration version is V16 because of a calibration reset.

const (
	CalibrationNetworkNodeType = "calibrationnet"
	CalibrationNetwork         = "calibration"
	MainnetNetwork             = "mainnet"
)

type version struct {
	calibration abi.ChainEpoch
	mainnet     abi.ChainEpoch

	nodeVersion    uint
	currentNetwork string
}

var (
	LatestMainnetVersion     version = V25
	LatestCalibrationVersion version = V27

	supportedVersions     = []version{V0, V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22, V23, V24, V25, V27}
	supportedVersionsList *list.List

	// V0 genesis, spec-actors: v1, calibration: 0, mainnet: 0
	V0 version = version{calibration: 0, mainnet: 0, nodeVersion: 0}
	// V1 Breeze, spec-actors: v1, calibration: 0, mainnet: 41280
	V1 version = version{calibration: 0, mainnet: buildconstants.UpgradeBreezeHeight, nodeVersion: 1}
	// V2 Smoke spec-actors: v1, calibration: 0, mainnet: 51000
	V2 version = version{calibration: 0, mainnet: buildconstants.UpgradeSmokeHeight, nodeVersion: 2}
	// V3 Ignition, spec-actors: v1, calibration: 0, mainnet: 94000
	V3 version = version{calibration: 0, mainnet: buildconstants.UpgradeIgnitionHeight, nodeVersion: 3}

	// V4 Refuel, spec-actors: v2, calibration: 0, mainnet: 130800
	V4 version = version{calibration: 0, mainnet: buildconstants.UpgradeRefuelHeight, nodeVersion: 4}
	// V5 Tape, spec-actors: v2, calibration: 0, mainnet: 140760
	V5 version = version{calibration: 0, mainnet: buildconstants.UpgradeTapeHeight, nodeVersion: 5}
	// V6 Kumquat, spec-actors: v2, calibration: 0, mainnet: 170000
	V6 version = version{calibration: 0, mainnet: buildconstants.UpgradeKumquatHeight, nodeVersion: 6}
	// V7 Calico, spec-actors: v2, calibration: 0, mainnet: 265200
	V7 version = version{calibration: 0, mainnet: buildconstants.UpgradeCalicoHeight, nodeVersion: 7}
	// V8 Persian, spec-actors: v2, calibration: 0, mainnet: 272400
	V8 version = version{calibration: 0, mainnet: buildconstants.UpgradePersianHeight, nodeVersion: 8}
	// V9 Orange, spec-actors: v2, calibration: 0, mainnet: 336458
	V9 version = version{calibration: 0, mainnet: buildconstants.UpgradeOrangeHeight, nodeVersion: 9}

	// V10 Trust, spec-actors: v3, calibration: 0, mainnet: 550321
	V10 version = version{calibration: 0, mainnet: buildconstants.UpgradeTrustHeight, nodeVersion: 10}
	// V11 Norwegian, spec-actors: v3, calibration: 0, mainnet: 665280
	V11 version = version{calibration: 0, mainnet: buildconstants.UpgradeNorwegianHeight, nodeVersion: 11}

	// V12 Turbo, spec-actors: v4, calibration: 0 (actual: 193789), mainnet: 712320
	V12 version = version{calibration: 0, mainnet: buildconstants.UpgradeTurboHeight, nodeVersion: 12}

	// V13 Hyperdrive, spec-actors: v5, calibration: 0, mainnet: 892800
	//
	// calibration reset
	V13 version = version{calibration: 0, mainnet: buildconstants.UpgradeHyperdriveHeight, nodeVersion: 13}

	// V14 Chocolate, spec-actors: v6, calibration: 0 (actual: 312746), mainnet: 1231620
	V14 version = version{calibration: 0, mainnet: buildconstants.UpgradeChocolateHeight, nodeVersion: 14}

	// V15 OhSnap,spec-actors: v7, calibration: 0 (actual: 682006), mainnet: 1594680
	V15 version = version{calibration: 0, mainnet: buildconstants.UpgradeOhSnapHeight, nodeVersion: 15}

	// V16 Skyr, builtin-actors(go-state-types): v8, calibration: 0 (actual: 1044660), mainnet: 1960320.
	//
	// parsing all calibration heights from 0->16799 with V16.
	V16 version = version{calibration: 0, mainnet: buildconstants.UpgradeSkyrHeight, nodeVersion: 16}

	// V17 Shark, builtin-actors(go-state-types): v9, calibration: 16800, mainnet: 2383680
	//
	// calibration reset
	V17 version = version{calibration: 16800, mainnet: buildconstants.UpgradeSharkHeight, nodeVersion: 17}

	// V18 Hygge, builtin-actors(go-state-types): v10, calibration: 322354, mainnet: 2683348
	V18 version = version{calibration: 322354, mainnet: buildconstants.UpgradeHyggeHeight, nodeVersion: 18}

	// V19 Lightning, builtin-actors(go-state-types): v11, calibration: 489094, mainnet: 2809800
	V19 version = version{calibration: 489094, mainnet: buildconstants.UpgradeLightningHeight, nodeVersion: 19}
	// V20 Thunder, builtin-actors(go-state-types): v11, calibration: 492214, mainnet: 2870280
	V20 version = version{calibration: 492214, mainnet: buildconstants.UpgradeThunderHeight, nodeVersion: 20}

	// V21 Watermelon, builtin-actors(go-state-types): v12, calibration: 1108174, mainnet: 3469380
	V21 version = version{calibration: 1108174, mainnet: buildconstants.UpgradeWatermelonHeight, nodeVersion: 21}

	// V22 Dragon, builtin-actors(go-state-types): v13, calibration: 1427974, mainnet: 3855360
	V22 version = version{calibration: 1427974, mainnet: buildconstants.UpgradeDragonHeight, nodeVersion: 22}

	// V23 Waffle, builtin-actors(go-state-types): v14, calibration: 1779094, mainnet: 4154640
	V23 version = version{calibration: 1779094, mainnet: buildconstants.UpgradeWaffleHeight, nodeVersion: 23}

	// V24 Tuktuk, builtin-actors(go-state-types): v15, calibration: 2081674, mainnet: 4461240
	V24 version = version{calibration: 2081674, mainnet: buildconstants.UpgradeTuktukHeight, nodeVersion: 24}

	// V25 Teep, builtin-actors(go-state-types): v16, calibration: 2523454, mainnet: 4878840
	V25 version = version{calibration: 2523454, mainnet: buildconstants.UpgradeTeepHeight, nodeVersion: 25}

	// V26 was skipped

	// V27 GoldenWeek, builtin-actors(go-state-types): v17, calibration: 3007294, mainnet: <unknown>
	V27 version = version{calibration: 3007294, mainnet: buildconstants.UpgradeGoldenWeekHeight, nodeVersion: 27}
)

func init() {
	supportedVersionsList = list.New()
	for _, v := range supportedVersions {
		supportedVersionsList.PushBack(v)
	}
}

func LatestVersion(network string) version {
	if network == CalibrationNetwork {
		return LatestCalibrationVersion
	}
	return LatestMainnetVersion
}

func ParseRawNetworkName(network string) string {
	if network == CalibrationNetworkNodeType || network == CalibrationNetwork {
		return CalibrationNetwork
	}
	return MainnetNetwork
}

// IsSupported returns true if the height is within the version range for a given network
func (v version) IsSupported(network string, height int64) bool {
	iter := NewVersionIterator(v, network)
	return isSupported(network, abi.ChainEpoch(height), iter)
}

func isSupported(network string, height abi.ChainEpoch, iter *VersionIterator) bool {
	v, ok := iter.Peek()
	if !ok {
		return false
	}
	if network == CalibrationNetwork {
		return checkCalibrationEdgeCases(network, height, iter)
	}
	// edge case: the calibration upgrade is done before mainnet.
	// so if the version is greater than the latest mainnet version, it is not supported
	if v.nodeVersion > LatestMainnetVersion.nodeVersion {
		return false
	}

	if height >= LatestMainnetVersion.mainnet {
		return v.nodeVersion == LatestMainnetVersion.nodeVersion
	}

	// edge case: check if two new versions have the same mainnet height
	next, ok := iter.PeekNext()
	if ok && v.mainnet == next.mainnet && !IsLatestVersion(v) {
		iter.Next()
		return isSupported(network, height, iter)
	}

	if height >= v.mainnet && height < next.mainnet {
		return true
	}
	return false

}

func checkCalibrationEdgeCases(network string, height abi.ChainEpoch, iter *VersionIterator) bool {
	v, ok := iter.Peek()
	if !ok {
		return false
	}
	if height == 0 && v.calibration == 0 {
		return true
	}
	if height >= LatestCalibrationVersion.calibration {
		return v.nodeVersion == LatestCalibrationVersion.nodeVersion
	}
	if v.nodeVersion < V16.nodeVersion {
		// on calibration, all versions before V16 are not used because of a calibration reset
		return false
	}

	// if height < V19.calibration {
	// parse all calibration heights before V19 with the V18 network version parsers because there was a calibration reset somewhere between V16 and V17.
	// and through testing we have determined there was a reset in V18
	// return v.nodeVersion == V18.nodeVersion
	// }

	next, ok := iter.PeekNext()
	// edge case: check if two new versions have the same calibration height
	if ok && v.calibration == next.calibration && !IsLatestVersion(v) {
		iter.Next()
		return isSupported(network, height, iter)
	}
	// check if the height is greater than the current version  but less than the next version
	if height >= v.calibration && height < next.calibration {
		return true
	}
	return false
}

// IsLatestVersion returns true if the version is the latest version
func IsLatestVersion(version version) bool {
	return version.nodeVersion == LatestVersion(version.currentNetwork).nodeVersion
}

// AnyIsSupported returns true if any of the versions are supported for a given network and height
func AnyIsSupported(network string, height int64, versions ...version) bool {
	for _, v := range versions {
		if v.IsSupported(network, height) {
			return true
		}
	}
	return false
}

// String returns the version as a string
func (v version) String() string {
	return fmt.Sprintf("V%d", v.nodeVersion)
}

// Height returns the height of a given version
// if the version is on the calibration network, it returns the calibration height
// otherwise, it returns the mainnet height
func (v version) Height() int64 {
	if v.currentNetwork == CalibrationNetwork {
		return int64(v.calibration)
	}
	return int64(v.mainnet)
}

func (v version) NodeVersion() uint {
	return v.nodeVersion
}

func (v version) FilNetworkVersion() network.Version {
	return network.Version(v.nodeVersion)
}

// GetSupportedVersions returns all supported versions for a given network
func GetSupportedVersions(network string) []version {
	var result []version
	iter := NewVersionIterator(V0, network)
	for v, ok := iter.Begin(); ok; v, ok = iter.Next() {
		v.currentNetwork = network
		result = append(result, v)
	}
	return result
}

// VersionsBefore returns all versions before the given version (inclusive of the start version)
func VersionsBefore(uptoIncluding version) []version {
	var result []version
	iter := NewVersionIterator(V0, uptoIncluding.currentNetwork)
	for v, ok := iter.Begin(); ok; v, ok = iter.Next() {
		if v.nodeVersion > uptoIncluding.nodeVersion {
			break
		}
		v.currentNetwork = uptoIncluding.currentNetwork
		result = append(result, v)
	}
	return result
}

// VersionsAfter returns all versions after the given version (inclusive of the start version)
func VersionsAfter(start version) []version {
	var result []version
	iter := NewVersionIterator(start, start.currentNetwork)
	for v, ok := iter.Begin(); ok; v, ok = iter.Next() {
		v.currentNetwork = start.currentNetwork
		result = append(result, v)
	}
	return result
}

// VersionRange returns the height range of a given network version
// if the version is the latest version, it returns the height range of the latest version
// if configured version heights are 0, it returns 0 to the latest known version heights (V8 and above)
func VersionRange(version version) (int64, int64) {
	var min, max int64
	if version.nodeVersion == LatestVersion(version.currentNetwork).nodeVersion {
		min, max = version.Height(), math.MaxInt64
	} else {
		iter := NewVersionIterator(version, version.currentNetwork)
		next, ok := iter.PeekNext()
		if !ok {
			return min, math.MaxInt64
		}
		min, max = version.Height(), next.Height()
	}
	if min == 0 && max == 0 {
		return 0, V8.Height()
	}
	return min, max
}

// VersionFromString returns the version struct from a given version string
// if the version is not found, it returns the first version (V1)
func VersionFromString(version string) version {
	for _, v := range supportedVersions {
		if v.String() == version {
			return v
		}
	}
	return V0
}

// VersionFromHeight returns the version for a given network and height.
// https://github.com/filecoin-project/go-state-types/blob/master/network/version.go
// The minimum calibration version is V16 ( height 0 -> 16799 will always return V16 for calibration).
func VersionFromHeight(network string, height int64) version {
	switch {
	case V0.IsSupported(network, height):
		return V0
	case V1.IsSupported(network, height):
		return V1
	case V2.IsSupported(network, height):
		return V2
	case V3.IsSupported(network, height):
		return V3
	case V4.IsSupported(network, height):
		return V4
	case V5.IsSupported(network, height):
		return V5
	case V6.IsSupported(network, height):
		return V6
	case V7.IsSupported(network, height):
		return V7
	case V8.IsSupported(network, height):
		return V8
	case V9.IsSupported(network, height):
		return V9
	case V10.IsSupported(network, height):
		return V10
	case V11.IsSupported(network, height):
		return V11
	case V12.IsSupported(network, height):
		return V12
	case V13.IsSupported(network, height):
		return V13
	case V14.IsSupported(network, height):
		return V14
	case V15.IsSupported(network, height):
		return V15
	case V16.IsSupported(network, height):
		return V16
	case V17.IsSupported(network, height):
		return V17
	case V18.IsSupported(network, height):
		return V18
	case V19.IsSupported(network, height):
		return V19
	case V20.IsSupported(network, height):
		return V20
	case V21.IsSupported(network, height):
		return V21
	case V22.IsSupported(network, height):
		return V22
	case V23.IsSupported(network, height):
		return V23
	case V24.IsSupported(network, height):
		return V24
	case V25.IsSupported(network, height):
		return V25
	}
	return LatestVersion(network)
}
