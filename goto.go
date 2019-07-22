package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var jumpDirMap = make(map[string]string)

func main() {
	cwd, err := os.Getwd()
	check(err)

	home, err := os.UserHomeDir()
	check(err)

	base := filepath.Base(cwd)

	f, err := os.OpenFile(
		filepath.Join(home, ".gotorc"),
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		0644,
	)
	check(err)
	defer f.Close()

	// Create a map of our .gotorc file
	// { "shortcut": "/path/to/directory/" }
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		j, dir := split(scanner.Text(), ",")
		jumpDirMap[j] = dir
	}

	// Our flags
	name, add, list := makeFlags(base)
	jump := flag.Arg(0)

	// Oops -- this is clearly not going to work. I forgot that each
	// executable runs in its own process and therefore cannot change the
	// current working directory of the shell.
	// To fix this -- we need to integrate with the shell
	if jump != "" {
		if dir, ok := jumpDirMap[jump]; ok {
			fmt.Println(dir)
			//cmd := exec.Command("cd", dir)
			//err := cmd.Run()
			//check(err)
			return
		}
		fmt.Println("Shortcut not found.")
		return
	}

	if list {
		fmt.Println("Currently installed shortcuts:")
		for k, v := range jumpDirMap {
			fmt.Printf("%v\t%v\n", k, v)
		}
		return
	}

	if add {
		_, err := f.WriteString(fmt.Sprintf("%v,%v\n", base, cwd))
		check(err)
		return
	}

	if name != "" && name != base {
		_, err := f.WriteString(fmt.Sprintf("%v,%v\n", name, cwd))
		check(err)
		return
	}

	flag.Usage()
	return
}

// Helpers
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func split(s, sep string) (string, string) {
	if len(s) == 0 {
		return s, s
	}

	slice := strings.SplitN(s, sep, 2)
	if len(slice) == 1 {
		return slice[0], ""
	}

	return slice[0], slice[1]
}

func makeFlags(base string) (string, bool, bool) {
	namePtr := flag.String("n",
		"",
		"Bind current file path to alias",
	)
	addPtr := flag.Bool("a", false, "Bind current file path to base name")
	listPtr := flag.Bool("l", false, "Lists the currently installed shortcuts")
	flag.Parse()
	return *namePtr, *addPtr, *listPtr
}
