package grep

import (
	"flag"
	"os"
	"fmt"
	"regexp"
	"io"
	"bufio"
)

var (
	flagSet  = flag.NewFlagSet("grep", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Grep(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() < 1 || *helpFlag {
		println("`grep` <pattern> [<file>...]")
		flagSet.PrintDefaults()
		return nil
	}

	pattern, err := regexp.Compile(flagSet.Arg(0))
	if err != nil {
		return err
	}

	if flagSet.NArg() == 1 {
		doGrep(pattern, os.Stdin, "<stdin>", false)
	} else {
		for _, fn := range flagSet.Args()[1:] {
			if fh, err := os.Open(fn); err == nil {
				doGrep(pattern, fh, fn, flagSet.NArg() > 2)
				fh.Close()
			} else {
				fmt.Fprintf(os.Stderr, "grep: %s: %v\n", fn, err)
			}
		}
	}

	return nil
}

func doGrep(pattern *regexp.Regexp, fh io.Reader, fn string, print_fn bool) {
	buf := bufio.NewReader(fh)

	for {
		line, err := readLine(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while reading from %s: %v\n", fn, err)
			return
		}
		if line == nil {
			break
		}

		if pattern.Match(line) {
			if print_fn {
				fmt.Printf("%s:", fn)
			}
			fmt.Printf("%s\n", line)
		}
	}
}

func readLine(buf *bufio.Reader) ([]byte, os.Error) {
	str := []byte{}
	for {
		data, prfx, err := buf.ReadLine()
		str = append(str, data...)
		if err != nil {
			if err == os.EOF {
				return nil, nil
			} else {
				return nil, err
			}
		}
		if !prfx {
			break
		}
	}
	return str, nil
}
