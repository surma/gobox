package ifconfig

import (
	"flag"
	"fmt"
	"net"
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
		iface := Interface {
			Name: flagSet.Arg(0),
			Address: ip,
			Netmask: nm,
		}
		return iface.Set()
	}
	return nil
}

func ListAllInterfaces() error {
	list, e := GetInterfaceNames()
	if e != nil {
		return e
	}
	var iface Interface
	for _, name := range list {
		iface.Name = name
		e = iface.Load()
		if e != nil {
			fmt.Printf("Could not obtain data of %s: %s\n", name, e)
			continue
		}
		fmt.Printf("%s: %s/%s\n", iface.Name, iface.Address, iface.Netmask)
	}
	return nil
}

