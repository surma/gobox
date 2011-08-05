package shell

import (
	"bufio"
	"common"
	"os"
)

func Shell(call []string) os.Error {
	var in *bufio.Reader
	if len(call) > 2 {
		call = call[0:1]
	}
	if len(call) == 2 {
		f, e := os.Open(call[1])
		if e != nil {
			return e
		}
		defer f.Close()
		in = bufio.NewReader(f)
	} else {
		in = bufio.NewReader(os.Stdin)
	}

	var e os.Error
	var line string
	for e == nil {
		line, e = getNextLine(in)
		if e != nil {
			return e
		}
		ce := executeLine(line)
		if ce != nil {
			common.DumpError(ce)
		}
	}
	return nil
}

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

func executeLine(line string) os.Error {
	println("Command:", line)
	return nil
}
