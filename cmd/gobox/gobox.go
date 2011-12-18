package main

import (
	flag "appletflag"
	"common"
	"path/filepath"
)

const (
	VERSION = "0.3"
)

var (
	listFlag    = flag.Bool("list", false, "List applets")
	installFlag = flag.String("install", "", "Create symlinks for applets in given path")
	helpFlag    = flag.Bool("help", false, "Show help")
)

func GoboxMain() {
	flag.Parse()

	if *listFlag {
		list()
	} else if *installFlag != "" {
		_ = install(*installFlag)
	} else {
		help()
	}
	return
}

func help() {
	println("`gobox` [options]")
	flag.PrintDefaults()
	println()
	println("Version", VERSION)
	list()
}

func list() {
	println("List of compiled applets:\n")
	for name, _ := range Applets {
		print(name, ", ")
	}
	println("")
}

func install(path string) error {
	goboxpath, e := common.GetGoboxBinaryPath()
	if e != nil {
		return e
	}
	for name, _ := range Applets {
		// Don't overwrite the executable
		if name == "gobox" {
			continue
		}
		newpath := filepath.Join(path, name)
		e = common.ForcedSymlink(goboxpath, newpath)
		if e != nil {
			common.DumpError(e)
		}
	}
	return nil
}
