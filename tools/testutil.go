package tools

func DeterministicTestHeight(version version) int64 {
	if version.currentNetwork == CalibrationNetwork {
		// min calibration parsing height is V16
		if version.nodeVersion < V16.nodeVersion {
			return V16.Height()
		}
	}
	iter := NewVersionIterator(version, version.currentNetwork)
	next, ok := iter.PeekNext()
	if !ok {
		return version.Height()
	}
	return (version.Height() + next.Height()) / 2
}
