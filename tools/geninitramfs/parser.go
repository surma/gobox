package main

import (
	"common"
	"github.com/surma/gocpio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(in io.ReadCloser, c chan<- *Entry) {
	buf := common.NewBufferedReader(in)
	for {
		line, e := buf.ReadWholeLine()
		if e != nil && e != io.EOF {
			log.Printf("Warning: Could not read whole file: %s", e.Error())
		}
		stripped_line := strings.TrimSpace(line)
		if len(stripped_line) == 0 || stripped_line[0] == '#' {
			if e != io.EOF {
				continue
			} else {
				break
			}
		}
		ent := parseLine(line)
		if ent != nil {
			c <- ent
		}

		if e == io.EOF {
			break
		}
	}
	close(c)
}

func parseLine(line string) *Entry {
	lineparts := strings.Split(line, " ")
	if len(lineparts) < 1 {
		return nil
	}

	switch lineparts[0] {
	case "file":
		return parseFile(lineparts[1:])
	case "dir":
		return parseDir(lineparts[1:])
	case "nod":
		return parseNod(lineparts[1:])
	case "slink":
		return parseSlink(lineparts[1:])
	case "pipe":
		return parsePipe(lineparts[1:])
	case "sock":
		return parseSock(lineparts[1:])
	default:
		log.Printf("Warning: %s is in invalid type\n", lineparts[0])
	}
	return nil
}

func parseFile(parts []string) *Entry {
	if len(parts) != 5 {
		return nil
	}
	name := parts[0]
	localname := parts[1]
	mode, uid, gid, e := parseModeUidGid(parts[2], parts[3], parts[4])
	if e != nil {
		log.Printf("Invalid permission settings: %s\n", e.Error())
		return nil
	}

	f, e := os.Open(localname)
	if e != nil {
		log.Printf("Could not open file %s: %s\n", localname, e.Error())
		return nil
	}
	finfo, e := f.Stat()
	if e != nil {
		log.Printf("Could not obtain file size of %s: %s\n")
		return nil
	}

	return &Entry{
		hdr: cpio.Header{
			Mode: mode,
			Uid:  uid,
			Gid:  gid,
			Size: finfo.Size(),
			Type: cpio.TYPE_REG,
			Name: name,
		},
		data: f,
	}
}

func parseDir(parts []string) *Entry {
	if len(parts) != 4 {
		return nil
	}
	name := parts[0]
	mode, uid, gid, e := parseModeUidGid(parts[1], parts[2], parts[3])
	if e != nil {
		log.Printf("Invalid permission settings: %s\n", e.Error())
		return nil
	}

	return &Entry{
		hdr: cpio.Header{
			Mode: mode,
			Uid:  uid,
			Gid:  gid,
			Type: cpio.TYPE_DIR,
			Name: name,
		},
	}
}

func parseNod(parts []string) *Entry {
	if len(parts) != 7 {
		return nil
	}
	name := parts[0]
	mode, uid, gid, e := parseModeUidGid(parts[1], parts[2], parts[3])
	if e != nil {
		log.Printf("Invalid permission settings: %s\n", e.Error())
		return nil
	}

	var dev_type int64
	switch parts[4] {
	case "b":
		dev_type = cpio.TYPE_BLK
	case "c":
		dev_type = cpio.TYPE_CHAR
	default:
		log.Printf("Invalid device type: %s\n", parts[4])
		return nil
	}

	maj, e := strconv.ParseInt(parts[5], 10, 64)
	if e != nil {
		log.Printf("Invalid major device: %s\n", e.Error())
		return nil
	}
	min, e := strconv.ParseInt(parts[6], 10, 64)
	if e != nil {
		log.Printf("Invalid major device: %s\n", e.Error())
		return nil
	}

	return &Entry{
		hdr: cpio.Header{
			Mode:     mode,
			Uid:      uid,
			Gid:      gid,
			Type:     dev_type,
			Devmajor: maj,
			Devminor: min,
			Name:     name,
		},
	}

}

func parseSlink(parts []string) *Entry {
	if len(parts) != 5 {
		return nil
	}

	name := parts[0]
	target := parts[1]

	mode, uid, gid, e := parseModeUidGid(parts[2], parts[3], parts[4])
	if e != nil {
		log.Printf("Invalid permission settings: %s\n", e.Error())
		return nil
	}

	return &Entry{
		hdr: cpio.Header{
			Mode: mode,
			Uid:  uid,
			Gid:  gid,
			Type: cpio.TYPE_SYMLINK,
			Size: int64(len(target) + 1),
			Name: name,
		},
		data: strings.NewReader(target),
	}

}

func parsePipe(parts []string) *Entry {
	if len(parts) != 4 {
		return nil
	}
	name := parts[0]
	mode, uid, gid, e := parseModeUidGid(parts[1], parts[2], parts[3])
	if e != nil {
		log.Printf("Invalid permission settings: %s\n", e.Error())
		return nil
	}

	return &Entry{
		hdr: cpio.Header{
			Mode: mode,
			Uid:  uid,
			Gid:  gid,
			Type: cpio.TYPE_FIFO,
			Name: name,
		},
	}
}

func parseSock(parts []string) *Entry {
	if len(parts) != 4 {
		return nil
	}
	name := parts[0]
	mode, uid, gid, e := parseModeUidGid(parts[1], parts[2], parts[3])
	if e != nil {
		log.Printf("Invalid permission settings: %s\n", e.Error())
		return nil
	}

	return &Entry{
		hdr: cpio.Header{
			Mode: mode,
			Uid:  uid,
			Gid:  gid,
			Type: cpio.TYPE_SOCK,
			Name: name,
		},
	}
}

func parseModeUidGid(s_mode, s_uid, s_gid string) (int64, int, int, error) {
	mode, err := strconv.ParseInt(s_mode, 0, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	tuid, err := strconv.ParseInt(s_uid, 0, 0)
	if err != nil {
		return 0, 0, 0, err
	}
	tgid, err := strconv.ParseInt(s_gid, 0, 0)
	return mode, int(tuid), int(tgid), nil
}
