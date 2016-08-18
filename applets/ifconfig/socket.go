package ifconfig

import (
	"flag"

	"net"
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

func getFdFromSocket(sock *net.TCPListener) (uintptr, error) {
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

func cstrToString(str []byte) string {
	for i := range str {
		if str[i] == 0 {
			return string(str[0:i])
		}
	}
	return ""
}
