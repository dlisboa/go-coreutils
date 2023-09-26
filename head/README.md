# head

Display first lines (or bytes) of a file.

    head [-n count | -c bytes] [file ...]

- [X] Open one or more files from command line. Exit if any of them err
- [X] Read at most n lines from the file
- [X] Read at most c bytes from the file (exclusive, either this or lines)
  ```sh
  $ head -n 1 -c 20 main.go
  head: can't combine line and byte counts
  ```
- [X] Spit result to os.Stdout, in the order they came
- [X] If more than one file, write it out like so:
  ```sh
  $ head -c 5 bar foo
  ==> bar <==
  hello
  ==> foo <==
  world$
  ```
- [X] Optimization: don't read all before spitting, do chunks at a time
