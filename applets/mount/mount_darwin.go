package mount

import (
	"errors"
	"flag"

	//"syscall"
	//"strings"
)

var (
	flagSet   = flag.NewFlagSet("mount", flag.PanicOnError)
	typeFlag  = flagSet.String("t", "", "Filesystem type of the mount")
	flagsFlag = flagSet.String("o", "", "Comma-separated list of flags for the mount (ro, noexec, nosuid, nodev, synchronous, remount)")
	helpFlag  = flagSet.Bool("help", false, "Show this help")
)

func Mount(call []string) error {
	/*e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 2 || *helpFlag {
		println("`mount` [options] <device> <dir>")
		flagSet.PrintDefaults()
		return nil
	}

	return nil*/
	return errors.New("`mount` not implemented for darwin")
}
