package tools

import (
	"container/list"
	"fmt"
	"math"
)

const (
	CalibrationNetwork = "calibration"
	MainnetNetwork     = "mainnet"
)

type version struct {
	calibration int64
	mainnet     int64

	nodeVersion    int64
	currentNetwork string
}

var (
	LatestMainnetVersion     version = V24
	LatestCalibrationVersion version = V25

	supportedVersions     = []version{V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22, V23, V24, V25}
	supportedVersionsList *list.List

	V1  version = version{calibration: 0, mainnet: 0, nodeVersion: 1}
	V2  version = version{calibration: 0, mainnet: 0, nodeVersion: 2}
	V3  version = version{calibration: 0, mainnet: 0, nodeVersion: 3}
	V4  version = version{calibration: 0, mainnet: 0, nodeVersion: 4}
	V5  version = version{calibration: 0, mainnet: 0, nodeVersion: 5}
	V6  version = version{calibration: 0, mainnet: 0, nodeVersion: 6}
	V7  version = version{calibration: 0, mainnet: 0, nodeVersion: 7}
	V8  version = version{calibration: 0, mainnet: 170000, nodeVersion: 8}
	V9  version = version{calibration: 0, mainnet: 265200, nodeVersion: 9}
	V10 version = version{calibration: 0, mainnet: 550321, nodeVersion: 10}
	V11 version = version{calibration: 0, mainnet: 665280, nodeVersion: 11}
	V12 version = version{calibration: 193789, mainnet: 712320, nodeVersion: 12}
	V13 version = version{calibration: 0, mainnet: 892800, nodeVersion: 13} // calibration reset
	V14 version = version{calibration: 312746, mainnet: 1231620, nodeVersion: 14}
	V15 version = version{calibration: 682006, mainnet: 1594680, nodeVersion: 15}
	V16 version = version{calibration: 1044660, mainnet: 1960320, nodeVersion: 16}
	V17 version = version{calibration: 16800, mainnet: 2383680, nodeVersion: 17} // calibration reset
	V18 version = version{calibration: 322354, mainnet: 2683348, nodeVersion: 18}
	V19 version = version{calibration: 489094, mainnet: 2809800, nodeVersion: 19}
	V20 version = version{calibration: 492214, mainnet: 2809800, nodeVersion: 20}
	V21 version = version{calibration: 1108174, mainnet: 3469380, nodeVersion: 21}
	V22 version = version{calibration: 1427974, mainnet: 3817920, nodeVersion: 22}
	V23 version = version{calibration: 1779094, mainnet: 4154640, nodeVersion: 23}
	V24 version = version{calibration: 2081674, mainnet: 4461240, nodeVersion: 24}
	V25 version = version{calibration: 2523454, mainnet: -1, nodeVersion: 25}
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
	if network == "calibrationnet" || network == CalibrationNetwork {
		return CalibrationNetwork
	}
	return MainnetNetwork
}

// IsSupported returns true if the height is within the version range for a given network
func (v version) IsSupported(network string, height int64) bool {
	iter := NewVersionIterator(v, network)
	return isSupported(network, height, iter)
}

func isSupported(network string, height int64, iter *VersionIterator) bool {
	v, ok := iter.Peek()
	if !ok {
		return false
	}
	if network == CalibrationNetwork {
		return checkCalibrationEdgeCases(network, height, iter)
	} else {
		if height == 0 && v.mainnet == 0 {
			return true
		}
		if height >= LatestMainnetVersion.mainnet {
			return v.nodeVersion == LatestMainnetVersion.nodeVersion
		}
		if height <= V7.mainnet {
			return v.nodeVersion == V7.nodeVersion
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
}

func checkCalibrationEdgeCases(network string, height int64, iter *VersionIterator) bool {
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
	if v.nodeVersion < V19.nodeVersion {
		if v.nodeVersion != V18.nodeVersion {
			// on calibration, all versions before V18 are not used because of a calibration reset
			return false
		}
	}

	if height < V19.calibration {
		// parse all calibration heights before V19 with the V18 network version parsers because there was a calibration reset somewhere between V16 and V17.
		// and through testing we have determined there was a reset in V18
		return v.nodeVersion == V18.nodeVersion
	}

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
		return v.calibration
	}
	return v.mainnet
}

// GetSupportedVersions returns all supported versions for a given network
func GetSupportedVersions(network string) []version {
	var result []version
	iter := NewVersionIterator(V1, network)
	for v, ok := iter.Begin(); ok; v, ok = iter.Next() {
		v.currentNetwork = network
		result = append(result, v)
	}
	return result
}

// VersionsBefore returns all versions before the given version (inclusive of the start version)
func VersionsBefore(uptoIncluding version) []version {
	var result []version
	iter := NewVersionIterator(V1, uptoIncluding.currentNetwork)
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
	return V1
}
