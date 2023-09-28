# ls

List files.

    ls [-l] [-a] [path] [path...]

- [] If no path is given, list from current directory
- [] Path can be a glob: *, **, foo*. List all the files that match it
- [] Order files alphabetically
- [] When output fits in the number of columns, print in one line
- [] When output overflow the columns, print in columns
- [] When multiple paths given, print header:
  ```sh
  $ ls /usr /etc
  /usr:
  bin lib libexec ...

  /etc:
  aliases
  apache2
  ...
  ```
- [] If -a is set, include hidden files
- [] If -l is set, display expanded version in columns:
  ```sh
  $ ls -l
  total 16 # can disregard this for now
  -rw-r--r--  1 diogo  staff  277 Sep 27 23:26 README.md
  -rw-r--r--  1 diogo  staff  517 Sep 27 23:26 main.go
  ```
  From the ls man page on "The Long Format", these columns mean:
  > file mode, number of links, owner name, group name, number of bytes in the
  > file, abbreviated month, day-of-month file was last modified, hour file
  > last modified, minute file last modified, and the pathname.
- [] Lower priority: the "total 16" part of the Long Format

