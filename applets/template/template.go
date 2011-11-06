package template

import "flag"

var (
	flagSet  = flag.NewFlagSet("template", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
)

func Template(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 1 || *helpFlag {
		println("`template` [options] <parameter1>")
		flagSet.PrintDefaults()
		return nil
	}
	return nil
}
