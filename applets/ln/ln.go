package ln

import (
	"flag"

	"log"
	"os"
)

var (
	flagSet    = flag.NewFlagSet("ln", flag.PanicOnError)
	parentFlag = flag.Bool("s", false, "Create a symlink")
	helpFlag   = flag.Bool("help", false, "Show this help")
)

func Ln(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flag.NArg() != 2 || *helpFlag {
      println("`Ln` [OPTION]... TARGET... DIRECTORY")
		flag.PrintDefaults()
		return nil
	}

	if *parentFlag {
      e = os.Symlink(flag.Arg(0), flag.Arg(1))
	} else {
      log.Fatalf("Not implemented. Use -s")
    }
	if e != nil {
		log.Fatalf("Could not create: %s\n", e)
	}

	return nil
}
