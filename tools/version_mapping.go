package tools

import "fmt"

var (
	V1            version = version{calibration: 0, mainnet: 0}
	V2            version = version{calibration: 0, mainnet: 0}
	V3            version = version{calibration: 0, mainnet: 0}
	V4            version = version{calibration: 0, mainnet: 0}
	V5            version = version{calibration: 0, mainnet: 0}
	V6            version = version{calibration: 0, mainnet: 0}
	V7            version = version{calibration: 0, mainnet: 0}
	V8            version = version{calibration: 0, mainnet: 170000}
	V9            version = version{calibration: 0, mainnet: 265200}
	V10           version = version{calibration: 0, mainnet: 550321}
	V11           version = version{calibration: 0, mainnet: 665280}
	V12           version = version{calibration: 193789, mainnet: 712320}
	V13           version = version{calibration: 0, mainnet: 892800} // calibration reset
	V14           version = version{calibration: 312746, mainnet: 1231620}
	V15           version = version{calibration: 682006, mainnet: 1594680}
	V16           version = version{calibration: 1044660, mainnet: 1960320}
	V17           version = version{calibration: 16800, mainnet: 2383680} // calibration reset
	V18           version = version{calibration: 322354, mainnet: 2683348}
	V19           version = version{calibration: 489094, mainnet: 2809800}
	V20           version = version{calibration: 492214, mainnet: 2809800}
	V21           version = version{calibration: 1108174, mainnet: 3469380}
	V22           version = version{calibration: 1427974, mainnet: 3817920}
	V23           version = version{calibration: 1779094, mainnet: 4154640}
	V24           version = version{calibration: 2081674, mainnet: 4461240}
	LatestVersion version = V24
)

type version struct {
	calibration int64
	mainnet     int64
}

var supportedVersions = []version{V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22}

func (v version) IsSupported(network string, height int64) bool {
	if network == "calibration" {
		return height <= v.calibration
	}
	return height <= v.mainnet
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
	return fmt.Sprintf("V%d", v.startHeight)
}
func GetSupportedVersions() []version {
	// Create a new slice with same length and copy all elements
	result := make([]version, len(supportedVersions))
	copy(result, supportedVersions)
	return result
}
