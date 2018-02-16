package shasum

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

var (
	flagSet  = flag.NewFlagSet("shasum", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

//Sha1sum calculates the checksum of file with sha1 hashing algorithm
func Sha1sum(call []string) error {
	if e := showHelp(call[1:], "sha1sum"); e != nil {
		return e

	}
	h := sha1.New()
	return hashFiles(h)

}

//Sha256sum calculates the checksum of file with sha256 hashing algorithm
func Sha256sum(call []string) error {
	if e := showHelp(call[1:], "sha256sum"); e != nil {
		return e

	}
	h := sha256.New()
	return hashFiles(h)

}

//Sha512sum calculates the checksum of file with sha512 hashing algorithm
func Sha512sum(call []string) error {
	if e := showHelp(call[1:], "sha512sum"); e != nil {
		return e
	}
	h := sha512.New()
	return hashFiles(h)

}

func hashFiles(h hash.Hash) error {
	for _, file := range flagSet.Args() {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := io.Copy(h, f); err != nil {
			return err
		}
		fmt.Printf("%x %v\n", h.Sum(nil), file)
	}
	return nil
}

func showHelp(call []string, commandName string) error {
	e := flagSet.Parse(call)
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println(commandName, " <files>")
		flagSet.PrintDefaults()
		return nil
	}

	return nil
}
