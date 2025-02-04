package v2

import (
	"reflect"
	"slices"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/rt"
)

var utilityMethods = []string{"Exports", "Code", "IsSingleton", "State"}

func RuntimeActorMethods(actor rt.VMActor) []string {
	t := reflect.TypeOf(actor)
	methods := []string{}
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if slices.Contains(utilityMethods, method.Name) {
			continue
		}
		methods = append(methods, method.Name)
	}
	return methods
}

func BuiltinActorMethods(methodMap map[abi.MethodNum]builtin.MethodMeta) []string {
	methods := []string{}
	for _, method := range methodMap {
		methods = append(methods, method.Name)
	}
	return methods
}
