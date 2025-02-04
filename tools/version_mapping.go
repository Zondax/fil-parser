package tools

import (
	"fmt"
	"math"
)

type version struct {
	calibration int64
	mainnet     int64

	nodeVersion    int64
	currentNetwork string
}

var (
	supportedVersions         = []version{V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22, V23, V24}
	V1                version = version{calibration: 0, mainnet: 0, nodeVersion: 1}
	V2                version = version{calibration: 0, mainnet: 0, nodeVersion: 2}
	V3                version = version{calibration: 0, mainnet: 0, nodeVersion: 3}
	V4                version = version{calibration: 0, mainnet: 0, nodeVersion: 4}
	V5                version = version{calibration: 0, mainnet: 0, nodeVersion: 5}
	V6                version = version{calibration: 0, mainnet: 0, nodeVersion: 6}
	V7                version = version{calibration: 0, mainnet: 0, nodeVersion: 7}
	V8                version = version{calibration: 0, mainnet: 170000, nodeVersion: 8}
	V9                version = version{calibration: 0, mainnet: 265200, nodeVersion: 9}
	V10               version = version{calibration: 0, mainnet: 550321, nodeVersion: 10}
	V11               version = version{calibration: 0, mainnet: 665280, nodeVersion: 11}
	V12               version = version{calibration: 193789, mainnet: 712320, nodeVersion: 12}
	V13               version = version{calibration: 0, mainnet: 892800, nodeVersion: 13} // calibration reset
	V14               version = version{calibration: 312746, mainnet: 1231620, nodeVersion: 14}
	V15               version = version{calibration: 682006, mainnet: 1594680, nodeVersion: 15}
	V16               version = version{calibration: 1044660, mainnet: 1960320, nodeVersion: 16}
	V17               version = version{calibration: 16800, mainnet: 2383680, nodeVersion: 17} // calibration reset
	V18               version = version{calibration: 322354, mainnet: 2683348, nodeVersion: 18}
	V19               version = version{calibration: 489094, mainnet: 2809800, nodeVersion: 19}
	V20               version = version{calibration: 492214, mainnet: 2809800, nodeVersion: 20}
	V21               version = version{calibration: 1108174, mainnet: 3469380, nodeVersion: 21}
	V22               version = version{calibration: 1427974, mainnet: 3817920, nodeVersion: 22}
	V23               version = version{calibration: 1779094, mainnet: 4154640, nodeVersion: 23}
	V24               version = version{calibration: 2081674, mainnet: 4461240, nodeVersion: 24}
	LatestVersion     version = V24
)

func (v version) IsSupported(network string, height int64) bool {
	if network == "calibration" {
		if height == 0 && v.calibration == 0 {
			return true
		}
		if height >= LatestVersion.calibration {
			return v == LatestVersion
		}

		// edge case: check if two new versions have the same calibration height
		if v.calibration == v.next().calibration && !IsLatestVersion(v) {
			return v.next().IsSupported(network, height)
		}
		// check if the height is greater than the current version  but less than the next version
		if height >= v.calibration && height < v.next().calibration {
			return true
		}
	} else {
		if height == 0 && v.mainnet == 0 {
			return true
		}
		if height >= LatestVersion.mainnet {
			return v == LatestVersion
		}

		// edge case: check if two new versions have the same mainnet height
		if v.mainnet == v.next().mainnet && !IsLatestVersion(v) {
			return v.next().IsSupported(network, height)
		}
		// check if the height is greater than the current version  but less than the next version
		if height >= v.mainnet && height < v.next().mainnet {
			return true
		}
	}

	return false
}

func IsLatestVersion(version version) bool {
	return version.nodeVersion == LatestVersion.nodeVersion
}

func AnyIsSupported(network string, height int64, versions ...version) bool {
	for _, v := range versions {
		if v.IsSupported(network, height) {
			return true
		}
	}
	return false
}

func (v version) next() version {
	for i, version := range supportedVersions {
		if version.nodeVersion == v.nodeVersion {
			if i == len(supportedVersions)-1 {
				return v
			}
			return supportedVersions[i+1]
		}
	}
	return v
}

func (v version) String() string {
	return fmt.Sprintf("V%d", v.nodeVersion)
}

func (v version) Height() int64 {
	if v.currentNetwork == "calibration" {
		return v.calibration
	}
	return v.mainnet
}

func GetSupportedVersions(network string) []version {
	var result []version
	for _, v := range supportedVersions {
		v.currentNetwork = network
		result = append(result, v)
	}
	return result
}

// VersionsBefore returns all versions before the given version (inclusive of the start version)
func VersionsBefore(version version) []version {
	offset := 1 // we start at version 1
	for i, v := range supportedVersions {
		if i+offset > len(supportedVersions) {
			return supportedVersions[:i]
		}
		if v.nodeVersion == version.nodeVersion {
			return supportedVersions[:i+offset]
		}
	}
	return nil
}

// VersionsAfter returns all versions after the given version (inclusive of the start version)
func VersionsAfter(start version) []version {
	var result []version
	offset := 1 // we start at version 1
	for i, v := range supportedVersions {
		if i+offset > len(supportedVersions) {
			result = append(result, v)
			break
		}
		if v.nodeVersion == start.nodeVersion {
			result = append(result, supportedVersions[i:]...)
			break
		}
	}
	return result
}

func VersionRange(version version) (int64, int64) {
	var min, max int64
	if version.next().nodeVersion == version.nodeVersion {
		min, max = version.Height(), math.MaxInt64
	} else {
		min, max = version.Height(), version.next().Height()
	}
	if min == 0 && max == 0 {
		return 0, V8.Height()
	}
	return min, max
}

func VersionFromString(version string) version {
	for _, v := range supportedVersions {
		if v.String() == version {
			return v
		}
	}
	return V1
}
