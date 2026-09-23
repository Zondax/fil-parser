package miner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// getControlAddress dispatches on the concrete go-state-types type of each version; feed it
// every type the return map can produce.
func TestGetControlAddressHandlesEveryMappedVersion(t *testing.T) {
	for version, newReturn := range getControlAddressesReturn {
		ret := newReturn()
		_, err := getControlAddress(ret)
		require.NoError(t, err, "GetControlAddresses return at %s (%T) is not handled", version, ret)
	}
}
