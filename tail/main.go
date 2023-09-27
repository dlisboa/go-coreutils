package main

import (
	"bufio"
	"container/ring"
	"errors"
	"flag"
	"fmt"
	"os"
)

var nflag = flag.Int("n", 10, "Print the last n lines of each of the specified files.")
var fflag = flag.Bool("f", false, "Does not stop when endof file is reached, waits for additional data")

func main() {
	flag.Parse()
	if *nflag <= 0 {
		die(errors.New("line count has to be positive"))
	}

	files := flag.Args()
	if len(files) == 0 {
		tail(os.Stdout, false)
		os.Exit(0)
	}

	for i, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("tail: %w", err))
			continue
		}

		tail(f, len(files) > 1)

		if i < len(files)-1 {
			fmt.Println()
		}
	}
}

// tail prints at most *nflag lines of a file, and optionally a header
func tail(f *os.File, header bool) {
	// uses a ring buffer when iterating over the lines to keep a sliding window
	// of lines in memory.
	ring := ring.New(*nflag)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ring.Value = scanner.Text()
		ring = ring.Next()
	}

	if header {
		fmt.Printf("==> %s <==\n", f.Name())
	}

	ring.Do(func(line any) {
		if line != nil {
			fmt.Println(line)
		}
	})
}

func die(e error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("tail: %w", e))
	os.Exit(1)
}
