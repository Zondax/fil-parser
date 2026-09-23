package power

import (
	"reflect"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/stretchr/testify/require"

	"github.com/zondax/fil-parser/parser"
)

// getAddressInfo dispatches on the concrete go-state-types type of each version and silently
// returns nil for an unknown one; feed it every type the return map can produce.
func TestGetAddressInfoHandlesEveryMappedVersion(t *testing.T) {
	idAddr, err := address.NewIDAddress(1234)
	require.NoError(t, err)
	for version, newReturn := range createMinerReturn {
		ret := newReturn()
		f := reflect.ValueOf(ret).Elem().FieldByName("IDAddress")
		require.True(t, f.IsValid(), "%T has no IDAddress", ret)
		f.Set(reflect.ValueOf(idAddr))

		pr, ok := ret.(powerReturn)
		require.True(t, ok, "%T does not satisfy powerReturn", ret)
		info := getAddressInfo(pr, &parser.LotusMessage{})
		require.NotNil(t, info, "CreateMiner return at %s (%T) is not handled", version, ret)
		require.Equal(t, idAddr.String(), info.Short)
	}
}
