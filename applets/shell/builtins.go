package shell

import (
	"os"
	"strconv"
)

type BuiltinHandler func(call []string) os.Error
var (
	Builtins map[string]BuiltinHandler = map[string]BuiltinHandler{
		"cd": cd,
		"pwd": pwd,
		"exit": exit,
	}
)

func pwd(call []string) os.Error {
	pwd, e := os.Getwd()
	if e != nil {
		return e
	}
	println(pwd)
	return nil
}

func cd(call []string) os.Error {
	if len(call) != 2 {
		return os.NewError("`cd` takes 1 paramter")
	}
	e := os.Chdir(call[1])
	return e
}

func exit(call []string) (e os.Error) {
	code := 0
	if len(call) >= 2 {
		code, e = strconv.Atoi(call[1])
		if e != nil {
			return e
		}
	}
	os.Exit(code)
	return nil
}
