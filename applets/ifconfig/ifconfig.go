package ifconfig

import (
	"flag"
	"fmt"
	"net"
	"unsafe"
	"syscall"
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

	if (flagSet.NArg() != 0 && flagSet.NArg() != 4) || *helpFlag {
		println("`ifconfig` [interface ip netmask {up|down}]")
		flagSet.PrintDefaults()
		return nil
	}
	list, e := GetIfaceNames()
	if e != nil {
		panic(e)
	}
	for _, iface := range list {
		fmt.Printf("%s\n", iface)
	}
	ip, e := GetAddrFromIface("eth0")
	if e != nil {
		panic(e)
	}
	fmt.Printf("%v\n", ip)
	return nil
}

func GetAddrFromIface(name string) (ip net.IP, e error) {
	if len(name) >= syscall.IFNAMSIZ {
		return ip, ErrInvalidIfaceName
	}

	var req ifreq_sockaddr
	copy(req.ifr_name[0:], []byte(name))
	ptr := uintptr(unsafe.Pointer(&req))
	e = socketIoctl(syscall.SIOCGIFADDR, ptr)
	if e != nil {
		return ip, e
	}
	ipaddr := (*sockaddr_in)(unsafe.Pointer(&req.ifr_addr))
	ip = uint32ToByteArray(ipaddr.sin_addr.addr)
	return ip, nil
}

func GetIfaceNames() ([]string, error) {
	// Apparently due to the union the struct is
	// 8 bytes bigger than it’s Go pendant
	structsize := int(unsafe.Sizeof(ifreq_sockaddr{}))+8
	list := ifconf_list {
		ifc_len: 1,
	}
	for i := 1; ; i++ {
		list.ifc_req = make([]ifreq_sockaddr, i)
		list.ifc_len = i*structsize
		ptr := uintptr(unsafe.Pointer(&list))
		e := socketIoctl(syscall.SIOCGIFCONF, ptr)
		if e != nil {
			return nil, e
		}
		list.ifc_len /= structsize
		// ifc_len is set by the kernel to the number
		// of bytes of interface descriptions put into the struct.
		// When that value is smaller than the size
		// of the array, we got everyting.
		// ...
		// And yes, this is how the `man 7 netdevice` tells
		// me to go about this.
		if i > list.ifc_len {
			list.ifc_req = list.ifc_req[0:i-1]
			break
		}
	}
	ret := make([]string, 0)
	for _, iface := range list.ifc_req {
		ret = append(ret, string(iface.ifr_name[0:]))
	}
	return ret, nil
}
