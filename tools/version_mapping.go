package tools

import "fmt"

type version int64

var supportedVersions = []version{V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22}

func (v version) IsSupported(height int64) bool {
	return height <= int64(v.next())
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
	return fmt.Sprintf("V%d", v)
}
func GetSupportedVersions() []version {
	// Create a new slice with same length and copy all elements
	result := make([]version, len(supportedVersions))
	copy(result, supportedVersions)
	return result
}

const (
	LatestVersion version = V22
	V1            version = 0
	V2            version = 100
	V3            version = 1000
	V4            version = 2000
	V5            version = 3000
	V6            version = 4000
	V7            version = 5000
	V8            version = 6000
	V9            version = 7000
	V10           version = 8000
	V11           version = 9000
	V12           version = 10000
	V13           version = 11000
	V14           version = 12000
	V15           version = 13000
	V16           version = 14000
	V17           version = 15000
	V18           version = 16000
	V19           version = 17000
	V20           version = 18000
	V21           version = 19000
	V22           version = 20000
)
