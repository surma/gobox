package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func init() {
	// "gobox" has to be added here for two reasons
	// 1.) So it can't be removed, it's core functionalty
	// 2.) It causes cyclic dependencies otherwise
	Applets["gobox"] = GoboxMain
}

func main() {
	callname := filepath.Base(os.Args[0])
	defer func() {
		if p := recover(); p != nil {
			e, ok := p.(error)
			if !ok {
				e = fmt.Errorf("Some error occured")
			}
			log.Fatalf("%s: %s", callname, e)
		}
	}()
	applet, ok := Applets[callname]
	if !ok {
		log.Fatalf("Could not find applet \"%s\"", callname)
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

	applet(args)
}
