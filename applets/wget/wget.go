package wget

import (
	"flag"
	"os"
	"http"
	"url"
	"path"
	"io"
)

var (
	flagSet  = flag.NewFlagSet("wget", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Wget(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`wget` [options] <url>")
		flagSet.PrintDefaults()
		return nil
	}

	filename, e := getFilenameFromURL(flagSet.Arg(0))
	if e != nil {
		return e
	}

	output, e := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
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

func getFilenameFromURL(rawurl string) (string, os.Error) {
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
