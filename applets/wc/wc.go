package wc

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/surma/gobox/pkg/common"
)

var (
	flagSet  = flag.NewFlagSet("wc", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

//Wc counts lines and bytes for given files.
func Wc(call []string) error {
	err := flagSet.Parse(call[1:])
	if err != nil {
		return err
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`wc` <files>")
		flagSet.PrintDefaults()
		return nil
	}
	totalLineCount, totalWordCount, totalByteCount := 0, 0, 0
	for _, file := range flagSet.Args() {

		lineCount, wordCount, byteCount, err := lineCounter(file)
		totalLineCount += lineCount
		totalWordCount += wordCount
		totalByteCount += byteCount
		fmt.Printf("%8v%8v%8v %s\n", lineCount, wordCount, byteCount, file)
		if err != nil {
			return err
		}
	}
	if flagSet.NArg() > 1 {
		fmt.Printf("%8v%8v%8v %s\n", totalLineCount, totalWordCount, totalByteCount, "total")
	}
	return nil
}

func lineCounter(path string) (int, int, int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, 0, 0, err
	}
	defer f.Close()
	buf := make([]byte, 32*1024)
	lineCount := 0
	byteCount := 0
	wordCount := 0
	lineSep := []byte{'\n'}

	r := common.NewBufferedReader(f)

	for {
		c, err := r.Read(buf)
		lineCount += bytes.Count(buf[:c], lineSep)
		wordCount += len(bytes.Fields(buf[:c]))
		byteCount += c
		switch {
		case err == io.EOF:
			return lineCount, wordCount, byteCount, nil

		case err != nil:
			return lineCount, wordCount, byteCount, err
		}
	}

}
