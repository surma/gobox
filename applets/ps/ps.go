package ps

import (
	"flag"
	"fmt"
	"github.com/surma/gobox/pkg/common"
	"os"
	"text/tabwriter"
)

var (
	flagSet  = flag.NewFlagSet("ps", flag.PanicOnError)
	helpFlag = flagSet.Bool("help", false, "Show this help")
	out      = tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
)

func Ps(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() != 0 || *helpFlag {
		println("`ps` [options]")
		flagSet.PrintDefaults()
		return nil
	}

	pids, e := common.GetAllPids()
	if e != nil {
		return e
	}

	fmt.Fprintf(out, "Pid\tParent\tState\tOwner\tMem (kB)\tName\t\n")
	for _, pid := range pids {
		proc, e := common.GetProcessByPid(pid)
		if e != nil {
			return e
		}
		printProcess(proc)
	}

	out.Flush()
	return nil
}

func printProcess(p *common.Process) {
	fmt.Fprintf(out, "%d\t%d\t%s\t%d\t%d\t%s\t\n", p.Process.Pid, p.Parent, p.State, p.Owner, p.MemUsage, p.Name)
}
