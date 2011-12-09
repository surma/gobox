package kill

import (
	flag "appletflag"
	"log"
	"os"
	"strconv"
)

var (
	signalFlag = flag.Int("sig", 9, "Number of the signal to send")
	helpFlag   = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() != 1 || *helpFlag {
		println("`kill` [options] <pid>")
		flag.PrintDefaults()
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
		return
	}

	pid, e := strconv.Atoi(flag.Arg(0))
	if e != nil {
		log.Fatalf("Invald PID %s: %s\n", flag.Arg(0), e)
	}

	p, e := os.FindProcess(pid)
	if e != nil {
		log.Fatalf("Could not find process: %s\n", e)
	}

	e = p.Signal(os.UnixSignal(int32(*signalFlag)))
	if e != nil {
		log.Fatalf("Could not send signal: %s\n", e)
	}
}
