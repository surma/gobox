package common

import (
	"errors"
	"regexp"
	"strings"
)

const (
	unquotedParam = "([^ \t\"]+)"
	quotedParam   = "\"((\\\\.|[^\\\"])*)\""
	paramRegexp   = "^[ \t]*(" + unquotedParam + "|" + quotedParam + ")"
)

var (
	paramMatcher *regexp.Regexp = regexp.MustCompile(paramRegexp)
)

func Parameterize(line string) ([]string, error) {
	result := make([]string, 0)
	for len(line) != 0 {
		indices := paramMatcher.FindStringIndex(line)
		if indices == nil {
			println(paramRegexp)
			return nil, errors.New("Ill-formatted line")
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
