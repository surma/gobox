package umount

import (
	flag "appletflag"
	"log"
	"syscall"
)

var (
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() != 1 || *helpFlag {
		println("`umount` [options] <mount point>")
		flag.PrintDefaults()
		return
	}

	e := syscall.Unmount(flag.Arg(0), 0)
	if e != nil {
		log.Fatalf("Could not unmount: %s\n", e)
	}
}
