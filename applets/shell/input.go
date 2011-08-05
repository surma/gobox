package shell

import (
	"bufio"
	"regexp"
	"os"
	"strings"
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

const (
	unquotedParam = "([^ \t\"]+)"
	quotedParam   = "\"((\\\\.|[^\\\"])*)\""
	paramRegexp   = "^[ \t]*(" + unquotedParam + "|" + quotedParam + ")"
)

var (
	paramMatcher *regexp.Regexp = regexp.MustCompile(paramRegexp)
)

func parameterize(line string) ([]string, os.Error) {
	result := make([]string, 0)
	for len(line) != 0 {
		indices := paramMatcher.FindStringIndex(line)
		if indices == nil {
			println(paramRegexp)
			return nil, os.NewError("Ill-formatted line")
		}
		result = append(result, cleanMatch(line[indices[0]:indices[1]]))
		line = strings.TrimSpace(line[indices[1]:])
	}
	return result, nil
}

func cleanMatch(s string) string {
	s = strings.TrimSpace(s)
	if s[0:1] == "\"" {
		s = s[1 : len(s)-1]
	}
	s = strings.Replace(s, "\\\"", "\"", -1)
	s = strings.Replace(s, "\\\\", "\\", -1)
	return s
}
