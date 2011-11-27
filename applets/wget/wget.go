package wget

import (
	"flag"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

var (
	flagSet  = flag.NewFlagSet("wget", flag.PanicOnError)
	outFlag  = flagSet.String("o", "", "Filename to save output to")
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Wget(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 1 || *helpFlag {
		println("`wget` [options] <url>")
		flagSet.PrintDefaults()
		return nil
	}

	output, e := getOutputFile(flagSet.Arg(0))
	if e != nil {
		return e
	}
	defer output.Close()

	c := http.Client{}
	r, e := c.Get(flagSet.Arg(0))
	if e != nil {
		return e
	}
	defer r.Body.Close()

	_, e = io.Copy(output, r.Body)

	return e
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
		filename, e = getFilenameFromURL(flagSet.Arg(0))
		if e != nil {
			return nil, e
		}
	} else {
		filename = *outFlag
	}
	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

}
