package ifconfig

import (
	"flag"

	"net"
	"syscall"
	"unsafe"
)

func SetAddrForIface(name string, ip net.IP) error {
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

func SetNetmaskForIface(name string, ip net.IP) error {
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

func SetStateForIface(name string, state bool) error {
	if len(name) >= syscall.IFNAMSIZ {
		return ErrInvalidIfaceName
	}

	var req ifreq_flags
	copy(req.ifr_name[0:], []byte(name))
	ptr := uintptr(unsafe.Pointer(&req))
	e := socketIoctl(syscall.SIOCGIFFLAGS, ptr)
	if e != nil {
		return e
	}
	if state {
		req.ifr_flags |= uint16(syscall.IFF_UP)
	} else {
		req.ifr_flags &= ^uint16(syscall.IFF_UP)
	}
	return socketIoctl(syscall.SIOCSIFFLAGS, ptr)
}
