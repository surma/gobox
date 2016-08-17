package mkdir

import (
	flag "gobox/appletflag"
	"log"
	"os"
)

var (
	parentFlag = flag.Bool("p", false, "Create parent directories, if necessary")
	helpFlag   = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() != 1 || *helpFlag {
		println("`Mkdir` [options] <path>")
		flag.PrintDefaults()
		return
	}

	var e error
	if *parentFlag {
		e = os.MkdirAll(flag.Arg(0), 0755)
	} else {
		e = os.Mkdir(flag.Arg(0), 0755)
	}
	if e != nil {
		log.Fatalf("Could not create: %s\n", e)
	}
}
