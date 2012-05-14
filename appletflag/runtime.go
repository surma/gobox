package appletflag

import (
	"runtime"
	"strings"
)

func getCallerPackage() string {
	ptr, _, _, ok := runtime.Caller(2)
	if !ok {
		panic("Could not obtain caller’s function pointer")
	}
	name := runtime.FuncForPC(ptr).Name()
	return strings.Split(name, ".")[0]
}
