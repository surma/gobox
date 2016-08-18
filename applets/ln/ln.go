package ln

import (
	"flag"

	"log"
	"os"
)

var (
	flagSet    = flag.NewFlagSet("ln", flag.PanicOnError)
	parentFlag = flagSet.Bool("s", false, "Create a symlink")
	helpFlag   = flagSet.Bool("help", false, "Show this help")
)

func Ln(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 2 || *helpFlag {
      println("`Ln` [OPTION]... TARGET... DIRECTORY")
		flagSet.PrintDefaults()
		return nil
	}

	if *parentFlag {
      e = os.Symlink(flagSet.Arg(0), flagSet.Arg(1))
	} else {
      log.Fatalf("Not implemented. Use -s")
    }
	if e != nil {
		log.Fatalf("Could not create: %s\n", e)
	}

	return nil
}
