package ifconfig

import (
	"net"
	"syscall"
	"unsafe"
)

func createSocket() (*net.TCPListener, error) {
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:1")
	for i := 65535; i > 1024; i-- {
		addr.Port = i
		listener, e := net.ListenTCP("tcp", addr)
		if e == nil {
			return listener, nil
		}
	}
	return nil, ErrNoPortAvailable
}

func getFdFromSocket(sock *net.TCPListener) (int, error) {
	f, e := sock.File()
	if e != nil {
		return 0, e
	}
	return f.Fd(), nil
}

func socketIoctl(request, data uintptr) error {
	sock, e := createSocket()
	if e != nil {
		return e
	}
	defer sock.Close()

	fd, e := getFdFromSocket(sock)
	if e != nil {
		return e
	}

	return Ioctl(uintptr(fd), request, data)
}

func uint32ToByteArray(t uint32) []byte {
	r := make([]byte, 4)
	for i := uint(0); i < 4; i++ {
		r[i] = byte((t >> (i * 8)) & 0xFF)
	}
	return r
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

func GetNetmaskFromIface(name string) (ip net.IP, e error) {
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

func cstrToString(str []byte) string {
	for i := range str {
		if str[i] == 0 {
			return string(str[0:i])
		}
	}
	return ""
}
