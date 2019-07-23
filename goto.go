package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var jumpDirMap = make(map[string]string)

func main() {
	cwd, err := os.Getwd()
	Check(err)

	home, err := os.UserHomeDir()
	Check(err)

	base := filepath.Base(cwd)

	f, err := os.OpenFile(
		filepath.Join(home, ".gotorc"),
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		0644,
	)
	Check(err)
	defer f.Close()

	// Create a map of our .gotorc file
	// { "shortcut": "/path/to/directory/" }
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		j, dir := Split(scanner.Text(), ",")
		jumpDirMap[j] = dir
	}

	// Our flags and arguments
	cli := ParseCommandLine(base)

	if cli.Add {
		shortcutName := ""
		if cli.Arg != "" {
			shortcutName = cli.Arg
		} else {
			shortcutName = base
		}
		_, err := f.WriteString(fmt.Sprintf("%v,%v/\n", shortcutName, cwd))
		fmt.Printf("Added shortcut: %s | %s\n", shortcutName, cwd)
		Check(err)
		return
	}

	if cli.List {
		fmt.Println("Currently installed shortcuts:")
		for k, v := range jumpDirMap {
			fmt.Printf("%v - \t%v\n", k, v)
		}
		return
	}

	if cli.Init {
		text := PrintShellIntegration(Bash)
		fmt.Println(text)
		return
	}

	if cli.Arg != "" {
		if dir, ok := jumpDirMap[cli.Arg]; ok {
			fmt.Printf("%s", dir)
			return
		}
		fmt.Println("Shortcut not found.")
		return
	}

	flag.Usage()
	return
}
