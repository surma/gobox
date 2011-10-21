package main

import (
	"os"
)

// Applet imports
import (
	"applets/echo"
	"applets/shell"
	"applets/telnetd"
	"applets/ls"
	"applets/rm"
	"applets/httpd"
	"applets/wget"
	"applets/kill"
	"applets/cat"
	"applets/mknod"
	"applets/mount"
	"applets/umount"
	"applets/chroot"
	"applets/ps"
	"applets/mkdir"
	"applets/head"
)

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet{
	"echo":    echo.Echo,
	"shell":   shell.Shell,
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
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) os.Error
