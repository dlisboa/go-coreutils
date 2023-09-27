package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	switch {
	case len(args) < 1 || len(args) > 2:
		fmt.Println(`usage: basename path [suffix]`)
		os.Exit(1)
	case len(args) == 1:
		basename(args[0], "")
	case len(args) == 2:
		basename(args[0], args[1])
	}
}

func basename(path, suffix string) {
	s := filepath.Base(path)
	if suffix != "" {
		s, _ = strings.CutSuffix(s, suffix)
	}
	fmt.Println(s)
}
