package ifconfig

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	flagSet  = flag.NewFlagSet("ifconfig", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Ifconfig(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if (flagSet.NArg() != 0 && flagSet.NArg() != 3) || *helpFlag {
		println("`ifconfig` [interface ip netmask]")
		flagSet.PrintDefaults()
		return nil
	}
	if flagSet.NArg() == 0 {
		return ListAllInterfaces()
	} else {
		ip := net.ParseIP(flagSet.Arg(1))
		nm := net.ParseIP(flagSet.Arg(2))
		if ip == nil || nm == nil {
			return ErrInvalidAddressFormat
		}
		return SetInterface(flagSet.Arg(0), ip, nm)
	}
	return nil
}

func ListAllInterfaces() error {
	list, e := net.Interfaces()
	if e != nil {
		return e
	}
	for _, iface := range list {
		fmt.Printf("%s (%v)\n", iface.Name, iface.HardwareAddr)
		addrs, e := iface.Addrs()
		if e != nil {
			fmt.Fprintf(os.Stderr, "Error while getting addresses: %s\n", e)
			continue
		}
		for _, addr := range addrs {
			fmt.Printf("\t%s\n", addr)
		}
		fmt.Println()
	}
	return nil
}

