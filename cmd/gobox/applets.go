package main

import (
	// Applet imports: start
	"applets/echo"
	// Applet imports: end
	"os"
)

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) os.Error

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet {
	// Applet functions: start
	"echo": echo.Echo,
	// Applet functions: end
}


