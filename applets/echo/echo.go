package echo

import (
	flag "appletflag"
	"fmt"
	"strings"
)

// A simple, parameterless echo implementation
func Main() {
	parameters := flag.Parameters[1:]
	fmt.Println(strings.Join(parameters, " "))
	return
}
