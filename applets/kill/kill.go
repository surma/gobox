package kill

import (
	"flag"
	"os"
	"strconv"
	"syscall"
)

var (
	flagSet    = flag.NewFlagSet("kill", flag.PanicOnError)
	signalFlag = flagSet.Int("sig", 9, "Number of the signal to send")
	helpFlag   = flagSet.Bool("help", false, "Show this help")
)

func Kill(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 1 || *helpFlag {
		println("`kill` [options] <pid>")
		flagSet.PrintDefaults()
		println("1     SIGHUP       terminal line hangup")
		println("2     SIGINT       interrupt program")
		println("3     SIGQUIT      quit program")
		println("4     SIGILL       illegal instruction")
		println("5     SIGTRAP      trace trap")
		println("6     SIGABRT      abort program (formerly SIGIOT)")
		println("7     SIGEMT       emulate instruction executed")
		println("8     SIGFPE       floating-point exception")
		println("9     SIGKILL      kill program")
		println("10    SIGBUS       bus error")
		println("11    SIGSEGV      segmentation violation")
		println("12    SIGSYS       non-existent system call invoked")
		println("13    SIGPIPE      write on a pipe with no reader")
		println("14    SIGALRM      real-time timer expired")
		println("15    SIGTERM      software termination signal")
		println("16    SIGURG       urgent condition present on socket")
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

	return p.Signal(syscall.Signal(int32(*signalFlag)))
}
