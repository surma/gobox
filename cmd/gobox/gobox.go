package main

import (
	"flag"
	"os"
	"common"
	"path/filepath"
)

var (
	helpFlag = flag.Bool("help", false, "Show help")
	listFlag = flag.Bool("list", false, "List applets")
	installFlag = flag.String("install", "", "Create symlinks for applets in given path")
)

func help() {
	flag.PrintDefaults()
	println()
	list()
}

func list() {
	println("List of compiled applets:\n")
	for name, _ := range Applets {
		print(name,", ")
	}
	println("")
}

func install(path string) {
	goboxpath, e := common.GetGoboxBinaryPath()
	if e != nil {
		common.DumpError(e)
		return
	}
	for name, _ := range Applets {
		newpath := filepath.Join(path, name)
		e = common.ForcedSymlink(goboxpath, newpath)
		if e != nil {
			common.DumpError(e)
		}
	}
}

func run() {
	callname := filepath.Base(os.Args[0])
	// "gobox" has to be handled separately. Putting it in
	// the Applets map results in dependency cylces
	if callname == "gobox" {
		help()
		return
	}
	applet, ok := Applets[callname]
	if !ok {
		panic(os.NewError("Could not find applet \""+callname+"\""))
	}
	applet(os.Args)
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
	flag.Parse()

	if *helpFlag {
		help()
	} else if *listFlag {
		list()
	} else if *installFlag != ""{
		install(*installFlag)
	} else {
		run()
	}

}
