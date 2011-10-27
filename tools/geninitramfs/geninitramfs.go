package main

import (
	"os"
	"cpio"
)

func main() {
	f, _ := os.OpenFile("test.cpio", os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0644)
	defer f.Close()
	w := cpio.NewWriter(f)
	defer w.Close()
}
