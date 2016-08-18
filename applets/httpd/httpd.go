package httpd

import (
	"flag"

	"log"
	"net/http"
)

var (
	flagSet  = flag.NewFlagSet("httpd", flag.PanicOnError)
	addrFlag = flag.String("addr", ":8080", "Address to listen on")
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Httpd(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flag.NArg() != 1 || *helpFlag {
		println("`httpd` [options] <dir>")
		flag.PrintDefaults()
		return nil
	}

	e = http.ListenAndServe(*addrFlag, http.FileServer(http.Dir(flag.Arg(0))))
	if e != nil {
		log.Fatalf("Could not start server: %s\n", e)
	}

	return nil
}
