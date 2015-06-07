package date

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	flagSet = flag.NewFlagSet("date", flag.PanicOnError)
	uflag   = flagSet.Bool("u", false, "Report Coordinated Universal Time (UTC) rather than local time.")
	nflag   = flagSet.Bool("n", false, "Report the date as the number of seconds since the epoch, 00:00:00 UTC, January 1, 1970")
)

func Date(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	var t time.Time
	switch flagSet.NArg() {
	case 0:
		t = time.Now()
	case 1:
		i, err := strconv.ParseInt(flag.Arg(0), 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "date: error parsing time:", err)
			os.Exit(1)

		}
		t = time.Unix(i, 0)
	default:
		fmt.Printf("Usage of %s\n", "date")
		flagSet.PrintDefaults()
		return nil
	}

	switch {
	case *nflag:
		fmt.Println(t.Unix())
	case *uflag:
		fmt.Println(t.UTC())
	default:
		fmt.Println(t)
	}

	return nil
}
