package common

import (
	"unicode"
)

func IsNumeric(s string) bool {
	for _, rune := range s {
		if !unicode.IsDigit(rune) {
			return false
		}
	}
	return true
}
