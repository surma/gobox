package head

import (
	"flag"
	"os"
	"fmt"
	"bufio"
)

var (
	flagSet  = flag.NewFlagSet("head", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
	lines    = flagSet.Uint("n", 10, "Print -n <number> of lines. Default is 10.")
	quiet    = flagSet.Bool("q", false, "Don't print file names in multi-file mode.")
)

func Head(call []string) os.Error {
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

	printNames := *quiet == false && argn > 1
	for _, file := range flagSet.Args() {
		if printNames {
			fmt.Fprintf(os.Stdout, "==> %s <==\n", file)
		}

		err = dumpFile(file)
		if err != nil {
			return err
		}
	}

	return nil
}

func dumpFile(path string) os.Error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var lineAt uint = 0
	r := bufio.NewReader(f)

	for bytes, partialLine, err := r.ReadLine(); err == nil; bytes, partialLine, err = r.ReadLine() {

		os.Stdout.Write(bytes)

		if partialLine == false {
			lineAt++
			os.Stdout.WriteString("\n")
		}

		if *lines == lineAt || err == os.EOF {
			return nil
		}
	}
	return err
}
