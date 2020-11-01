package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	printVersion = false
	version      = "0.0"
)

func init() {
	flag.BoolVar(&printVersion, "version", false,
		"Print the current version of this program and exit.")
}

func main() {

	flag.Parse()

	if printVersion {
		fmt.Println("golang_starter_repo v" + version)
		os.Exit(0)
	}

	fmt.Println("Hello, World!")
}
