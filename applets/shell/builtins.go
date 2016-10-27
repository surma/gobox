package shell

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type BuiltinHandler func(call []string) error

var (
	Builtins map[string]BuiltinHandler
)

func init() {
	Builtins = map[string]BuiltinHandler{
		"cd":       cd,
		"pwd":      pwd,
		"exit":     exit,
		"env":      env,
		"getenv":   getenv,
		"setenv":   setenv,
		"unsetenv": unsetenv,
		"fork":     fork,
		"sleep":    sleep,
	}
}

func pwd(call []string) error {
	pwd, e := os.Getwd()
	if e != nil {
		return e
	}
	println(pwd)
	return nil
}

func cd(call []string) error {
	homedir := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
	if homedir == "" {
		homedir = os.Getenv("USERPROFILE")
	}
	if homedir == "" {
		homedir = os.Getenv("HOME")
	}
	if len(call) == 1 {
		e := os.Chdir(homedir)
		return e
	}
	if call[1] == "~" && len(call) == 2 {
		e := os.Chdir(homedir)
		return e
	}
	if len(call) != 2 {
		return errors.New("`cd <directory>`")
	}
	e := os.Chdir(call[1])
	return e
}

func exit(call []string) (e error) {
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

func env(call []string) error {
	for _, envvar := range os.Environ() {
		println(envvar)
	}
	return nil
}

func getenv(call []string) error {
	if len(call) != 2 {
		return errors.New("`getenv <variable name>`")
	}
	println(os.Getenv(call[1]))
	return nil
}

func setenv(call []string) error {
	if len(call) != 3 {
		return errors.New("`setenv <variable name> <value>`")
	}
	return os.Setenv(call[1], call[2])
}

func unsetenv(call []string) error {
	if len(call) != 2 {
		return errors.New("`unsetenv <variable name>`")
	}
	return os.Setenv(call[1], "")
}

func fork(call []string) error {
	if len(call) < 2 {
		return errors.New("`fork <command...>`")
	}
	go execute(call[1:])
	return nil
}

func sleep(call []string) error {
	if len(call) < 2 {
		return errors.New("`sleep duration(S)`")
	}
	t, e := strconv.ParseFloat(call[1], 64)
	if e != nil {
		return e
	}
	d := time.Duration(t * 1000 * 1000 * 1000)
	time.Sleep(d)
	return nil
}
