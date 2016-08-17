package main

// Applet imports
import (
    "gobox/applets/cal"
	"gobox/applets/cat"
	"gobox/applets/chroot"
    "gobox/applets/checksum"
    "gobox/applets/date"
	"gobox/applets/echo"
	"gobox/applets/grep"
	"gobox/applets/gzip"
	"gobox/applets/head"
	"gobox/applets/httpd"
	"gobox/applets/kill"
	"gobox/applets/ls"
	"gobox/applets/mkdir"
	"gobox/applets/mknod"
	"gobox/applets/mount"
	"gobox/applets/ps"
	"gobox/applets/rm"
	"gobox/applets/shell"
    "gobox/applets/strings"
	"gobox/applets/telnetd"
    "gobox/applets/touch"
    "gobox/applets/tar"
	"gobox/applets/umount"
    "gobox/applets/wc"
	"gobox/applets/wget"
)

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet{
	"echo":      echo.Echo,
	"shell":     shell.Shell,
    "sh":        shell.Shell,
	"strings":   strings.Strings,
	"telnetd":   telnetd.Telnetd,
	"touch":     touch.Touch,
	"md5sum":    checksum.Hash,
	"sha1sum":   checksum.Hash,
	"sha256sum": checksum.Hash,
	"sha512sum": checksum.Hash,
	"crc32":     checksum.Hash,
	"ls":        ls.Ls,
	"rm":        rm.Rm,
	"httpd":     httpd.Httpd,
	"wget":      wget.Wget,
	"tar":       tar.Tar,
	"wc":        wc.Wc,
	"kill":      kill.Kill,
	"cat":       cat.Cat,
	"cal":       cal.Cal,
	"mknod":     mknod.Mknod,
	"mount":     mount.Mount,
	"umount":    umount.Umount,
	"chroot":    chroot.Chroot,
	"ps":        ps.Ps,
	"mkdir":     mkdir.Mkdir,
	"head":      head.Head,
	"grep":      grep.Grep,
	"gzip":      gzip.Gzip,
	"gunzip":    gzip.Gunzip,
	"zcat":      gzip.Zcat,
	"date":      date.Date,
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) error
