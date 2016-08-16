package rm

import (
	flag "../../appletflag"
	"io/ioutil"
	"log"
	"os"
)

var (
	recursiveFlag = flag.Bool("r", false, "Recurse into directories")
	helpFlag      = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() <= 0 || *helpFlag {
		println("`rm` [options] <files...>")
		flag.PrintDefaults()
		return
	}

	for _, file := range flag.Args() {
		e := delete(file)
		if e != nil {
			log.Fatalf("Could not delete file: %s\n", e)
		}
	}
	return
}

func delete(file string) error {
	fi, e := os.Stat(file)
	if e != nil {
		return e
	}
	if fi.IsDir() && *recursiveFlag {
		e := deleteDir(file)
		if e != nil {
			return e
		}
	}
	return os.Remove(file)
}

func deleteDir(dir string) error {
	files, e := ioutil.ReadDir(dir)
	if e != nil {
		return e
	}
	for _, file := range files {
		e = delete(dir + "/" + file.Name())
		if e != nil {
			return e
		}
	}
	return nil
}
