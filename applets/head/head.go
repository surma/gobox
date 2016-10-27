package head

import (
	"flag"

	"fmt"
	"gobox/common"
	"io"
	"log"
	"os"
)

var (
	flagSet  = flag.NewFlagSet("head", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
	numLines = flagSet.Uint("n", 10, "Print -n <number> of lines. Default is 10.")
	quiet    = flagSet.Bool("q", false, "Don't print file names in multi-file mode.")
)

func Head(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
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

		e := dumpFile(file)
		if e != nil {
			log.Printf("Could not read file %s: %s\n", file, e)
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
