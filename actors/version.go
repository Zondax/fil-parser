package actors

var heightToVersion = map[string][2]int64{
	"v11": {0, 100000},
}

const defaultVersion = "v14"

func getVersionFromHeight(height int64) (string, error) {
	for version, heightRange := range heightToVersion {
		if height >= heightRange[0] && height <= heightRange[1] {
			return version, nil
		}
	}
	return defaultVersion, nil
}
