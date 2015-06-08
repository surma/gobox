package tar

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	flagSet = flag.NewFlagSet("tar", flag.PanicOnError)
	tfile   = flagSet.String("f", "", "tar file")
	extract = flagSet.Bool("x", false, "extract")
	create  = flagSet.Bool("c", false, "create")

	tw *tar.Writer
	tr *tar.Reader
)

func walkpath(path string, f os.FileInfo, err error) error {
	header, _ := tar.FileInfoHeader(f, "")
	header.Name = path
	tw.WriteHeader(header)
	ifile, _ := os.Open(path)
	io.Copy(tw, ifile)
	fmt.Printf("%s with %d bytes\n", path, f.Size())
	return nil

}

func Tar(call []string) error {
	flagSet.Parse(call[1:])

	if *tfile == "" {
		fmt.Printf("Usage for %[1]s: %[1]s [-x] [-c] [-f file] [files ...]\n", "tar")
		flagSet.PrintDefaults()
	}

	if *extract {
		ifile, _ := os.Open(*tfile)
		tr := tar.NewReader(ifile)

		// Iterate through the files in the archive.
		for {
			hdr, err := tr.Next()
			if err == io.EOF {
				// end of tar archive
				break

			}
			if err != nil {
				log.Fatalln(err)

			}
			fi := hdr.FileInfo()
			if fi.IsDir() {
				os.MkdirAll(hdr.Name, 0755)
			} else {
				os.MkdirAll(filepath.Dir(hdr.Name), 0755)
				ofile, _ := os.Create(hdr.Name)
				io.Copy(ofile, tr)
			}
			fmt.Println(hdr.Name)
		}
	} else if *create {
		ofile, _ := os.Create(*tfile)
		tw = tar.NewWriter(ofile)
		for _, incpath := range flagSet.Args() {
			filepath.Walk(incpath, walkpath)
		}
		tw.Close()
		ofile.Close()
	}

	return nil
}
