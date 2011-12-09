package telnetd

import (
	"common"
	"errors"
	flag "appletflag"
	"io"
	"log"
	"net"
	"os/exec"
)

var (
	addrFlag = flag.String("addr", ":23", "Port to listen on")
	helpFlag = flag.Bool("help", false, "Show this help")
)

func Main() {
	flag.Parse()

	if flag.NArg() <= 0 || *helpFlag {
		println("telnet [options] <command to serve...>")
		flag.PrintDefaults()
		return
	}

	e := startServer(*addrFlag, flag.Args())
	if e != nil {
		log.Fatalf("Could not start server: %s\n", e)
	}
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
