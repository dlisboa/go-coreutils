package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Only works for Linux, macOS keeps logged user names in a different file

// given the line:
//
//	username:*:uid:guid:fullname:home:shell
//
// captures username and uid
var re = regexp.MustCompile(`^(.+?):.+?:(.+?):.*$`)

func main() {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("whoami: %w", err))
		os.Exit(1)
	}

	uid := os.Getuid()
	username := search(file, uid)

	if username == "" {
		fmt.Println(uid)
		os.Exit(0)
	}

	fmt.Println(username)
}

func search(file *os.File, uid int) string {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matches) == 0 {
			continue
		}

		name := matches[0][1]
		stringUID := matches[0][2]
		intUID, err := strconv.Atoi(stringUID)
		if err != nil {
			die(err)
		}

		if intUID == uid {
			return name
		}
	}

	return ""
}

func die(e error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("whoami: %w", e))
	os.Exit(1)
}
