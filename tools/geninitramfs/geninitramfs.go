package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"github.com/surma/gocpio"
	"io"
	"log"
	"os"
)

var (
	outputFlag = flag.String("o", "initramfs", "Define output filename")
	helpFlag   = flag.Bool("h", false, "Show this help")
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 || *helpFlag {
		flag.PrintDefaults()
		return
	}

	out, e := os.OpenFile(*outputFlag, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if e != nil {
		log.Fatalf("Error while opening output file: %s", e.Error())
	}
	defer out.Close()
	comp_out, e := gzip.NewWriterLevel(out, gzip.BestCompression)
	if e != nil {
		log.Fatalf("Error while setting up the compressor: %s", e.Error())
	}
	defer comp_out.Close()
	cpio_out := cpio.NewWriter(comp_out)
	defer cpio_out.Close()

	in, e := os.Open(flag.Arg(0))
	if e != nil {
		log.Fatalf("Error while opening input file: %s", e.Error())
	}
	defer in.Close()

	c := make(chan *Entry)
	go parseInput(in, c)
	createCpioArchive(cpio_out, c)
}

type Entry struct {
	hdr  cpio.Header
	data io.Reader
}

func createCpioArchive(w *cpio.Writer, c <-chan *Entry) {
	for entry := range c {
		fmt.Printf("Entry: %s\n", entry.hdr.Name)
		w.WriteHeader(&entry.hdr)
		if entry.data != nil {
			io.Copy(w, entry.data)
			if closer, ok := entry.data.(io.Closer); ok {
				closer.Close()
			}
		}
	}
}
