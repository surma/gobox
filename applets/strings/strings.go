// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Strings is a more capable, UTF-8 aware version of the standard strings utility.
//
// Flags(=default) are:
//
//-ascii(=false)    restrict strings to ASCII
//-min(=6)          minimum length of UTF-8 strings printed, in runes
//-max(=256)        maximum length of UTF-8 strings printed, in runes
//-offset(=false)   show file name and offset of start of each string
//
package strings

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	flagSet = flag.NewFlagSet("strings", flag.PanicOnError)
	min     = flagSet.Int("min", 6, "minimum length of UTF-8 strings printed, in runes")
	max     = flagSet.Int("max", 256, "maximum length of UTF-8 strings printed, in runes")
	ascii   = flagSet.Bool("ascii", false, "restrict strings to ASCII")
	offset  = flagSet.Bool("offset", false, "show file name and offset of start of each string")
)

var stdout *bufio.Writer

func Strings(call []string) error {
	log.SetFlags(0)
	log.SetPrefix("strings: ")
	stdout = bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	flagSet.Parse(call[1:])
	if *max < *min {
		*max = *min

	}

	if flagSet.NArg() == 0 {
		do("<stdin>", os.Stdin)
	} else {
		for _, arg := range flagSet.Args() {
			fd, err := os.Open(arg)
			if err != nil {
				log.Print(err)
				continue
			}
			do(arg, fd)
			stdout.Flush()
			fd.Close()
		}
	}
	return nil
}

func do(name string, file *os.File) {
	in := bufio.NewReader(file)
	str := make([]rune, 0, *max)
	filePos := int64(0)
	print := func() {
		if len(str) >= *min {
			s := string(str)
			if *offset {
				fmt.Printf("%s:#%d:\t%s\n", name, filePos-int64(len(s)), s)
			} else {
				fmt.Println(s)
			}
		}
		str = str[0:0]
	}
	for {
		var (
			r   rune
			wid int
			err error
		)
		// One string per loop.
		for ; ; filePos += int64(wid) {
			r, wid, err = in.ReadRune()
			if err != nil {
				if err != io.EOF {
					log.Print(err)
				}
				return
			}
			if !strconv.IsPrint(r) || *ascii && r >= 0xFF {
				print()
				continue
			}
			// It's printable. Keep it.
			if len(str) >= *max {
				print()
			}
			str = append(str, r)
		}
	}
}
