package main

import (
	flag "./appletflag"
	"./common"
	"errors"
	"os"
	"path/filepath"
)

func init() {
	// "gobox" has to be added here for two reasons
	// 1.) So it can't be removed, it's core functionalty
	// 2.) It causes cyclic dependencies otherwise
	Applets["gobox"] = GoboxMain
}

func run() {
	callname := filepath.Base(os.Args[0])
	applet, ok := Applets[callname]
	if !ok {
		panic(errors.New("Could not find applet \"" + callname + "\""))
	}

	// If the Gobox applet is called (i.e. the executable itself)
	// check, if the second parameter is an applet name.
	// If so, call that applet instead
	args := os.Args
	if callname == "gobox" && len(args) >= 2 {
		subapplet, ok := Applets[args[1]]
		if ok {
			applet = subapplet
			args = args[1:]
		}
	}

	flag.Parameters = args
	applet()
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			e, ok := p.(error)
			if !ok {
				e = errors.New("Some error occured")
			}
			common.DumpError(e)
		}
	}()
	run()
}
