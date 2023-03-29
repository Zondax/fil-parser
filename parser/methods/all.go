package methods

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
)

func All() map[string]func() map[string]map[abi.MethodNum]builtin.MethodMeta {
	return map[string]func() map[string]map[abi.MethodNum]builtin.MethodMeta{
		"v8":  V8Methods,
		"v9":  V9Methods,
		"v10": V10Methods,
		"v11": V11Methods,
	}
}
