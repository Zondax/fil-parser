package tools

func DeterministicTestHeight(version version) int64 {
	iter := NewVersionIterator(version, version.currentNetwork)
	next, ok := iter.PeekNext()
	if !ok {
		return version.Height()
	}
	return (version.Height() + next.Height()) / 2
}
