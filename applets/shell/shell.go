package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/surma/gobox/pkg/common"
)

func Shell(call []string) error {
	var in *common.BufferedReader
	interactive := true
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
		interactive = false
	} else {
		in = common.NewBufferedReader(os.Stdin)
	}

	var e error
	var line string
	for e == nil {
		if interactive {
			fmt.Print("> ")
		}
		line, e = in.ReadWholeLine()
		if e != nil {
			return e
		}
		if isComment(line) {
			continue
		}
		line = expandVariables(line)
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

func expandVariables(line string) string {
	var expanded []string
	var isquote bool
	words := strings.Split(line, " ")
	for _, word := range words {
		isquote = false
		if strings.HasPrefix(word, `'`) {
			isquote = true
		}
		if !isquote && strings.HasPrefix(word, "$") {
			word = os.ExpandEnv(word)
		}
		expanded = append(expanded, word)
	}
	return strings.Join(expanded, " ")
}

func isComment(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, "#")
}

func execute(cmd []string) error {
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
