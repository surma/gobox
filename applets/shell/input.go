package shell

import "bufio"

func getNextLine(in *bufio.Reader) (string, error) {
	var e error
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
