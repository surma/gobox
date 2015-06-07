// touch - set modification date of a file
package touch

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	flagSet = flag.NewFlagSet("touch", flag.PanicOnError)
	aflag   = flagSet.Bool("a", false, "Change access time instead of modification time")
	tflag   = flagSet.Int("t", 0, "Set to time provided (in seconds)")
)

func Touch(call []string) error {
	flagSet.Parse(call[1:])
	if flagSet.NArg() < 1 {
		fmt.Printf("Usage for %[1]s: %[1]s [ -a  ] [ -t time  ] files...\n", "touch")
		flagSet.PrintDefaults()
		return nil
	}

	var atime, mtime time.Time
	if *tflag == 0 {
		if *aflag {
			atime = time.Now()
		} else {
			mtime = time.Now()
		}
	} else {
		if *aflag {
			atime = time.Unix(int64(*tflag), 0)
		} else {
			mtime = time.Unix(int64(*tflag), 0)
		}
	}

	for _, name := range flagSet.Args() {
		if err := os.Chtimes(name, atime, mtime); err != nil {
			fmt.Fprintln(os.Stderr, "touch:", err)
		}
	}

	return nil
}
