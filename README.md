## goto

Jump between different directories using shortcuts.

### Why

I frequently jump around between very nested directories and it would be nice
to have some shortcut to do that. There are already existing tools that do this
(much better), but as an exercise I wanted to roll my own.

### How

As we cannot change directories from within the tool itself since it is spawned
as a subprocess that does not have access to our shell, we must integrate it
and pass the directory upwards to our shell to `cd`.

The implementation works by generating a file `~/.gotorc` that stores a
shortcut to directory mapping. The tool generates a shell (bash-only right now)
script that is called when `goto --init` is run. This is `eval`'d in a
bash startup script (`~/.bashrc` or `~/.bash_profile`), and creates a function
that reads the directory and `cd`'s into the directory.

### Usage

1. `go get github.com/aos/goto`
2. Add `eval $(goto --init)` in your `~/.bashrc` or `~/.bash_profile`
3. Run it and follow the usage instructions: `g`

### Todo

- [ ] Custom shortcut keybinding
- [ ] Shortcut removal
