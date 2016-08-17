package mknod

import (
	flag "gobox/appletflag"
	"log"
	"strconv"
	"strings"
	"syscall"
)

var (
	majorFlag = flag.Int("major", -1, "Major number of the block device")
	minorFlag = flag.Int("minor", -1, "Minor number of the block device")
	typeFlag  = flag.String("type", "", "Type of the node to create (i.e. socket, link, regular, block, directory, character, fifo)")
	modeFlag  = flag.String("mode", "644", "Mode (in octal) to create the node with")
	helpFlag  = flag.Bool("help", false, "Show this help")

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

func Main() {
	flag.Parse()

	if flag.NArg() != 1 || *helpFlag {
		println("`mknod` [options] <file>")
		flag.PrintDefaults()
		return
	}

	mode, ok := typemap[strings.ToLower(*typeFlag)]
	if !ok {
		log.Fatalf("Invalid node type: %s\n", *typeFlag)
	}

	if mode == syscall.S_IFBLK && (*majorFlag == -1 || *minorFlag == -1) {
		log.Fatalf("Major and minor device have to be set when creating a block device\n")
	}

	fmode, e := strconv.ParseInt(*modeFlag, 8, 32)
	if e != nil {
		log.Fatalf("Invalid number: %s\n", e)
	}
	mode |= uint32(fmode)

	e = syscall.Mknod(flag.Arg(0), mode, *majorFlag<<8|*minorFlag)
	if e != nil {
		log.Fatalf("Could not create node: %s\n", e)
	}
}
