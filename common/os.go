package common

import (
	"os"
)

// Checks if the given path exists
// i.e. if a os.Stat() succeeds
func PathExists(path string) bool {
	_, e := os.Stat(path)
	return e == nil

}

// Prints an error to stdout in a "nice" way.
func DumpError(e os.Error) {
	println("gobox:", "Error:", e.String())
}

// Creates a symlink, and overwrites the target, if it
// happens to exist.
func ForcedSymlink(oldname, newname string) os.Error {
	if PathExists(newname) {
		e := os.Remove(newname)
		if e != nil {
			return e
		}
	}
	return os.Symlink(oldname, newname)
}
