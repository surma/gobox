package shell

import (
	flag "../../appletflag"
	"../../common"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func Main() {
	flag.Parse()
	call := flag.Args()
	var in *common.BufferedReader
	interactive := true
	if len(call) > 1 {
		call = call[0:1]
	}
	if len(call) == 1 {
		f, e := os.Open(call[0])
		if e != nil {
			log.Fatalf("Could not open input file %s: %s\n", call[1], e)
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
			log.Fatalf("Could not read line: %s\n", e)
		}
		if isComment(line) {
			continue
		}
		params, ce := common.Parameterize(line)
		if ce != nil {
			common.DumpError(ce)
			continue
		}
		params = expandEnvs(params)
		ce = execute(params)
		if ce != nil {
			common.DumpError(ce)
			continue
		}
	}
	return
}

// Replace environment variables with the content
func expandEnvs(params []string) []string {
	envReplaceFn := func(envVar string) string {
		return os.Getenv(envVar[1:])
	}

	envRe := regexp.MustCompile("([$]{1}[A-Z_]+)")
	for i, param := range params {
		params[i] = envRe.ReplaceAllStringFunc(param, envReplaceFn)
	}
	return params
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
