package shasum

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	flagSet  = flag.NewFlagSet("shasum", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

//Sha1sum calculates the checksum of file with sha1 hashing algorithm
func Sha1sum(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`sha1sum` <files>")
		flagSet.PrintDefaults()
		return nil
	}
	for _, file := range flagSet.Args() {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			return err
		}
		fmt.Printf("%x %v\n", h.Sum(nil), file)
	}
	return nil
}

//Sha256sum calculates the checksum of file with sha256 hashing algorithm
func Sha256sum(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`sha256sum` <files>")
		flagSet.PrintDefaults()
		return nil
	}
	for _, file := range flagSet.Args() {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			return err
		}
		fmt.Printf("%x %v\n", h.Sum(nil), file)
	}
	return nil
}

//Sha512sum calculates the checksum of file with sha512 hashing algorithm
func Sha512sum(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`sha512sum` <files>")
		flagSet.PrintDefaults()
		return nil
	}
	for _, file := range flagSet.Args() {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		h := sha512.New()
		if _, err := io.Copy(h, f); err != nil {
			return err
		}
		fmt.Printf("%x %v\n", h.Sum(nil), file)
	}
	return nil
}
