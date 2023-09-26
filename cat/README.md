# cat

Concatenate files.

- [] Open one or more files from command line. Exit if any of them err
- [] Read these files into a buffer
- [] Concatenate them and spit to os.Stdout, in the order they came
- [] Optimization: don't read all before spitting, do chunks at a time
