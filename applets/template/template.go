package template

import (
	"flag"
	"os"
)

var (
	flagSet  = flag.NewFlagSet("template", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Template(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`template` [options] <parameters>")
		flagSet.PrintDefaults()
		return nil
	}
	return nil
}
