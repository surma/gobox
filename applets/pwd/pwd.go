package pwd

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagSet  = flag.NewFlagSet("pwd", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

//Pwd outputs the current working directory
func Pwd(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() > 0 {
		fmt.Println("Expected 0 args, got 1")
		return nil

	}
	if *helpFlag {
		println("`pwd`")
		flagSet.PrintDefaults()
		return nil
	}

	pwd, _ := os.Getwd()
	println(pwd)
	return nil
}
