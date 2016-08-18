package mkdir

import (
	"flag"

	"log"
	"os"
)

var (
	flagSet    = flag.NewFlagSet("mkdir", flag.PanicOnError)
	parentFlag = flag.Bool("p", false, "Create parent directories, if necessary")
	helpFlag   = flag.Bool("help", false, "Show this help")
)

func Mkdir(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flag.NArg() != 1 || *helpFlag {
		println("`Mkdir` [options] <path>")
		flag.PrintDefaults()
		return nil
	}

	if *parentFlag {
		e = os.MkdirAll(flag.Arg(0), 0755)
	} else {
		e = os.Mkdir(flag.Arg(0), 0755)
	}
	if e != nil {
		log.Fatalf("Could not create: %s\n", e)
	}

	return nil
}
