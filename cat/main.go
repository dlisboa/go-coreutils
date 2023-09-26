package main

import (
	"fmt"
	"io"
	"os"
)

func cat(f *os.File) {
	buf := make([]byte, 4096) // read 4 KB at a time

	for {
		_, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			die(err)
		}
		fmt.Printf("%s", buf)
	}
}

func die(e error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("cat: %w", e))
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		cat(os.Stdout)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				die(err)
			}
			cat(f)
		}
	}
}
