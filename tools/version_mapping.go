package tools

import "fmt"

type version struct {
	calibration int64
	mainnet     int64

	nodeVersion    int64
	currentNetwork string
}

var (
	supportedVersions         = []version{V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22}
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
		if height >= LatestVersion.calibration {
			return v == LatestVersion
		}
		// check if the height is greater than the current version  but less than the next version
		if height > v.calibration && height < v.next().calibration {
			return true
		}
	} else {
		if height >= LatestVersion.mainnet {
			return v == LatestVersion
		}
		// check if the height is greater than the current version  but less than the next version
		if height > v.mainnet && height < v.next().mainnet {
			return true
		}
	}

	return false
}

func (v version) next() version {
	for i, version := range supportedVersions {
		if version == v {
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
	result := make([]version, len(supportedVersions))
	for _, v := range supportedVersions {
		v.currentNetwork = network
		result = append(result, v)
	}
	return result
}
