package main

import (
	"flag"
	"fmt"
	"os"
)

var rflag = flag.Bool("r", false, "remove recursively")
var fflag = flag.Bool("f", false, "force removal")

func main() {
	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "usage: rm [-r] [-f] file ...\n")
		os.Exit(64) // EX_USAGE, check `man sysexits`
	}

	force := *fflag
	recur := *rflag
	for _, file := range files {
		err := rm(file, recur, force)
		if err != nil {
			printErr(err)
			continue
		}
	}
}

func printErr(e error) {
	fmt.Fprintf(os.Stderr, "rm: %s\n", e)
}

func rm(file string, recur, force bool) (err error) {
	info, err := os.Stat(file)
	switch {
	case err != nil && force:
		return nil
	case err != nil && !force:
		return err
	case info.IsDir() && !recur:
		return fmt.Errorf("%s: is a directory", info.Name())
	}

	if recur {
		err = os.RemoveAll(file)
	} else {
		err = os.Remove(file)
	}

	return err
}
