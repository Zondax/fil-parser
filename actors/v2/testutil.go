package v2

import (
	"reflect"
	"slices"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/rt"
)

var utilityMethods = []string{"Exports", "Code", "IsSingleton", "State"}

func runtimeActorMethods(actor rt.VMActor) []string {
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

func builtinActorMethods(methodMap map[abi.MethodNum]builtin.MethodMeta) []string {
	methods := []string{}
	for _, method := range methodMap {
		methods = append(methods, method.Name)
	}
	return methods
}

func MissingMethods(actor Actor, actorVersions []any) []string {
	missingMethods := map[string]bool{}
	txTypes := actor.TransactionTypes()
	for _, actorVersion := range actorVersions {
		switch actorVersion := actorVersion.(type) {
		case rt.VMActor:
			methods := runtimeActorMethods(actorVersion)
			for _, method := range methods {
				_, ok := txTypes[method]
				if !ok {
					missingMethods[method] = true
				}
			}
		case map[abi.MethodNum]builtin.MethodMeta:
			methods := builtinActorMethods(actorVersion)
			for _, method := range methods {
				_, ok := txTypes[method]
				if !ok {
					missingMethods[method] = true
				}
			}
		}
	}
	missingMethodsList := []string{}
	for method := range missingMethods {
		missingMethodsList = append(missingMethodsList, method)
	}

	return missingMethodsList
}
