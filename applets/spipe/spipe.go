package spipe

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"

	"github.com/dchest/spipe"
)

var (
	flagSet  = flag.NewFlagSet("spipe", flag.PanicOnError)
	fAddr    = flagSet.String("t", "", "target socket")
	sAddr    = flagSet.String("s", "", "source socket")
	fKeyFile = flagSet.String("k", "", "key file name")
)

func Spipe(call []string) error {
	flagSet.Parse(call[1:])
	if *fKeyFile == "" {
		fmt.Printf("Usage of %s:\n", "spipe")
		flagSet.PrintDefaults()
		return nil

	}
	// Read key file.
	key, err := ioutil.ReadFile(*fKeyFile)
	if err != nil {
		return fmt.Errorf("key file: %s", err)
	}

	// Dial.
	conn, err := spipe.Dial(key, "tcp", *fAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	errc := make(chan error, 1)

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		errc <- err

	}()
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		errc <- err

	}()
	<-errc
	conn.Close()

	return nil
}

func Spiped(call []string) error {
	flagSet.Parse(call[1:])
	if *fKeyFile == "" {
		fmt.Printf("Usage of %s:\n", "spiped")
		flagSet.PrintDefaults()
		return nil

	}
	// Read key file.
	key, err := ioutil.ReadFile(*fKeyFile)
	if err != nil {
		return fmt.Errorf("key file: %s", err)

	}
	s, _ := net.Listen("tcp", *sAddr)
	for {
		c, _ := s.Accept()
		// Dial.
		conn, err := spipe.Dial(key, "tcp", *fAddr)
		if err != nil {
			return fmt.Errorf("Dial: %s", err)

		}
		defer conn.Close()
		errc := make(chan error, 1)
		go func() {
			_, err := io.Copy(conn, c)
			errc <- err

		}()
		go func() {
			_, err := io.Copy(c, conn)
			errc <- err

		}()
		<-errc
		conn.Close()
	}

	return nil
}
