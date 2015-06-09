package ls

import (
	flag "../../appletflag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"text/tabwriter"
	"time"
)

var (
	longFlag      = flag.Bool("l", false, "Long, detailed listing")
	recursiveFlag = flag.Bool("r", false, "Recurse into directories")
	humanFlag     = flag.Bool("h", false, "Output sizes in a human readable format")
	helpFlag      = flag.Bool("help", false, "Show this help")
	out           = tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
)

func Main() {
	flag.Parse()

	if *helpFlag {
		println("`ls` [options] [dirs...]")
		flag.PrintDefaults()
		return
	}

	dirs, e := getDirList()
	if e != nil {
		log.Fatalf("Could not get cwd: %s\n", e)
	}

	for _, dir := range dirs {
		e := list(dir, "")
		if e != nil {
			log.Printf("Error while listing directory: %s\n", e)
		}
	}
	out.Flush()
	return
}

func getDirList() ([]string, error) {
	if flag.NArg() <= 0 {
		cwd, e := os.Getwd()
		return []string{cwd}, e
	}
	return flag.Args(), nil
}

func list(dir, prefix string) error {
	entries, e := ioutil.ReadDir(dir)
	if e != nil {
		return e
	}

	for _, entry := range entries {
		printEntry(entry)
		if entry.IsDir() && *recursiveFlag {
			folder := prefix + "/" + entry.Name()
			fmt.Fprintf(out, "%s:\n", folder)
			e := list(dir+"/"+entry.Name(), folder)
			if e != nil {
				log.Printf("Failed listing %s: %s", entry.Name(), e)
				continue
			}
		}
	}
	return nil
}

func printEntry(e os.FileInfo) {
	if *longFlag {
		if e.IsDir() {
			fmt.Print("d")
		} else {
			fmt.Print("-")
		}
		fmt.Fprintf(out, "%s ", getModeString(e.Mode()))
		fmt.Fprintf(out, "%s ", getSizeString(e.Size()))
		fmt.Fprintf(out, "%16s ", getTimeString(e.ModTime()))
	}
	fmt.Fprintf(out, "%s%s", e.Name(), getEntryTypeString(e))
	fmt.Fprintln(out, "")
}

func getTimeString(mtime time.Time) (s string) {
	return mtime.Format("2006-01-02 15:04")
}

var accessSymbols = "xwr"

func getModeString(mode os.FileMode) (s string) {
	for i := 8; i >= 0; i-- {
		if uint32(mode)&(1<<uint(i)) == 0 {
			s += "-"
		} else {
			char := i % 3
			s += accessSymbols[char : char+1]
		}
	}
	return
}

var sizeSymbols = "BkMGT"

func getSizeString(size int64) (s string) {
	if !*humanFlag {
		return fmt.Sprintf("%6dB", size)
	}
	var power int
	if size == 0 {
		power = 0
	} else {
		power = int(math.Log(float64(size)) / math.Log(1024.0))
	}
	if power > len(sizeSymbols)-1 {
		power = len(sizeSymbols) - 1
	}
	rSize := float64(size) / math.Pow(1024, float64(power))
	return fmt.Sprintf("% 6.1f%s", rSize, sizeSymbols[power:power+1])
}

func getEntryTypeString(e os.FileInfo) string {
	if e.IsDir() {
		return "/"
	} else if e.IsBlock() {
		return "<>"
	} else if e.Mode()&os.ModeNamedPipe != 0 {
		return ">>"
	} else if e.Mode()&os.ModeSymlink != 0 {
		return "@"
	} else if e.Mode()&os.ModeSocket != 0 {
		return "&"
	} else if e.Mode().IsRegular() && (e.Mode()&0001 == 0001) {
		return "*"
	}
	return ""
}

func getUserString(id int) string {
	return fmt.Sprintf("%03d", id)
}
