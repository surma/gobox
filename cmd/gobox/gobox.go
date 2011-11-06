package main

import (
	"flag"
	"common"
	"path/filepath"
)

const (
	VERSION = "0.2"
)

var (
	flagSet     = flag.NewFlagSet("gobox", flag.ExitOnError)
	listFlag    = flagSet.Bool("list", false, "List applets")
	installFlag = flagSet.String("install", "", "Create symlinks for applets in given path")
	helpFlag    = flagSet.Bool("help", false, "Show help")
)

func Gobox(call []string) (e error) {
	e = flagSet.Parse(call[1:])
	if e != nil {
		return
	}

	if *listFlag {
		list()
	} else if *installFlag != "" {
		e = install(*installFlag)
	} else {
		help()
	}
	return
}

func help() {
	println("`gobox` [options]")
	flagSet.PrintDefaults()
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
