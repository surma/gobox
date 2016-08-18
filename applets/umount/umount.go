package umount

import (
	"flag"

	"log"
	"syscall"
)

var (
	flagSet  = flag.NewFlagSet("umount", flag.PanicOnError)
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Umount(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flag.NArg() != 1 || *helpFlag {
		println("`umount` [options] <mount point>")
		flag.PrintDefaults()
		return nil
	}

	e = syscall.Unmount(flag.Arg(0), 0)
	if e != nil {
		log.Fatalf("Could not unmount: %s\n", e)
	}
	return e
}
