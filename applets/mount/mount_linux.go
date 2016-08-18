package mount

import (
	"flag"

	"errors"
	"log"
	"strings"
	"syscall"
)

var (
	flagSet   = flag.NewFlagSet("mount_linux", flag.PanicOnError)
	typeFlag  = flag.String("t", "", "Filesystem type of the mount")
	flagsFlag = flag.String("o", "defaults", "Comma-separated list of flags for the mount")
	helpFlag  = flag.Bool("help", false, "Show this help")
)

func Mount(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flag.NArg() != 2 || *helpFlag {
		println("`mount` [options] <device> <dir>")
		flag.PrintDefaults()
		println("\nAvailable options are:")
		for opt := range flagMap {
			print(opt, ", ")
		}
		println()
		return nil
	}

	flags, e := parseFlags()
	if e != nil {
		log.Fatalf("Could not parse options: %s\n", e)
	}

	e = syscall.Mount(flag.Arg(0), flag.Arg(1), *typeFlag, uintptr(flags), "")
	if e != nil {
		log.Fatalf("Could not mount: %s\n", e)
	}
	return e
}

var (
	flagMap = map[string]uint32{
		"defaults":   0,
		"noatime":    syscall.MS_NOATIME,
		"nodev":      syscall.MS_NODEV,
		"nodiratime": syscall.MS_NODIRATIME,
		"noexec":     syscall.MS_NOEXEC,
		"nosuid":     syscall.MS_NOSUID,
		"remount":    syscall.MS_REMOUNT,
		"ro":         syscall.MS_RDONLY,
		"sync":       syscall.MS_SYNCHRONOUS,
	}
)

func parseFlags() (uint32, error) {
	ret := uint32(0)
	parts := strings.Split(*flagsFlag, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		val, ok := flagMap[strings.ToLower(part)]
		if !ok {
			return 0, errors.New("Invalid flag \"" + part + "\"")
		}
		ret |= val
	}
	return ret, nil
}
