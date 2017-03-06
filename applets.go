package main

// Applet imports
import (
	"github.com/surma/gobox/applets/cat"
	"github.com/surma/gobox/applets/chroot"
	"github.com/surma/gobox/applets/echo"
	"github.com/surma/gobox/applets/grep"
	"github.com/surma/gobox/applets/gzip"
	"github.com/surma/gobox/applets/head"
	"github.com/surma/gobox/applets/httpd"
	"github.com/surma/gobox/applets/kill"
	"github.com/surma/gobox/applets/ls"
	"github.com/surma/gobox/applets/mkdir"
	"github.com/surma/gobox/applets/mknod"
	"github.com/surma/gobox/applets/mount"
	"github.com/surma/gobox/applets/ps"
	"github.com/surma/gobox/applets/rm"
	"github.com/surma/gobox/applets/shell"
	"github.com/surma/gobox/applets/telnetd"
	"github.com/surma/gobox/applets/umount"
	"github.com/surma/gobox/applets/wget"
)

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet{
	"echo":    echo.Echo,
	"shell":   shell.Shell,
	"sh":      shell.Shell,
	"telnetd": telnetd.Telnetd,
	"ls":      ls.Ls,
	"rm":      rm.Rm,
	"httpd":   httpd.Httpd,
	"wget":    wget.Wget,
	"kill":    kill.Kill,
	"cat":     cat.Cat,
	"mknod":   mknod.Mknod,
	"mount":   mount.Mount,
	"umount":  umount.Umount,
	"chroot":  chroot.Chroot,
	"ps":      ps.Ps,
	"mkdir":   mkdir.Mkdir,
	"head":    head.Head,
	"grep":    grep.Grep,
	"gzip":    gzip.Gzip,
	"gunzip":  gzip.Gunzip,
	"zcat":    gzip.Zcat,
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) error
