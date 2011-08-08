package rm

import (
	"os"
)

func Rm(call []string) os.Error {
	if len(call) <= 1 {
		return os.NewError("`rm` <files...>")
	}
	for _, file := range call[1:] {
		e := os.Remove(file)
		if e != nil {
			return e
		}
	}
	return nil
}
