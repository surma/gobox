package telnetd

import (
	"flag"
	"net"
	"io"
	"os"
	"exec"
	"common"
)

var (
	flagSet  = flag.NewFlagSet("telnetd", flag.PanicOnError)
	addrFlag = flagSet.String("addr", ":23", "Port to listen on")
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Telnetd(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() <= 0 || *helpFlag {
		println("telnet [options] <command to serve...>")
		flagSet.PrintDefaults()
		return nil
	}

	return startServer(*addrFlag, flagSet.Args())
}

func startServer(addr string, call []string) os.Error {
	ta, e := net.ResolveTCPAddr("tcp4", addr)
	if e != nil {
		return e
	}
	l, e := net.ListenTCP("tcp4", ta)
	if e != nil {
		return e
	}
	defer l.Close()

	for {
		c, e := l.Accept()
		if e == nil {
		} else {
			common.DumpError(e)
		}
		go serve(c, call)
	}
	return nil
}

func serve(c io.ReadWriteCloser, call []string) {
	defer c.Close()
	e := serveExec(c, call)
	if e != nil {
		common.FDumpError(c, e)
	}
}

func serveExec(inout io.ReadWriter, call []string) os.Error {
	if len(call) < 1 {
		return os.NewError("Trying to serve an empty command")
	}
	c := exec.Command(call[0], call[1:]...)
	c.Stdin = inout
	c.Stdout = inout
	c.Stderr = inout
	e := c.Run()
	if e != nil {
		return e
	}
	return nil
}
