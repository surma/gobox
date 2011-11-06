package chroot

import (
	"errors"
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

	errno := syscall.Chroot(flagSet.Arg(0))
	if errno != 0 {
		return errors.New(syscall.Errstr(errno))
	}

	errno = syscall.Exec(flagSet.Arg(1), flagSet.Args()[1:], os.Envs)
	if errno != 0 {
		return errors.New(syscall.Errstr(errno))
	}
	return nil
}
