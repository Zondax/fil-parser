package tools

import "fmt"

var (
	LatestVersion version = version{calibration: 22000, mainnet: 22000}
	V1            version = version{calibration: 0, mainnet: 0}
	V2            version = version{calibration: 100, mainnet: 100}
	V3            version = version{calibration: 1000, mainnet: 1000}
	V4            version = version{calibration: 2000, mainnet: 2000}
	V5            version = version{calibration: 3000, mainnet: 3000}
	V6            version = version{calibration: 4000, mainnet: 4000}
	V7            version = version{calibration: 5000, mainnet: 5000}
	V8            version = version{calibration: 6000, mainnet: 6000}
	V9            version = version{calibration: 7000, mainnet: 7000}
	V10           version = version{calibration: 8000, mainnet: 8000}
	V11           version = version{calibration: 9000, mainnet: 9000}
	V12           version = version{calibration: 10000, mainnet: 10000}
	V13           version = version{calibration: 11000, mainnet: 11000}
	V14           version = version{calibration: 12000, mainnet: 12000}
	V15           version = version{calibration: 13000, mainnet: 13000}
	V16           version = version{calibration: 14000, mainnet: 14000}
	V17           version = version{calibration: 15000, mainnet: 15000}
	V18           version = version{calibration: 16000, mainnet: 16000}
	V19           version = version{calibration: 17000, mainnet: 17000}
	V20           version = version{calibration: 18000, mainnet: 18000}
	V21           version = version{calibration: 19000, mainnet: 19000}
	V22           version = version{calibration: 20000, mainnet: 20000}
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
