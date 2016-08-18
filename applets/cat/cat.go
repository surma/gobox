package cat

import (
	"flag"

	"io"
	"log"
	"os"
)

var (
	flagSet  = flag.NewFlagSet("cat", flag.PanicOnError)
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Cat(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flag.NArg() <= 0 || *helpFlag {
		println("`cat` [options] <files>")
		flag.PrintDefaults()
		return nil
	}

	for _, file := range flag.Args() {
		e := dumpFile(file)
		if e != nil {
			log.Printf("Could not print file %s: %s\n", file, e)
		}
	}

	return nil
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
