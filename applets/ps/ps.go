package ps

import (
	"common"
	flag "appletflag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

var (
	helpFlag = flag.Bool("help", false, "Show this help")
	out      = tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
)

func Main() {
	flag.Parse()

	if flag.NArg() != 0 || *helpFlag {
		println("`ps` [options]")
		flag.PrintDefaults()
		return
	}

	pids, e := common.GetAllPids()
	if e != nil {
		log.Fatalf("Could not obtain PIDs: %s\n", e)
	}

	fmt.Fprintf(out, "Pid\tParent\tState\tOwner\tMem (kB)\tName\t\n")
	for _, pid := range pids {
		proc, e := common.GetProcessByPid(pid)
		if e != nil {
			log.Printf("Could not get process info of %d: %s\n", pid, e)
			continue
		}
		printProcess(proc)
	}

	out.Flush()
	return
}

func printProcess(p *common.Process) {
	fmt.Fprintf(out, "%d\t%d\t%s\t%d\t%d\t%s\t\n", p.Process.Pid, p.Parent, p.State, p.Owner, p.MemUsage, p.Name)
}
