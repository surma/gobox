package ifconfig

import (
	"flag"

	"syscall"
)

const (
	SOCKADDR_DATA = 14
)

type sa_family_t uint16

type in_port uint16
type in_addr struct {
	addr uint32
}
type sockaddr_in struct {
	sin_family sa_family_t
	sin_port   in_port
	sin_addr   in_addr
}

type ifreq_sockaddr_in struct {
	ifr_name [syscall.IFNAMSIZ]byte
	ifr_addr sockaddr_in
}

type ifreq_flags struct {
	ifr_name  [syscall.IFNAMSIZ]byte
	ifr_flags uint16
}

func Ioctl(fd, request, data uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, request, data)
	if errno == 0 {
		return nil
	}
	return errno
}

func byteArrayToUint32(t []byte) uint32 {
	r := uint32(0)
	for i := uint(0); i < 4; i++ {
		r = (r << 8) | uint32(t[15-i])
	}
	return r
}
