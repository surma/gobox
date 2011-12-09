package echo

import (
	"fmt"
	flag "appletflag"
	"strings"
)

// A simple, parameterless echo implementation
func Main() {
	parameters := flag.Parameters
	fmt.Println(strings.Join(parameters, " "))
	return
}
