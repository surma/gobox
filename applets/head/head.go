package head

import (
	"flag"
	"fmt"
	"github.com/surma/gobox/pkg/common"
	"io"
	"os"
)

var (
	flagSet  = flag.NewFlagSet("head", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
	numLines = flagSet.Uint("n", 10, "Print -n <number> of lines. Default is 10.")
	quiet    = flagSet.Bool("q", false, "Don't print file names in multi-file mode.")
)

func Head(call []string) error {
	err := flagSet.Parse(call[1:])
	if err != nil {
		return err
	}

	argn := flagSet.NArg()
	if argn <= 0 || *helpFlag {
		println("`head` [options] <files>")
		flagSet.PrintDefaults()
		return nil
	}

	for _, file := range flagSet.Args() {
		if !*quiet {
			fmt.Fprintf(os.Stdout, "==> %s <==\n", file)
		}

		err = dumpFile(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func dumpFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := common.NewBufferedReader(f)

	line, err := r.ReadWholeLine()
	for i := uint(0); i < *numLines && err == nil; i++ {
		fmt.Println(line)
		line, err = r.ReadWholeLine()
	}
	if err == io.EOF {
		err = nil
	}
	return err
}
