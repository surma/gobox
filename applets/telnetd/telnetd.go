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
	portFlag = flagSet.Int("port", 23, "Port to listen on")
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Telnetd(call []string) os.Error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if *helpFlag || flagSet.NArg() <= 0 {
		println("telnet [options] <command to serve...>")
		flagSet.PrintDefaults()
		return nil
	}

	return startServer(*portFlag, flagSet.Args())
}

func startServer(port int, call []string) os.Error {
	l, e := net.ListenTCP("tcp4", &net.TCPAddr{net.IPv4zero, port})
	if e != nil {
		return e
	}
	defer l.Close()

	for {
		c, e := l.Accept()
		go func() {
			defer func() {
				if p := recover(); p != nil {
					e, ok := p.(os.Error)
					if !ok {
						e = os.NewError("Some error occured")
					}
					common.FDumpError(e, c)
				}
			}()
			defer c.Close()
			if e != nil {
				return
			}
			serveExec(c, call)
			return
		}()
	}
	return nil
}

func serveExec(inout io.ReadWriter, call []string) {
	if len(call) < 1 {
		panic(os.NewError("Trying to serve an empty command"))
	}
	c := exec.Command(call[0], call[1:]...)
	c.Stdin = inout
	c.Stdout = inout
	c.Stderr = inout
	e := c.Run()
	if e != nil {
		panic(e)
	}
}
