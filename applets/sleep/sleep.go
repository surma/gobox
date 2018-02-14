package sleep

import (
	"flag"
	"strconv"
	"time"
)

var (
	flagSet  = flag.NewFlagSet("sleep", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Sleep(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 1 || *helpFlag {
		println("`sleep` <seconds>")
		flagSet.PrintDefaults()
		return nil
	}
	seconds, err := strconv.Atoi(flagSet.Arg(0))
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	return nil
}
