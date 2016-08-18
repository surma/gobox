package gzip

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
)

var (
	flagSet    = flag.NewFlagSet("gzip", flag.PanicOnError)
	helpFlag   = flagSet.Bool("help", false, "Show this help")
	forceFlag  = flagSet.Bool("f", false, "Force decompression (ignore extension)")
	decompress = flagSet.Bool("d", false, "Decompress")
)

func Gzip(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() < 1 || *helpFlag {
		println("`gzip` <file>...")
		flagSet.PrintDefaults()
		return nil
	}

	if *decompress {
		Gunzip(call)
		return nil
	}

	for _, fn := range flagSet.Args() {
		doGzip(fn)
	}

	return nil
}

func Gunzip(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSet.NArg() < 1 || *helpFlag {
		println("`gunzip` <file>...")
		flagSet.PrintDefaults()
		return nil
	}

	for _, fn := range flagSet.Args() {
		doGunzip(fn)
	}
	return nil
}

var (
	flagSetZcat  = flag.NewFlagSet("zcat", flag.PanicOnError)
	helpFlagZcat = flagSetZcat.Bool("help", false, "Show this help")
)

func Zcat(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}

	if flagSetZcat.NArg() < 1 || *helpFlagZcat {
		println("`zcat` <file>...")
		flagSetZcat.PrintDefaults()
		return nil
	}

	for _, fn := range flagSetZcat.Args() {
		doZcat(fn)
	}

	return nil
}

func doGzip(fn string) error {
	fh, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fn, err)
		return err
	}
	defer fh.Close()
	fi, err := fh.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fn, err)
		return err
	}
	if fi.IsDir() {
		fmt.Fprintf(os.Stderr, "%s: is a directory\n", fn)
		return err
	}
	newfn := fn + ".gz"
	tfh, err := os.OpenFile(newfn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, fi.Mode().Perm())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", newfn, err)
		return err
	}
	compressor, err := gzip.NewWriterLevel(tfh, gzip.BestCompression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gzip: %v\n", err)
		return err
	}
	defer compressor.Close()
	compressor.ModTime = fi.ModTime()
	compressor.Name = fn
	compressor.OS = 3 // Unix
	if _, err := io.Copy(compressor, fh); err != nil {
		fmt.Fprintf(os.Stderr, "gzip: %v\n", err)
	}
	if err := os.Remove(fn); err != nil {
		fmt.Fprintf(os.Stderr, "gzip: %v\n", err)
	}
	return nil
}

func doGunzip(fn string) error {
	fh, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fn, err)
		return err
	}
	decompressor, err := gzip.NewReader(fh)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fn, err)
		return err
	}
	defer decompressor.Close()
	fi, err := fh.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fn, err)
		return err
	}
	if path.Ext(fn) != ".gz" && !*forceFlag {
		fmt.Fprintf(os.Stderr, "gunzip: %v: unknown suffix -- ignored\n", fn)
		return nil
	}
	newfn := fn + ".gunzip"
	if !*forceFlag {
		newfn = fn[0 : len(fn)-3]
	}
	tfh, err := os.OpenFile(newfn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, fi.Mode().Perm())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", newfn, err)
		return err
	}
	defer tfh.Close()
	if _, err := io.Copy(tfh, decompressor); err != nil {
		fmt.Fprintf(os.Stderr, "gunzip: %v\n", err)
	}
	if err := os.Remove(fn); err != nil {
		fmt.Fprintf(os.Stderr, "gunzip: %v\n", err)
	}
	return nil
}

func doZcat(fn string) error {
	fh, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fn, err)
		return err
	}
	decompressor, err := gzip.NewReader(fh)
	defer decompressor.Close()
	if _, err := io.Copy(os.Stdout, decompressor); err != nil {
		fmt.Fprintf(os.Stderr, "zcat: %v\n", err)
	}
	return nil
}
