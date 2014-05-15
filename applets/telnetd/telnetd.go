package telnetd

import (
	"errors"
	"flag"
	"github.com/surma/gobox/pkg/common"
	"io"
	"net"
	"os/exec"
)

var (
	flagSet  = flag.NewFlagSet("telnetd", flag.PanicOnError)
	addrFlag = flagSet.String("addr", ":23", "Port to listen on")
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Telnetd(call []string) error {
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

func startServer(addr string, call []string) error {
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

func serveExec(inout io.ReadWriter, call []string) error {
	if len(call) < 1 {
		return errors.New("Trying to serve an empty command")
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
