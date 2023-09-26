# head

Display first lines (or bytes) of a file.

    head [-n count | -c bytes] [file ...]

- [] Open one or more files from command line. Exit if any of them err
- [] Read at most n lines from the file
- [] Read at most c bytes from the file (exclusive, either this or lines)
- [] Spit result to os.Stdout, in the order they came
- [] If more than one file, write it out like so:
  ```sh
  $ head -c 5 bar foo
  ==> bar <==
  hello
  ==> foo <==
  world$
  ```
- [] Optimization: don't read all before spitting, do chunks at a time
