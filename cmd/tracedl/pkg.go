package main

import (
	"fmt"
	"reflect"
)

// FunctionInfo holds information about an exported function
type FunctionInfo struct {
	Name        string
	NumParams   int
	ParamTypes  []reflect.Type
	ReturnTypes []reflect.Type
}

// PackageExporter handles discovery and execution of exported functions
type PackageExporter struct {
	pkg           interface{}
	functionCache map[string]reflect.Value
	functionInfo  map[string]FunctionInfo
}

// NewPackageExporter creates a new PackageExporter instance
func NewPackageExporter(pkg interface{}) *PackageExporter {
	pe := &PackageExporter{
		pkg:           pkg,
		functionCache: make(map[string]reflect.Value),
		functionInfo:  make(map[string]FunctionInfo),
	}
	pe.discoverFunctions()
	return pe
}

// discoverFunctions finds all exported functions and caches their information
func (pe *PackageExporter) discoverFunctions() {
	pkgType := reflect.TypeOf(pe.pkg)
	pkgValue := reflect.ValueOf(pe.pkg)

	for i := 0; i < pkgType.NumMethod(); i++ {
		method := pkgType.Method(i)
		if !method.IsExported() {
			continue
		}

		methodType := method.Type
		info := FunctionInfo{
			Name:        method.Name,
			NumParams:   methodType.NumIn() - 1, // subtract 1 for receiver
			ParamTypes:  make([]reflect.Type, methodType.NumIn()-1),
			ReturnTypes: make([]reflect.Type, methodType.NumOut()),
		}

		// Get parameter types
		for j := 1; j < methodType.NumIn(); j++ { // start from 1 to skip receiver
			info.ParamTypes[j-1] = methodType.In(j)
		}

		// Get return types
		for j := 0; j < methodType.NumOut(); j++ {
			info.ReturnTypes[j] = methodType.Out(j)
		}

		pe.functionInfo[method.Name] = info
		pe.functionCache[method.Name] = pkgValue.Method(i)
	}
}

// GetExportedFunctions returns information about all exported functions
func (pe *PackageExporter) GetExportedFunctions() []FunctionInfo {
	functions := make([]FunctionInfo, 0, len(pe.functionInfo))
	for _, info := range pe.functionInfo {
		functions = append(functions, info)
	}
	return functions
}

// CallFunction calls an exported function by name with the provided arguments
func (pe *PackageExporter) CallFunction(functionName string, args ...interface{}) ([]interface{}, error) {
	// Check if function exists
	method, exists := pe.functionCache[functionName]
	if !exists {
		return nil, fmt.Errorf("function %s not found", functionName)
	}

	info := pe.functionInfo[functionName]

	// Check number of arguments
	if len(args) != info.NumParams {
		return nil, fmt.Errorf("function %s expects %d arguments, got %d",
			functionName, info.NumParams, len(args))
	}

	// Convert and type check arguments
	reflectArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValue := reflect.ValueOf(arg)
		expectedType := info.ParamTypes[i]

		if !argValue.Type().AssignableTo(expectedType) {
			return nil, fmt.Errorf("argument %d type mismatch: expected %v, got %v",
				i, expectedType, argValue.Type())
		}
		reflectArgs[i] = argValue
	}

	// Call the function
	results := method.Call(reflectArgs)

	// Convert results to interface slice
	returnValues := make([]interface{}, len(results))
	for i, result := range results {
		returnValues[i] = result.Interface()
	}

	// Check if the last return value is an error
	if len(results) > 0 {
		lastResult := results[len(results)-1]
		if lastResult.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			if !lastResult.IsNil() {
				return returnValues[:len(returnValues)-1], lastResult.Interface().(error)
			}
		}
	}

	return returnValues, nil
}
