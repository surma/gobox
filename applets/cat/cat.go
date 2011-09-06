package cat

import (
	"flag"
	"os"
	"io"
)

var (
	flagSet  = flag.NewFlagSet("template", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Cat(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`cat` [options] <files>")
		flagSet.PrintDefaults()
		return nil
	}

	for _, file := range flagSet.Args() {
		e := dumpFile(file)
		if e != nil {
			return e
		}
	}

	return nil
}

func dumpFile(path string) os.Error {
	f, e := os.Open(path)
	if e != nil {
		return e
	}
	defer f.Close()

	_, e = io.Copy(os.Stdout, f)
	return e
}
