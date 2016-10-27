package mkdir

import (
	"flag"

	"log"
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
		e = os.MkdirAll(flagSet.Arg(0), 0755)
	} else {
		e = os.Mkdir(flagSet.Arg(0), 0755)
	}
	if e != nil {
		log.Fatalf("Could not create: %s\n", e)
	}

	return nil
}
