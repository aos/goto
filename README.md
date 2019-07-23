## Goto

Jump between different directories using shortcuts!

### Why
The idea behind this is that I frequently jump around between very nested
directories and it would be nice to have some shortcut to do that. There are
already existing tools that do this (much better), but as an exercise I wanted
to roll my own.

### How
The implementation works by generating a file `~/.gotorc` that stores a
shortcut to directory mapping. The tool creates an auto-generated shell
(bash-only right now) text that is called when you run `goto --init`. When this
is `eval`'d in a bash startup script (`~/.bashrc` or `~/.bash_profile`), it
creates a function that reads the directory and `cd`'s into the directory.

### Usage

1. `go get github.com/aos/goto`
2. Add `eval $(goto --init)` in your `~/.bashrc` or `~/.bash_profile`
3. Run it and follow the usage instructions: `$ g`
