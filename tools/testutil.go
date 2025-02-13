package tools

func DeterministicTestHeight(version version) int64 {
	height := version.Height()
	if version != version.next() {
		height = (version.Height() + version.next().Height()) / 2
	}

	return height
}
