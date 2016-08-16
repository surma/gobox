package head

import (
	flag "../../appletflag"
	"../../common"
	"fmt"
	"github.com/surma/gobox/pkg/common"
	"io"
	"log"
	"os"
)

var (
	helpFlag = flag.Bool("help", false, "Show this help")
	numLines = flag.Uint("n", 10, "Print -n <number> of lines. Default is 10.")
	quiet    = flag.Bool("q", false, "Don't print file names in multi-file mode.")
)

func Main() {
	flag.Parse()

	argn := flag.NArg()
	if argn <= 0 || *helpFlag {
		println("`head` [options] <files>")
		flag.PrintDefaults()
		return
	}

	for _, file := range flag.Args() {
		if !*quiet {
			fmt.Fprintf(os.Stdout, "==> %s <==\n", file)
		}

		e := dumpFile(file)
		if e != nil {
			log.Printf("Could not read file %s: %s\n", file, e)
		}
	}

	return
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
