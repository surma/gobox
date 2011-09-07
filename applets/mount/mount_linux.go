package mount

import (
	"flag"
	"os"
	"syscall"
	"strings"
)

var (
	flagSet   = flag.NewFlagSet("mount", flag.PanicOnError)
	typeFlag  = flagSet.String("t", "", "Filesystem type of the mount")
	flagsFlag = flagSet.String("o", "defaults", "Comma-separated list of flags for the mount")
	helpFlag  = flagSet.Bool("help", false, "Show this help")
)

func Mount(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 2 || *helpFlag {
		println("`mount` [options] <device> <dir>")
		flagSet.PrintDefaults()
		println("\nAvailable options are:")
		for opt, _ := range flagMap {
			print(opt, ", ")
		}
		println()
		return nil
	}

	flags, e := parseFlags()
	if e != nil {
		return e
	}

	errno := syscall.Mount(flagSet.Arg(0), flagSet.Arg(1), *typeFlag, flags, "")
	if errno != 0 {
		return os.NewError(syscall.Errstr(errno))
	}
	return nil
}

var (
	flagMap = map[string]int{
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

func parseFlags() (int, os.Error) {
	ret := 0
	parts := strings.Split(*flagsFlag, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		val, ok := flagMap[strings.ToLower(part)]
		if !ok {
			return 0, os.NewError("Invalid flag \"" + part + "\"")
		}
		ret |= val
	}
	return ret, nil
}
