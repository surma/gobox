package grep

import (
	"common"
	flag "appletflag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var (
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() < 1 || *helpFlag {
		println("`grep` <pattern> [<file>...]")
		flag.PrintDefaults()
		return
	}

	pattern, err := regexp.Compile(flag.Arg(0))
	if err != nil {
		log.Fatalf("Invalid regular expression: %s\n", err)
	}

	if flag.NArg() == 1 {
		doGrep(pattern, os.Stdin, "<stdin>", false)
	} else {
		for _, fn := range flag.Args()[1:] {
			if fh, err := os.Open(fn); err == nil {
				func() {
					defer fh.Close()
					doGrep(pattern, fh, fn, flag.NArg() > 2)
				}()
			} else {
				log.Printf("Could not open file %s: %s\n", fn, err)
			}
		}
	}

	return
}

func doGrep(pattern *regexp.Regexp, fh io.Reader, fn string, print_fn bool) {
	buf := common.NewBufferedReader(fh)

	for {
		line, err := buf.ReadWholeLine()
		if err != nil {
			log.Printf("Could not read from %s: %s\n", fn, err)
			return
		}
		if line == "" {
			break
		}

		if pattern.MatchString(line) {
			if print_fn {
				fmt.Printf("%s:", fn)
			}
			fmt.Printf("%s\n", line)
		}
	}
}
