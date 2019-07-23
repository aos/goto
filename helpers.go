package main

import (
	"flag"
	"strings"
)

// CommandLineArgs holds all the command-line flag and args information
type CommandLineArgs struct {
	Add  bool
	List bool
	Init bool
	Arg  string
}

// Check checks for errors and just panics
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Split takes a string with a single separator,
// splits and returns both variables
func Split(s, sep string) (string, string) {
	if len(s) == 0 {
		return s, s
	}

	slice := strings.SplitN(s, sep, 2)
	if len(slice) == 1 {
		return slice[0], ""
	}

	return slice[0], slice[1]
}

// ParseCommandLine initializes our command-line flags and args
func ParseCommandLine() CommandLineArgs {
	addPtr := flag.Bool("a", false, "Bind current directory to name (default: basename)")
	listPtr := flag.Bool("l", false, "Lists the currently installed shortcuts")
	initPtr := flag.Bool("init", false, "Prints out the Bash integration code")
	flag.Parse()
	arg := flag.Arg(0)
	return CommandLineArgs{*addPtr, *listPtr, *initPtr, arg}
}
