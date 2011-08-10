package common

import (
	"fmt"
	"io"
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
	FDumpError(os.Stdout, e)
}

// Prints an error to a writer in a "nice" way.
func FDumpError(w io.Writer, e os.Error) {
	fmt.Fprintf(w, "gobox: Error: %s\n", e.String())
}

// Creates a symlink and deletes the file blocking
// the name of the symlink.
func ForcedSymlink(oldname, newname string) os.Error {
	if PathExists(newname) {
		e := os.Remove(newname)
		if e != nil {
			return e
		}
	}
	return os.Symlink(oldname, newname)
}
