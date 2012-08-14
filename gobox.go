package main

import (
	flags "flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	VERSION = "0.3"
)

var (
	flag        = flags.NewFlagSet("Gobox", flags.ExitOnError)
	listFlag    = flag.Bool("list", false, "List applets")
	installFlag = flag.String("install", "", "Create symlinks for applets in given path")
	helpFlag    = flag.Bool("help", false, "Show help")
)

func GoboxMain(call []string) {
	flag.Parse(call[1:])

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
	fmt.Println("`gobox` [options]")
	flag.PrintDefaults()
	fmt.Println("Version", VERSION)
	list()
}

func list() {
	fmt.Println("List of available applets:")
	sep := ""
	for name, _ := range Applets {
		fmt.Print(sep, name)
		sep = ", "
	}
	fmt.Println("")
}

func install(path string) error {
	goboxpath, e := binaryLocation()
	if e != nil {
		return e
	}
	for name, _ := range Applets {
		// Don't overwrite the executable
		if name == "gobox" {
			continue
		}
		newpath := filepath.Join(path, name)
		e = os.Symlink(goboxpath, newpath)
		if e != nil {
			log.Printf("Symlinking %s failed: %s", newpath, e)
			continue
		}
	}
	return nil
}

func binaryLocation() (string, error) {
	callname := os.Args[0]
	path, e := exec.LookPath(callname)
	if e == nil {
		return filepath.Abs(path)
	}

	return "", fmt.Errorf("Could not find gobox binary")
}
