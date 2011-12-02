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


func SetInterface(name string, addr, netmask net.IP) (e error) {
	e = setAddrForIface(name, addr)
	if e != nil {
		return e
	}
	e = setNetmaskForIface(name, netmask)
	if e != nil {
		return e
	}
	return nil
}

func setAddrForIface(name string, ip net.IP) error {
	if len(name) >= syscall.IFNAMSIZ {
		return ErrInvalidIfaceName
	}

	var req ifreq_sockaddr_in
	copy(req.ifr_name[0:], []byte(name))
	req.ifr_addr.sin_family = syscall.AF_INET
	req.ifr_addr.sin_addr.addr = byteArrayToUint32(ip)
	ptr := uintptr(unsafe.Pointer(&req))
	return socketIoctl(syscall.SIOCSIFADDR, ptr)
}

func setNetmaskForIface(name string, ip net.IP) error {
	if len(name) >= syscall.IFNAMSIZ {
		return ErrInvalidIfaceName
	}

	var req ifreq_sockaddr_in
	copy(req.ifr_name[0:], []byte(name))
	req.ifr_addr.sin_family = syscall.AF_INET
	req.ifr_addr.sin_addr.addr = byteArrayToUint32(ip)
	ptr := uintptr(unsafe.Pointer(&req))
	return socketIoctl(syscall.SIOCSIFNETMASK, ptr)
}
