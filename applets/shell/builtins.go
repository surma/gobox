package shell

import (
	"os"
	"strconv"
)

type BuiltinHandler func(call []string) os.Error

var (
	Builtins map[string]BuiltinHandler = map[string]BuiltinHandler{
		"cd":       cd,
		"pwd":      pwd,
		"exit":     exit,
		"env":      env,
		"getenv":   getenv,
		"setenv":   setenv,
		"unsetenv": unsetenv,
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
		return os.NewError("`cd <directory>`")
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

func env(call []string) os.Error {
	for _, envvar := range os.Environ() {
		println(envvar)
	}
	return nil
}

func getenv(call []string) os.Error {
	if len(call) != 2 {
		return os.NewError("`getenv <variable name>`")
	}
	println(os.Getenv(call[1]))
	return nil
}

func setenv(call []string) os.Error {
	if len(call) != 3 {
		return os.NewError("`setenv <variable name> <value>`")
	}
	return os.Setenv(call[1], call[2])
}

func unsetenv(call []string) os.Error {
	if len(call) != 2 {
		return os.NewError("`unsetenv <variable name>`")
	}
	return os.Setenv(call[1], "")
}
