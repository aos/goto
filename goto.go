package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var addFlag string

const usage = "Bind current file path to name, defaults to basename"

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	base := filepath.Base(cwd)

	fmt.Println(filepath.Base(cwd))
	flag.StringVar(&addFlag, "add", base, usage)
	flag.StringVar(&addFlag, "a", base, usage+" (shorthand)")
}
