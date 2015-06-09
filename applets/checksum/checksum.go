package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"os"
)

var (
	flagSet = flag.NewFlagSet("checksum", flag.PanicOnError)
)

func Hash(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() < 1 {
		fmt.Printf("Usage of %s\n", call[0])
		fmt.Println("`%s` [file ...]", call[0])
		return nil
	}

	var h hash.Hash
	switch call[0] {
	case "md5sum":
		h = md5.New()
	case "sha1sum":
		h = sha1.New()
	case "sha256sum":
		h = sha256.New()
	case "sha512sum":
		h = sha512.New()
	case "crc32":
		h = crc32.NewIEEE()
	}

	for _, v := range flagSet.Args() {
		infile, err := os.Open(v)
		if err != nil {
			fmt.Fprint(os.Stderr, "md5sum: %s\n", err)
		}
		io.Copy(h, infile)
		fmt.Printf("%x %s\n", h.Sum(nil), v)
		h.Reset()
	}

	return nil
}
