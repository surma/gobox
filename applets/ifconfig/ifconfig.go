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
		fmt.Printf("Not implemented\n")
	}
	return nil
}

func ListAllInterfaces() error {
	list, e := GetInterfaceNames()
	if e != nil {
		return e
	}
	for _, name := range list {
		iface, e := GetInterface(name)
		if e != nil {
			fmt.Printf("Could not obtain data of %s: %s\n", name, e)
			continue
		}
		fmt.Printf("%s: %s/%s\n", iface.Name, iface.Address, iface.Netmask)
	}
	return nil
}

type Interface struct {
	Name    string
	Address net.IP
	Netmask net.IP
}

func GetInterface(name string) (*Interface, error) {
	ip, e := GetAddrFromIface(name)
	if e != nil {
		return nil, e
	}
	nm, e := GetNetmaskFromIface(name)
	if e != nil {
		return nil, e
	}

	return &Interface{
		Name:    name,
		Address: ip,
		Netmask: nm,
	}, nil
}
