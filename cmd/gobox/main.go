package main

import (
	"path/filepath"
	"os"
	"common"
)

func init() {
	// "gobox" has to be added here for two reasons
	// 1.) So it can't be removed, it's core functionalty
	// 2.) It causes cyclic dependencies
	Applets["gobox"] = Gobox
}

func run() {
	callname := filepath.Base(os.Args[0])
	applet, ok := Applets[callname]
	if !ok {
		panic(os.NewError("Could not find applet \"" + callname + "\""))
	}
	e := applet(os.Args)
	if e != nil {
		panic(e)
	}
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			e, ok := p.(os.Error)
			if !ok {
				e = os.NewError("Some error occured")
			}
			common.DumpError(e)
		}
	}()
	run()
}
