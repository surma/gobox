package chroot

import (
	flag "appletflag"
	"log"
	"os"
	"syscall"
)

var (
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() < 2 || *helpFlag {
		println("`chroot` [options] <new root> <command>")
		flag.PrintDefaults()
		return
	}

	e := syscall.Chroot(flag.Arg(0))
	if e != nil {
		log.Fatalf("Could not chroot: %s\n", e)
	}

	e = syscall.Exec(flag.Arg(1), flag.Args()[1:], os.Environ())
	if e != nil {
		log.Fatalf("Could not exec: %s\n", e)
	}
}
