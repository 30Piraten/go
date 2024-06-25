package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "remove trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {

	// pointers are a key to the flag package
	// this uses a program's cmd line args
	// to set values of variables in a program
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
