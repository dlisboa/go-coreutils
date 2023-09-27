# tail

Display last lines of a file.

    tail [-n count | -f ] [file ...]

- [X] Open one or more files from command line. Print error if any of them err
- [X] Print the last n lines of the file
- [X] Spit result to os.Stdout, in the order they came
- [X] If more than one file, write it out like so:
  ```sh
  $ tail -n 5 bar foo
  ==> bar <==
  hello
  frank
  ==> foo <==
  world$
  ```
- [] Keep reading from file if -f option is given. It combines with -n option
  ```sh
  $ tail -n 1 -f bar
  ==> bar <==
  hello
  # still reading
  ```
