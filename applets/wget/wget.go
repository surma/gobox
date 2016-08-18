package wget

import (
	"flag"

	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

var (
	outFlag  = flag.String("o", "", "Filename to save output to")
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() != 1 || *helpFlag {
		println("`wget` [options] <url>")
		flag.PrintDefaults()
		return
	}

	output, e := getOutputFile(flag.Arg(0))
	if e != nil {
		log.Fatalf("Could not open output file %s: %s", flag.Arg(0), e)
	}
	defer output.Close()

	c := http.Client{}
	r, e := c.Get(flag.Arg(0))
	if e != nil {
		log.Fatalf("Could not issue HTTP request: %s", e)
	}
	defer r.Body.Close()

	_, e = io.Copy(output, r.Body)
}

func getFilenameFromURL(rawurl string) (string, error) {
	url, e := url.Parse(rawurl)
	if e != nil {
		return "", e
	}

	fname := path.Base(url.Path)
	if len(fname) == 0 || fname == "." {
		fname = "index.html"
	}
	return fname, nil
}

func getOutputFile(rawurl string) (io.WriteCloser, error) {
	if *outFlag == "-" {
		return os.Stdout, nil
	}
	var filename string
	var e error
	if len(*outFlag) == 0 {
		filename, e = getFilenameFromURL(flag.Arg(0))
		if e != nil {
			return nil, e
		}
	} else {
		filename = *outFlag
	}
	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

}
