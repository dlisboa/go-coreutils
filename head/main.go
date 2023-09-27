package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var nflag = flag.Int("n", 10, "Print n lines of each of the specified files.")
var nflagGiven = false
var cflag = flag.Int("c", 10, "Print n bytes of each of the specified files.")
var cflagGiven = false

func main() {
	flag.Parse()
	checkFlags()

	// non-flag arguments, in our case the files
	files := flag.Args()
	if len(files) == 0 {
		head(os.Stdout, false)
		os.Exit(0)
	}

	for i, f := range files {
		f, err := os.Open(f)
		if err != nil {
			printerr(err)
			continue
		}
		head(f, len(files) > 1)
		// print new line between files unless it's the last one
		if i < len(files)-1 {
			fmt.Println()
		}
	}
}

func printerr(e error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("head: %w", e))
}

func die(e error) {
	printerr(e)
	os.Exit(1)
}

func checkFlags() {
	switch {
	case flag.NFlag() > 1:
		die(errors.New("can't combine line and byte counts"))
	case *nflag <= 0:
		die(fmt.Errorf("illegal line count -- %d", *nflag))
	case *cflag <= 0:
		die(fmt.Errorf("illegal byte count -- %d", *cflag))
	case flag.NFlag() == 0: // no flags at command line
		nflagGiven = true
		return
	}

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "n":
			nflagGiven = true
		case "c":
			cflagGiven = true
		}
	})
}

func readLines(f *os.File) {
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
		if count == *nflag {
			break
		}
	}
}

func readBytes(f *os.File) {
	reader := bufio.NewReader(f)
	buf := make([]byte, min(*cflag, 4096))
	count := 0
	for count < *cflag {
		n, err := reader.Read(buf)
		if err == io.EOF {
			return
		}
		if err != nil {
			die(err)
		}
		count += n
		fmt.Printf("%s", buf)
	}
}

func head(f *os.File, header bool) {
	if header {
		fmt.Printf("==> %s <==\n", f.Name())
	}
	switch {
	case nflagGiven:
		readLines(f)
	case cflagGiven:
		readBytes(f)
	}
}
