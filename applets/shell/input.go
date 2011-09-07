package shell

import (
	"bufio"
	"os"
)

func getNextLine(in *bufio.Reader) (string, os.Error) {
	var e os.Error
	var buffer []byte
	isPrefix := true
	for isPrefix {
		var subbuffer []byte
		subbuffer, isPrefix, e = in.ReadLine()
		if e != nil {
			break
		}
		buffer = append(buffer, subbuffer...)
	}
	return string(buffer), e
}
