package cat

import (
	flag "gobox/appletflag"
	"io"
	"log"
	"os"
)

var (
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() <= 0 || *helpFlag {
		println("`cat` [options] <files>")
		flag.PrintDefaults()
		return
	}

	for _, file := range flag.Args() {
		e := dumpFile(file)
		if e != nil {
			log.Printf("Could not print file %s: %s\n", file, e)
		}
	}

	return
}

func dumpFile(path string) error {
	f, e := os.Open(path)
	if e != nil {
		return e
	}
	defer f.Close()

	_, e = io.Copy(os.Stdout, f)
	return e
}
