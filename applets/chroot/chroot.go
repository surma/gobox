package chroot

import (
	"flag"

	"log"
	"os"
	"syscall"
)

var (
	flagSet  = flag.NewFlagSet("chroot", flag.PanicOnError)
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Chroot(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flag.NArg() < 2 || *helpFlag {
		println("`chroot` [options] <new root> <command>")
		flag.PrintDefaults()
		return nil
	}

	e = syscall.Chroot(flag.Arg(0))
	if e != nil {
		log.Fatalf("Could not chroot: %s\n", e)
	}

	e = syscall.Exec(flag.Arg(1), flag.Args()[1:], os.Environ())
	if e != nil {
		log.Fatalf("Could not exec: %s\n", e)
	}

	return nil
}
