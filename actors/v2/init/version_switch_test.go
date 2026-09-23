package init

import (
	"reflect"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/stretchr/testify/require"
	typegen "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
)

// setExecParams and setReturnParams dispatch on the concrete go-state-types type of each
// version. A version wired into the params/return maps but missing from those type switches
// only fails at run time, once that version is live on chain. These tests feed every type the
// maps can produce, so a new version cannot be half-wired again.

func TestSetExecParamsHandlesEveryMappedVersion(t *testing.T) {
	for name, m := range map[string]map[string]func() typegen.CBORUnmarshaler{
		"Exec":  execParams,
		"Exec4": exec4Params,
	} {
		for version, newParams := range m {
			_, _, err := setExecParams(newParams())
			require.NotErrorIs(t, err, actors.ErrUnsupportedHeight, "%s params at %s (%T) are not handled", name, version, newParams())
		}
	}
}

func TestSetReturnParamsHandlesEveryMappedVersion(t *testing.T) {
	idAddr, err := address.NewIDAddress(1234)
	require.NoError(t, err)

	for name, m := range map[string]map[string]func() typegen.CBORUnmarshaler{
		"Exec":  execReturn,
		"Exec4": exec4Return,
	} {
		for version, newReturn := range m {
			ret := newReturn()
			f := reflect.ValueOf(ret).Elem().FieldByName("IDAddress")
			require.True(t, f.IsValid(), "%s return at %s: %T has no IDAddress", name, version, ret)
			f.Set(reflect.ValueOf(idAddr))

			info := setReturnParams(&parser.LotusMessage{}, "", ret)
			require.Equal(t, idAddr.String(), info.Short, "%s return at %s (%T) is not handled", name, version, ret)
		}
	}
}
