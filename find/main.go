package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

var nflag = flag.String("name", "", "filters by file names that match the pattern")
var tflag = flag.String("type", "", "filters by files that are of type [d f l]")

const usage = "usage: find [-name pattern] [-type type] path [path ...]"

func main() {
	flag.Parse()
	paths := flag.Args()
	if len(paths) < 1 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	for _, dir := range paths {
		filepath.WalkDir(dir, printFile)
	}
}

func printFile(file string, d fs.DirEntry, err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "find: %s\n", err)
		return nil
	}

	entry := DirEntry{d}
	shouldPrint := entry.MatchPattern(*nflag) && entry.MatchType(*tflag)
	if shouldPrint {
		fmt.Printf("%s\n", file)
	}

	return nil
}

type DirEntry struct {
	fs.DirEntry
}

func (d DirEntry) MatchPattern(pattern string) bool {
	if pattern == "" {
		return true
	}

	matched, err := path.Match(pattern, d.Name())
	if err != nil {
		return false
	}
	return matched
}

func (d DirEntry) MatchType(typ string) bool {
	switch typ {
	case "d":
		return d.IsDir()
	case "l":
		return d.IsSymlink()
	case "f":
		return d.IsRegular()
	case "":
		return true
	default:
		return false
	}
}

func (d DirEntry) IsRegular() bool {
	return !d.IsDir() && !d.IsSymlink()
}

func (d DirEntry) IsSymlink() bool {
	return (d.Type() & fs.ModeSymlink) == fs.ModeSymlink
}
