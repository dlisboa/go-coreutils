package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func usage() {
	fmt.Println("usage: mv source target")
	fmt.Println("       mv source ... target")
}

func main() {
	files := os.Args[1:]

	var err error
	switch {
	case len(files) < 2:
		usage()
		os.Exit(64) // EX_USAGE, check `man sysexits`
	case len(files) == 2:
		err = rename(files[0], files[1])
	default:
		err = mv(files...)
	}

	if err != nil {
		die(err)
	}
}

func mv(files ...string) error {
	target := files[len(files)-1]

	dir, err := os.Open(target)
	if err != nil {
		return err
	}
	info, err := dir.Stat()
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", dir.Name())
	}

	sources := files[:len(files)-1]
	for _, source := range sources {
		dest := filepath.Join(target, path.Base(source))
		err = os.Rename(source, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func rename(old, new string) error {
	err := os.Rename(old, new)
	if err != nil {
		return err
	}
	return nil
}

func die(e error) {
	fmt.Fprintf(os.Stderr, "mv: %s\n", e)
	os.Exit(1)
}
