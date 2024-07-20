package actors

const (
	VersionV11 = "v11"
	VersionV14 = "v14"
)

var heightToVersion = map[string][2]int64{
	VersionV11: {0, 100000},
}

const defaultVersion = VersionV14

func getVersionFromHeight(height int64) string {
	for version, heightRange := range heightToVersion {
		if height >= heightRange[0] && height <= heightRange[1] {
			return version
		}
	}
	return defaultVersion
}
