package mknod

import (
	"flag"
	"os"
	"syscall"
	"strings"
)

var (
	flagSet   = flag.NewFlagSet("mknod", flag.PanicOnError)
	majorFlag = flagSet.Int("major", -1, "Major number of the block device")
	minorFlag = flagSet.Int("minor", -1, "Minor number of the block device")
	typeFlag  = flagSet.String("type", "", "Type of the node to create (i.e. socket, link, regular, block, directory, character, fifo)")
	modeFlag  = flagSet.Int("mode", 0644, "Mode to create the node with")
	helpFlag  = flagSet.Bool("help", false, "Show this help")

	typemap = map[string]uint32{
		"socket":    syscall.S_IFSOCK,
		"link":      syscall.S_IFLNK,
		"regular":   syscall.S_IFREG,
		"block":     syscall.S_IFBLK,
		"directory": syscall.S_IFDIR,
		"character": syscall.S_IFCHR,
		"fifo":      syscall.S_IFIFO,
	}
)

func Mknod(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`mknod` [options] <file>")
		flagSet.PrintDefaults()
		return nil
	}

	mode, ok := typemap[strings.ToLower(*typeFlag)]
	if !ok {
		return os.NewError("Invalid node type \"" + *typeFlag + "\"")
	}

	if mode == syscall.S_IFBLK && (*majorFlag == -1 || *minorFlag == -1) {
		return os.NewError("When creating a block device, both minor and major number have to be given")
	}
	mode |= uint32(*modeFlag)

	errno := syscall.Mknod(flagSet.Arg(0), mode, *majorFlag<<16|*minorFlag)
	if errno != 0 {
		return os.NewError(syscall.Errstr(errno))
	}

	return nil
}
