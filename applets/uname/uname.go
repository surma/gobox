package uname

import (
	"flag"
	"fmt"
	"syscall"
)

type Utsname syscall.Utsname

var (
	flagSet = flag.NewFlagSet("uname", flag.PanicOnError)
	aflag   = flagSet.Bool("a", false, "All info.")
	iflag   = flagSet.Bool("i", false, "")
	mflag   = flagSet.Bool("m", false, "")
	nflag   = flagSet.Bool("n", false, "")
	rflag   = flagSet.Bool("r", false, "")
	vflag   = flagSet.Bool("v", false, "")
	sflag   = flagSet.Bool("s", false, "")
	hflag   = flagSet.Bool("h", false, "Print this help")
)

func b2S(bs [65]int8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

func out(space bool, a string) {
	if space {
		fmt.Print(" ")
	}
	fmt.Print(a)
}

func Uname(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if *hflag {
		println("`Uname` [OPTION]...")
		flagSet.PrintDefaults()
		return nil
	}

	uts := &syscall.Utsname{}
	if err := syscall.Uname(uts); err != nil {
		return err
	}

	space := false
	if *sflag || *aflag {
		out(space, b2S(uts.Sysname))
		space = true
	}
	if *nflag || *aflag {
		out(space, b2S(uts.Nodename))
		space = true
	}
	if *rflag || *aflag {
		out(space, b2S(uts.Release))
		space = true
	}
	if *vflag || *aflag {
		out(space, b2S(uts.Version))
		space = true
	}
	if *mflag || *aflag {
		out(space, b2S(uts.Machine))
		space = true
	}
	if *iflag || *aflag {
		out(space, b2S(uts.Domainname))
		space = true
	}
	fmt.Print("\n")
	return nil
}
