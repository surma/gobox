package main

// Applet imports
import (
	"applets/cat"
	"applets/chroot"
	"applets/echo"
	"applets/grep"
	"applets/gzip"
	"applets/head"
	"applets/httpd"
	"applets/ifconfig"
	"applets/kill"
	"applets/ls"
	"applets/mkdir"
	"applets/mknod"
	"applets/mount"
	"applets/ps"
	"applets/rm"
	"applets/shell"
	"applets/telnetd"
	"applets/umount"
	"applets/wget"
)

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet{
	"echo":     echo.Echo,
	"shell":    shell.Shell,
	"telnetd":  telnetd.Telnetd,
	"ls":       ls.Ls,
	"rm":       rm.Rm,
	"httpd":    httpd.Httpd,
	"wget":     wget.Wget,
	"kill":     kill.Kill,
	"cat":      cat.Cat,
	"ifconfig": ifconfig.Ifconfig,
	"mknod":    mknod.Mknod,
	"mount":    mount.Mount,
	"umount":   umount.Umount,
	"chroot":   chroot.Chroot,
	"ps":       ps.Ps,
	"mkdir":    mkdir.Mkdir,
	"head":     head.Head,
	"grep":     grep.Grep,
	"gzip":     gzip.Gzip,
	"gunzip":   gzip.Gunzip,
	"zcat":     gzip.Zcat,
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) error
