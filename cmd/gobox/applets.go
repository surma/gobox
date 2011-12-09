package main

// Applet imports
import (
	// "applets/cat"
	// "applets/chroot"
	// "applets/echo"
	// "applets/grep"
	// "applets/gzip"
	// "applets/head"
	// "applets/httpd"
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
	// "echo":     echo.Echo,
	// "httpd":    httpd.Httpd,
	// "cat":      cat.Cat,
	// "chroot":   chroot.Chroot,
	// "head":     head.Head,
	// "grep":     grep.Grep,
	// "gzip":     gzip.Gzip,
	// "gunzip":   gzip.Gunzip,
	// "zcat":     gzip.Zcat,
	"ifconfig": ifconfig.Main,
	"kill":     kill.Main,
	"ls":       ls.Main,
	"mkdir":    mkdir.Main,
	"mknod":    mknod.Main,
	"mount":    mount.Main,
	"ps":       ps.Main,
	"rm":       rm.Main,
	"shell":    shell.Main,
	"telnetd":  telnetd.Main,
	"umount":   umount.Main,
	"wget":     wget.Main,
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func()
