package ifconfig

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	flagSet     = flag.NewFlagSet("ifconfig", flag.PanicOnError)
	addrFlag    = flagSet.String("addr", "", "Address to set")
	netmaskFlag = flagSet.String("netmask", "", "Netmask to set")
	stateFlag   = flagSet.String("state", "", "Set the interface up")
	listFlag    = flagSet.Bool("list", false, "List one or all interfaces")
	helpFlag    = flagSet.Bool("help", false, "Show this help")
)

func Help() {
	fmt.Println("`ifconfig` {-list [interface] | [options] [-state {up|down}] <interface>}")
	flagSet.PrintDefaults()
}

func Ifconfig(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	narg := flagSet.NArg()
	if (narg != 0 && narg != 1) || *helpFlag {
		Help()
		return nil
	}
	if *listFlag {
		if narg == 0 {
			return ListAllInterfaces()
		} else {
			return ListInterface(flagSet.Arg(0))
		}
	}
	if narg != 1 {
		Help()
		return nil
	}
	if *addrFlag != "" {
		ip := net.ParseIP(*addrFlag)
		if ip == nil {
			return ErrInvalidAddressFormat
		}
		e := SetAddrForIface(flagSet.Arg(0), ip)
		if e != nil {
			return e
		}
	}
	if *netmaskFlag != "" {
		nm := net.ParseIP(*netmaskFlag)
		if nm == nil {
			return ErrInvalidAddressFormat
		}
		e := SetNetmaskForIface(flagSet.Arg(0), nm)
		if e != nil {
			return e
		}
	}
	if *stateFlag == "up" || *stateFlag == "down" {
		up := *stateFlag == "up"
		e := SetStateForIface(flagSet.Arg(0), up)
		if e != nil {
			return e
		}
	} else if *stateFlag != "" {
		return ErrInvalidState
	}
	return nil
}

func ListInterface(name string) error {
	iface, _ := net.InterfaceByName(name)
	fmt.Printf("%s (%v)\n", iface.Name, iface.HardwareAddr)
	addrs, e := iface.Addrs()
	if e != nil {
	}
	for _, addr := range addrs {
		fmt.Printf("\t%s\n", addr)
	}
	fmt.Println()
	return nil
}

func ListAllInterfaces() error {
	list, e := net.Interfaces()
	if e != nil {
		return e
	}
	for _, iface := range list {
		e := ListInterface(iface.Name)
		if e != nil {
			fmt.Fprintf(os.Stderr, "Error while getting addresses: %s\n", e)
			continue
		}
	}
	return nil
}
