package common

import (
	"bufio"
	"os"
	"io"
)

type BufferedReader struct {
	*bufio.Reader
}

func NewBufferedReader(r io.Reader) *BufferedReader {
	return &BufferedReader{
		bufio.NewReader(r),
	}
}

func (b *BufferedReader) ReadWholeLine() (line string, err os.Error) {
	byteline := make([]byte, 0)
	prefix := true
	for prefix {
		var partial []byte
		partial, prefix, err = b.Reader.ReadLine()
		if err != nil {
			break
		}
		byteline = append(byteline, partial...)
	}
	return string(byteline), err
}
