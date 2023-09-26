package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
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

	var printHeader bool
	if len(files) > 1 {
		printHeader = true
	}
	for _, f := range files {
		f, err := os.Open(f)
		if err != nil {
			die(err)
		}
		head(f, printHeader)
	}
}

func die(e error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("head: %w", e))
	os.Exit(1)
}

func checkFlags() {
	if flag.NFlag() > 1 {
		die(errors.New("can't combine line and byte counts"))
	}

	if flag.NFlag() == 0 {
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
		if count == *nflag {
			break
		}
		fmt.Println(scanner.Text())
		count++
	}
}

func readBytes(f *os.File) {
	buf := make([]byte, *cflag)
	f.Read(buf)
	fmt.Println(string(buf))
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
