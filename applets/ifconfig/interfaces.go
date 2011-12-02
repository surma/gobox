package ifconfig

import (
	"syscall"
	"net"
	"unsafe"
)

type Interface struct {
	Name    string
	Address net.IP
	Netmask net.IP
}

func (this *Interface) Load() (e error) {
	this.Address, e = getAddrFromIface(this.Name)
	if e != nil {
		return e
	}
	this.Netmask, e = getNetmaskFromIface(this.Name)
	if e != nil {
		return e
	}
	return nil
}

func (this *Interface) Set() (e error) {
	e = setAddrForIface(this.Name, this.Address)
	if e != nil {
		return e
	}
	e = setNetmaskForIface(this.Name, this.Netmask)
	if e != nil {
		return e
	}
	return nil
}

func setAddrForIface(name string, ip net.IP) error {
	if len(name) >= syscall.IFNAMSIZ {
		return ErrInvalidIfaceName
	}

	var req ifreq_sockaddr
	copy(req.ifr_name[0:], []byte(name))
	ipaddr := (*sockaddr_in)(unsafe.Pointer(&req.ifr_addr))
	ipaddr.sin_family = syscall.AF_INET
	ipaddr.sin_addr.addr = byteArrayToUint32(ip)
	ptr := uintptr(unsafe.Pointer(&req))
	return socketIoctl(syscall.SIOCSIFADDR, ptr)
}

func getAddrFromIface(name string) (ip net.IP, e error) {
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

func setNetmaskForIface(name string, ip net.IP) error {
	if len(name) >= syscall.IFNAMSIZ {
		return ErrInvalidIfaceName
	}

	var req ifreq_sockaddr
	copy(req.ifr_name[0:], []byte(name))
	ipaddr := (*sockaddr_in)(unsafe.Pointer(&req.ifr_addr))
	ipaddr.sin_family = syscall.AF_INET
	ipaddr.sin_addr.addr = byteArrayToUint32(ip)
	ptr := uintptr(unsafe.Pointer(&req))
	return socketIoctl(syscall.SIOCSIFNETMASK, ptr)
}

func getNetmaskFromIface(name string) (ip net.IP, e error) {
	if len(name) >= syscall.IFNAMSIZ {
		return ip, ErrInvalidIfaceName
	}

	var req ifreq_sockaddr
	copy(req.ifr_name[0:], []byte(name))
	ptr := uintptr(unsafe.Pointer(&req))
	e = socketIoctl(syscall.SIOCGIFNETMASK, ptr)
	if e != nil {
		return ip, e
	}
	ipaddr := (*sockaddr_in)(unsafe.Pointer(&req.ifr_addr))
	ip = uint32ToByteArray(ipaddr.sin_addr.addr)
	return ip, nil
}

func GetInterfaceNames() ([]string, error) {
	structsize := int(unsafe.Sizeof(ifreq_sockaddr{}))
	list := ifconf_list{
		ifc_len: 1,
	}
	for i := 1; ; i++ {
		list.ifc_req = make([]ifreq_sockaddr, i)
		list.ifc_len = i * structsize
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
			list.ifc_req = list.ifc_req[0 : i-1]
			break
		}
	}
	ret := make([]string, 0)
	for _, iface := range list.ifc_req {
		ret = append(ret, cstrToString(iface.ifr_name[0:]))
	}
	return ret, nil
}
