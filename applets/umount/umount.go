package umount

import (
	"flag"

	"log"
	"syscall"
)

var (
	flagSet  = flag.NewFlagSet("umount", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Umount(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 1 || *helpFlag {
		println("`umount` [options] <mount point>")
		flagSet.PrintDefaults()
		return nil
	}

	e = syscall.Unmount(flagSet.Arg(0), 0)
	if e != nil {
		log.Fatalf("Could not unmount: %s\n", e)
	}
	return e
}
