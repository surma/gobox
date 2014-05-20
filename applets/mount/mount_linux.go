package mount

import (
	"errors"
	"flag"
	"strings"
	"syscall"
)

var (
	flagSet   = flag.NewFlagSet("mount", flag.PanicOnError)
	typeFlag  = flagSet.String("t", "", "Filesystem type of the mount")
	flagsFlag = flagSet.String("o", "defaults", "Comma-separated list of flags for the mount")
	helpFlag  = flagSet.Bool("help", false, "Show this help")
)

func Mount(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 2 || *helpFlag {
		println("`mount` [options] <device> <dir>")
		flagSet.PrintDefaults()
		println("\nAvailable options are:")
		for opt := range flagMap {
			print(opt, ", ")
		}
		println()
		return nil
	}

	flags, e := parseFlags()
	if e != nil {
		return e
	}

	e = syscall.Mount(flagSet.Arg(0), flagSet.Arg(1), *typeFlag, uintptr(flags), "")
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
