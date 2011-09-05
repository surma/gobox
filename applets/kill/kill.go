package kill

import (
	"flag"
	"os"
	"strconv"
)

var (
	flagSet  = flag.NewFlagSet("kill", flag.PanicOnError)
	signalFlag = flagSet.Int("sig", 9, "Number of the signal to send")
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Kill(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("`kill` [options] <pid>")
		flagSet.PrintDefaults()
		return nil
	}

	pid, e := strconv.Atoi(flagSet.Arg(0))
	if e != nil {
		return e
	}

	p, e := os.FindProcess(pid)
	if e != nil {
		return e
	}

	return p.Signal(os.UnixSignal(int32(*signalFlag)))
}
