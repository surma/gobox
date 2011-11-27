package mkdir

import (
	"flag"
	"os"
)

var (
	flagSet    = flag.NewFlagSet("mkdir", flag.PanicOnError)
	parentFlag = flagSet.Bool("p", false, "Create parent directories, if necessary")
	helpFlag   = flagSet.Bool("help", false, "Show this help")
)

func Mkdir(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 1 || *helpFlag {
		println("`Mkdir` [options] <path>")
		flagSet.PrintDefaults()
		return nil
	}

	if *parentFlag {
		return os.MkdirAll(flagSet.Arg(0), 0755)
	}
	return os.Mkdir(flagSet.Arg(0), 0755)
}
