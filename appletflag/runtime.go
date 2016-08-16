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
	elems := strings.Split(name, ".")
	return strings.Join(elems[0:len(elems)-1], ".")
}
