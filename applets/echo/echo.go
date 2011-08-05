package echo

import (
	"strings"
	"os"
)

// A simple, parameterless echo implementation
func Echo(call []string) os.Error {
	var parameters []string
	if len(call) <= 1 {
		parameters = make([]string, 0)
	} else {
		parameters = call[1:]
	}

	println(strings.Join(parameters, " "))
	return nil
}
