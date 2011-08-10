package rm

import (
	"os"
	"flag"
	"io/ioutil"
)

var (
	flagSet       = flag.NewFlagSet("rm", flag.PanicOnError)
	recursiveFlag = flagSet.Bool("r", false, "Recurse into directories")
	helpFlag      = flagSet.Bool("help", false, "Show this help")
)

func Rm(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`rm` [options] <files...>")
		flagSet.PrintDefaults()
		return nil
	}

	for _, file := range flagSet.Args() {
		e := delete(file)
		if e != nil {
			return e
		}
	}
	return nil
}

func delete(file string) os.Error {
	fi, e := os.Stat(file)
	if e != nil {
		return e
	}
	if fi.IsDirectory() && *recursiveFlag {
		e := deleteDir(file)
		if e != nil {
			return e
		}
	}
	return os.Remove(file)
}

func deleteDir(dir string) os.Error {
	files, e := ioutil.ReadDir(dir)
	if e != nil {
		return e
	}
	for _, file := range files {
		e = delete(dir + "/" + file.Name)
		if e != nil {
			return e
		}
	}
	return nil
}
