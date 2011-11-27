package common

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Checks if the given path exists
// i.e. if a os.Stat() succeeds
func PathExists(path string) bool {
	_, e := os.Stat(path)
	return e == nil

}

// Prints an error to stdout in a "nice" way.
func DumpError(e error) {
	FDumpError(os.Stdout, e)
}

// Prints an error to a writer in a "nice" way.
func FDumpError(w io.Writer, e error) {
	fmt.Fprintf(w, "gobox: Error: %s\n", e.Error())
}

// Creates a symlink and deletes the file blocking
// the name of the symlink.
func ForcedSymlink(oldname, newname string) error {
	if PathExists(newname) {
		e := os.Remove(newname)
		if e != nil {
			return e
		}
	}
	return os.Symlink(oldname, newname)
}

// Returns a slice of all pids currently existing
func GetAllPids() ([]int, error) {
	r := make([]int, 0)

	f, e := os.Open("/proc")
	if e != nil {
		e = errors.New("Could not open /proc")
		return nil, e
	}

	elems, e := f.Readdirnames(0)
	if e != nil {
		return nil, e
	}

	for _, elem := range elems {
		if IsNumeric(elem) {
			pid, _ := strconv.Atoi(elem)
			r = append(r, pid)
		}
	}
	return r, nil
}

type Process struct {
	*os.Process
	Name     string
	Owner    int
	Cmdline  string
	State    string
	Parent   int
	MemUsage int
}

func GetProcessByPid(pid int) (*Process, error) {
	var e error
	p := &Process{}
	p.Process, e = os.FindProcess(pid)
	if e != nil {
		return nil, e
	}

	data, e := readProcessStatusFile(pid)
	if e != nil {
		return nil, e
	}
	p.Name = data["Name"]
	p.Owner = getOwnerID(data["Uid"])
	p.State = data["State"]
	p.Parent = getParentPid(data["PPid"])
	p.MemUsage = getMemUsage(data["VmSize"])
	p.Cmdline, e = getCmdlineByPid(pid)
	if e != nil {
		return nil, e
	}
	return p, nil

}

func readProcessStatusFile(pid int) (map[string]string, error) {
	filename := fmt.Sprintf("/proc/%d/status", pid)
	f, e := os.Open(filename)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	vals := make(map[string]string)
	r := NewBufferedReader(f)
	for l, e := r.ReadWholeLine(); e == nil; {
		parts := strings.SplitN(l, ":", 2)
		vals[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		l, e = r.ReadWholeLine()
	}
	if e == io.EOF {
		e = nil
	}
	return vals, e
}

func getOwnerID(uid string) int {
	parts := strings.Split(uid, "\t")
	nuid, _ := strconv.Atoi(parts[0])
	return nuid
}

func getCmdlineByPid(pid int) (string, error) {
	filename := fmt.Sprintf("/proc/%d/cmdline", pid)
	f, e := os.Open(filename)
	if e != nil {
		return "", e
	}
	defer f.Close()

	r := NewBufferedReader(f)
	l, e := r.ReadWholeLine()
	if e != nil && e != io.EOF {
		return "", e
	}

	return l, nil
}

func getParentPid(ppid string) int {
	nppid, _ := strconv.Atoi(ppid)
	return nppid
}

func getMemUsage(vmsize string) (nvmsize int) {
	fmt.Sscanf(vmsize, "%d kB", &nvmsize)
	return
}
