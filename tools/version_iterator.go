package tools

import "container/list"

// VersionIterator provides safe concurrent traversal of versions
type VersionIterator struct {
	current *list.Element
	network string
}

// NewVersionIterator creates a new iterator starting at the beginning of the list
func NewVersionIterator(from version, network string) *VersionIterator {
	current := supportedVersionsList.Front()
	for current != nil && current.Value.(version).nodeVersion < from.nodeVersion {
		// skip versions that are before the start version
		// stop if we reach the latest version
		current = current.Next()
		if current.Value.(version).nodeVersion == LatestVersion(current.Value.(version).currentNetwork).nodeVersion {
			break
		}
	}
	return &VersionIterator{
		current: current,
		network: network,
	}
}

// Next moves to and returns the next version
// Returns false when there are no more versions or the latest version for the network is reached
func (vi *VersionIterator) Next() (version, bool) {
	if vi.current == nil {
		return version{}, false
	}
	v := vi.current.Value.(version)
	if v.nodeVersion > LatestVersion(vi.network).nodeVersion {
		return version{}, false
	}
	vi.current = vi.current.Next()
	return v, true
}

// Peek returns the current version in the list without advancing the iterator
func (vi *VersionIterator) Peek() (version, bool) {
	if vi.current == nil {
		return version{}, false
	}
	return vi.current.Value.(version), true
}

// PeekNext returns the next version in the list without advancing the iterator
func (vi *VersionIterator) PeekNext() (version, bool) {
	if vi.current == nil {
		return version{}, false
	}
	return vi.current.Next().Value.(version), true
}

// Begin moves the iterator to the first version to be returned for a for-loop init
func (vi *VersionIterator) Begin() (version, bool) {
	vi.Reset()
	return vi.Next()
}

// Reset moves the iterator back to the start of the list
func (vi *VersionIterator) Reset() {
	vi.current = supportedVersionsList.Front()
}
