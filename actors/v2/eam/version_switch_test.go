package eam

import (
	"reflect"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	typegen "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/golem/pkg/logger"
)

// newEamCreate and validateEamReturn dispatch on the concrete go-state-types type of each
// version; feed them every type the return maps can produce.
func TestCreateReturnHelpersHandleEveryMappedVersion(t *testing.T) {
	e := &Eam{logger: logger.NewDevelopmentLogger()}
	for name, m := range map[string]map[string]func() typegen.CBORUnmarshaler{
		"Create":         createReturn,
		"Create2":        create2Return,
		"CreateExternal": createExternalReturn,
	} {
		for version, newReturn := range m {
			ret := newReturn()
			_, _, _, err := e.newEamCreate(ret, cid.Undef)
			require.NoError(t, err, "%s return at %s (%T) is not handled by newEamCreate", name, version, ret)

			require.NoError(t, validateEamReturn(ret))
			robust := reflect.ValueOf(ret).Elem().FieldByName("RobustAddress")
			require.True(t, robust.IsValid(), "%T has no RobustAddress", ret)
			require.False(t, robust.IsNil(), "%s return at %s (%T): validateEamReturn left a nil RobustAddress", name, version, ret)
		}
	}
}
