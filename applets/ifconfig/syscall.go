package ifconfig

import (
	"syscall"
)

const (
	SOCKADDR_DATA = 14
)

type sa_family_t uint16
type sockaddr struct {
	sin_family sa_family_t
	data       [SOCKADDR_DATA]byte
	padding    [8]byte
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

type ifreq_sockaddr struct {
	ifr_name [syscall.IFNAMSIZ]byte
	ifr_addr sockaddr
}

type ifconf_list struct {
	ifc_len int
	ifc_req []ifreq_sockaddr
}

func Ioctl(fd, request, data uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, request, data)
	if errno == 0 {
		return nil
	}
	return errno
}

func uint32ToByteArray(t uint32) []byte {
	r := make([]byte, 16)
	for i := uint(0); i < 4; i++ {
		r[i+12] = byte((t >> (i * 8)) & 0xFF)
	}
	return r
}

func byteArrayToUint32(t []byte) uint32 {
	r := uint32(0)
	for i := uint(0); i < 4; i++ {
		r = (r<<8) | uint32(t[15-i])
	}
	return r
}

