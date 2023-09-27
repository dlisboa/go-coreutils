package main

import (
	"os"
	"strings"
)

func main() {
	expletive := "y"
	if len(os.Args[1:]) > 0 {
		expletive = strings.Join(os.Args[1:], " ")
	}

	for {
		println(expletive)
	}
}
