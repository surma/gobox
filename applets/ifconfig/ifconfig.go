package ifconfig

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"syscall"
	"unsafe"
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
	ip, e := getAddrFromIface("eth0")
	if e != nil {
		panic(e)
	}
	fmt.Printf("%v\n", ip)
	return nil
}

const (
	SOCKADDR_DATA = 14
	IFNAMSIZ      = 0x10
)

type sa_family_t uint16
type sockaddr struct {
	sin_family sa_family_t
	data       [14]byte
}
type in_port uint16
type in_addr struct {
	addr uint32
}
type sockaddr_in struct {
	sin_family sa_family_t
	sin_port   in_port
	sin_addr   in_addr
}

type ifreq struct {
	ifr_name [IFNAMSIZ]byte
	ifr_addr sockaddr
}

var (
	ErrInvalidIfaceName = errors.New("Invalid interface name")
	ErrNoPortAvailable  = errors.New("Could not find a free port")
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

func getAddrFromIface(name string) (ip net.IP, e error) {
	if len(name) >= IFNAMSIZ {
		return ip, ErrInvalidIfaceName
	}

	var req ifreq
	copy(req.ifr_name[0:], []byte(name))
	ptr := uintptr(unsafe.Pointer(&req))
	socketIoctl(syscall.SIOCGIFADDR, ptr)
	ipaddr := (*sockaddr_in)(unsafe.Pointer(&req.ifr_addr))
	ip = uint32ToByteArray(ipaddr.sin_addr.addr)
	return ip, nil
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

func Ioctl(fd, request, data uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, request, data)
	if errno == 0 {
		return nil
	}
	return errno
}

func uint32ToByteArray(t uint32) []byte {
	r := make([]byte, 4)
	for i := uint(0); i < 4; i++ {
		r[i] = byte((t >> (i * 8)) & 0xFF)
	}
	return r
}
