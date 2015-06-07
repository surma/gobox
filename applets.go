package main

// Applet imports
import (
<<<<<<< HEAD
	"./applets/cat"
	"./applets/chroot"
    "./applets/checksum"
    "./applets/date"
	"./applets/echo"
	"./applets/grep"
	"./applets/gzip"
	"./applets/head"
	"./applets/httpd"
	"./applets/kill"
	"./applets/ls"
	"./applets/mkdir"
	"./applets/mknod"
	"./applets/mount"
	"./applets/ps"
	"./applets/rm"
	"./applets/shell"
    "./applets/strings"
	"./applets/telnetd"
    "./applets/touch"
    "./applets/spipe"
	"./applets/umount"
    "./applets/wc"
	"./applets/wget"
)

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet{
	"echo":      echo.Echo,
	"shell":     shell.Shell,
	"spipe":     spipe.Spipe,
	"spiped":    spipe.Spiped,
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
	"wc":        wc.Wc,
	"kill":      kill.Kill,
	"cat":       cat.Cat,
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
