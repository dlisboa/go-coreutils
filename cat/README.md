# cat

Concatenate files.

- [X] Open one or more files from command line. Exit if any of them err
- [X] Read these files into a buffer
- [X] Concatenate them and spit to os.Stdout, in the order they came
- [X] Optimization: don't read all before spitting, do chunks at a time
