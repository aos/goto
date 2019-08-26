package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"
)

func main() {
	jumpDirMap := make(map[string]string)

	home, err := os.UserHomeDir()
	Check(err)

	// Create the ~/.gotorc file if it doesn't exist
	f, err := os.OpenFile(
		filepath.Join(home, ".gotorc"),
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		0644,
	)
	Check(err)
	defer f.Close()

	// Create a map from our .gotorc file
	// { "shortcut": "/path/to/directory/" }
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		j, dir := Split(scanner.Text(), ",")
		jumpDirMap[j] = dir
	}

	// Our flags and arguments
	cli := ParseCommandLine()

	if cli.Add {
		cwd, err := os.Getwd()
		Check(err)

		base := filepath.Base(cwd)
		shortcutName := ""
		if cli.Arg != "" {
			shortcutName = cli.Arg
		} else {
			shortcutName = base
		}

		_, err = f.WriteString(fmt.Sprintf("%v,%v/\n", shortcutName, cwd))
		fmt.Printf("Added shortcut: %s | %s\n", shortcutName, cwd)
		Check(err)
		return
	}

	if cli.List {
		fmt.Println("Currently installed shortcuts:")
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		for k, v := range jumpDirMap {
			fmt.Fprintf(w, "%s\t%s\n", k, v)
		}
		w.Flush()
		return
	}

	if cli.Init {
		PrintShellIntegration(Bash)
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
