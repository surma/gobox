package httpd

import (
	flag "../../appletflag"
	"log"
	"net/http"
)

var (
	addrFlag = flag.String("addr", ":8080", "Address to listen on")
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() != 1 || *helpFlag {
		println("`httpd` [options] <dir>")
		flag.PrintDefaults()
		return
	}

	e := http.ListenAndServe(*addrFlag, http.FileServer(http.Dir(flag.Arg(0))))
	if e != nil {
		log.Fatalf("Could not start server: %s\n", e)
	}
}
