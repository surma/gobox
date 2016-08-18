package ifconfig

import (
	"flag"

	"fmt"
	"log"
	"net"
	"os"
)

var (
	flagSet     = flag.NewFlagSet("ifconfig", flag.PanicOnError)
	addrFlag    = flag.String("addr", "", "Address to set")
	netmaskFlag = flag.String("netmask", "", "Netmask to set")
	stateFlag   = flag.String("state", "", "Set the interface up")
	listFlag    = flag.Bool("list", false, "List one or all interfaces")
	helpFlag    = flag.Bool("help", false, "Show this help")
)

func Help() {
	fmt.Println("`ifconfig` {-list [interface] | [options] [-state {up|down}] <interface>}")
	flag.PrintDefaults()
}

func Ifconfig(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	narg := flag.NArg()
	if (narg != 0 && narg != 1) || *helpFlag {
		Help()
		return nil
	}
	var e error
	if *listFlag {
		if narg == 0 {
			e = ListAllInterfaces()
		} else {
			e = ListInterface(flag.Arg(0))
		}
		if e != nil {
			log.Fatalf("Could not list interface(s): %s\n", e)
		}
	}
	if narg != 1 {
		Help()
		return nil
	}
	if *addrFlag != "" {
		ip := net.ParseIP(*addrFlag)
		if ip == nil {
			log.Fatalf("Invalid IP")
		}
		e := SetAddrForIface(flag.Arg(0), ip)
		if e != nil {
			log.Fatalf("Could not set address: %s\n", e)
		}
	}
	if *netmaskFlag != "" {
		nm := net.ParseIP(*netmaskFlag)
		if nm == nil {
			log.Fatalf("Invalid netmask")
		}
		e := SetNetmaskForIface(flag.Arg(0), nm)
		if e != nil {
			log.Fatalf("Could not set netmask: %s\n", e)
		}
	}
	if *stateFlag == "up" || *stateFlag == "down" {
		up := *stateFlag == "up"
		e := SetStateForIface(flag.Arg(0), up)
		if e != nil {
			log.Fatalf("Could not change state: %s\n", e)
		}
	} else if *stateFlag != "" {
		log.Fatalf("Invalid state")
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
