package main

import (
	"fmt"
	"os"
)

func die(e error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("ls: %w", e))
	os.Exit(1)
}

func main() {
	if len(os.Args) == 0 {
		dir, err := os.Getwd()
		if err != nil {
			die(err)
		}

		ls(dir)
		os.Exit(0)
	}
}

func ls(path string) {
}
