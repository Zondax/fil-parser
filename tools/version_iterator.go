package tools

import "container/list"

// VersionIterator provides safe concurrent traversal of versions
type VersionIterator struct {
	current *list.Element
	network string
	from    version
}

// NewVersionIterator creates a new iterator starting at the beginning of the list
func NewVersionIterator(from version, network string) *VersionIterator {
	v := &VersionIterator{
		network: network,
		from:    from,
	}
	v.Reset()
	return v
}

// Next returns the current version and advances the iterator
// Returns false when there are no more versions or the latest version for the network is reached
func (vi *VersionIterator) Next() (version, bool) {
	if vi.current == nil {
		return version{}, false
	}
	v := vi.current.Value.(version)
	v.currentNetwork = vi.network
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
	v := vi.current.Value.(version)
	v.currentNetwork = vi.network
	return v, true
}

// PeekNext returns the next version in the list without advancing the iterator
func (vi *VersionIterator) PeekNext() (version, bool) {
	if vi.current == nil {
		return version{}, false
	}
	if vi.current.Next() == nil {
		return version{}, false
	}
	v := vi.current.Next().Value.(version)
	v.currentNetwork = vi.network
	return v, true
}

// Begin moves the iterator to the first version to be returned for a for-loop init
func (vi *VersionIterator) Begin() (version, bool) {
	vi.Reset()
	return vi.Next()
}

// Reset moves the iterator back to the start of the list
func (vi *VersionIterator) Reset() {
	current := supportedVersionsList.Front()
	for current != nil && current.Value.(version).nodeVersion < vi.from.nodeVersion {
		// skip versions that are before the start version
		// stop if we reach the latest version
		if current.Value.(version).nodeVersion == LatestVersion(vi.network).nodeVersion {
			break
		}
		current = current.Next()
	}
	vi.current = current
}
