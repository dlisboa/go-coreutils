package main

import (
	"bufio"
	"fmt"
	"os"
)

func cat(f *os.File) {
	reader := bufio.NewReaderSize(f, 4096) // read 4 KB at a time
	reader.WriteTo(os.Stdout)
}

func die(e error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("cat: %w", e))
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		cat(os.Stdout)
		os.Exit(0)
	}

	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			die(err)
		}

		info, err := f.Stat()
		if err != nil {
			die(err)
		}
		if info.IsDir() {
			die(fmt.Errorf("%s: is a directory", f.Name()))
		}

		cat(f)
	}
}
