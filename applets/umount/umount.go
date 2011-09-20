package umount

import (
	"flag"
	"os"
	"syscall"
)

var (
	flagSet  = flag.NewFlagSet("umount", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Umount(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 1 || *helpFlag {
		println("`umount` [options] <mount point>")
		flagSet.PrintDefaults()
		return nil
	}

	errno := syscall.Unmount(flagSet.Arg(0), 0)
	if errno != 0 {
		return os.NewError(syscall.Errstr(errno))
	}
	return nil
}
