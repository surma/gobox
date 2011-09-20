package shell

import (
	"common"
	"exec"
	"os"
	"strings"
)

func Shell(call []string) os.Error {
	var in *common.BufferedReader
	if len(call) > 2 {
		call = call[0:1]
	}
	if len(call) == 2 {
		f, e := os.Open(call[1])
		if e != nil {
			return e
		}
		defer f.Close()
		in = common.NewBufferedReader(f)
	} else {
		in = common.NewBufferedReader(os.Stdin)
	}

	var e os.Error
	var line string
	for e == nil {
		print("> ")
		line, e = in.ReadWholeLine()
		if e != nil {
			return e
		}
		if isComment(line) {
			continue
		}
		params, ce := common.Parameterize(line)
		if ce != nil {
			common.DumpError(ce)
			continue
		}
		ce = execute(params)
		if ce != nil {
			common.DumpError(ce)
			continue
		}
	}
	return nil
}

func isComment(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, "#")
}

func execute(cmd []string) os.Error {
	if len(cmd) == 0 {
		return nil
	}
	if isBuiltIn(cmd[0]) {
		builtin := Builtins[cmd[0]]
		return builtin(cmd)
	} else {
		cmd := exec.Command(cmd[0], cmd[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		e := cmd.Run()
		return e
	}
	return nil
}

func isBuiltIn(cmd string) bool {
	_, ok := Builtins[cmd]
	return ok
}
