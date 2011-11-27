package chroot

import (
	"flag"
	"os"
	"syscall"
)

var (
	flagSet  = flag.NewFlagSet("chroot", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Chroot(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() < 2 || *helpFlag {
		println("`chroot` [options] <new root> <command>")
		flagSet.PrintDefaults()
		return nil
	}

	e = syscall.Chroot(flagSet.Arg(0))
	if e != nil {
		return e
	}

	e = syscall.Exec(flagSet.Arg(1), flagSet.Args()[1:], os.Environ())
	return e
}
