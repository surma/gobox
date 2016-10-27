package wc

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jasonmoo/wc"
)

var (
	flagSet    = flag.NewFlagSet("wc", flag.PanicOnError)
	multibytes = flagSet.Bool("m", false, "count the multibyte runes")
	lines      = flagSet.Bool("l", false, "count the lines")
	words      = flagSet.Bool("w", false, "count the words")
	bytes      = flagSet.Bool("c", false, "count the bytes")
)

func Wc(call []string) error {

	flagSet.Parse(call[1:])

	if flagSet.NArg() == 0 {
		fmt.Println("wc [-l] [-m] [-w] [-b] file [...fileN]")
		flagSet.PrintDefaults()
		return nil
	}

	// Default flags if non specified
	if flagSet.NFlag() == 0 {
		*lines = true
		*words = true
		*bytes = true
	}

	var multibytes_total, lines_total, words_total, bytes_total uint64

	for _, filepath := range flagSet.Args() {
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}

		c := wc.NewCounter(file)

		err = c.Count(*multibytes, *bytes, *lines, *words)
		if err != nil {
			log.Fatal(err)
		}

		file.Close()

		if *lines {
			lines_total += c.Lines
			fmt.Printf("% 10d ", c.Lines)
		}
		if *words {
			words_total += c.Words
			fmt.Printf("% 10d ", c.Words)
		}
		if *multibytes {
			multibytes_total += c.Multibytes
			fmt.Printf("% 10d ", c.Multibytes)
		}
		if *bytes {
			bytes_total += c.Bytes
			fmt.Printf("% 10d ", c.Bytes)
		}

		fmt.Printf("%s\n", filepath)
	}

	if flagSet.NArg() > 1 {
		if *lines {
			fmt.Printf("% 10d ", lines_total)
		}
		if *words {
			fmt.Printf("% 10d ", words_total)
		}
		if *multibytes {
			fmt.Printf("% 10d ", multibytes_total)
		}
		if *bytes {
			fmt.Printf("% 10d ", bytes_total)
		}

		fmt.Print("total\n")
	}

	return nil
}
