package util

import (
	"reflect"
	"runtime"
	"strings"
)

func FuncName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	if idx := strings.LastIndex(fullName, "."); idx != -1 {
		name := fullName[idx+1:]
		if strings.HasSuffix(name, "-fm") {
			return name[:len(name)-3]
		}
		return name
	}

	return fullName
}
