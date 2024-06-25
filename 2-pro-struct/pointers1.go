package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "remove trailing newline")
var sep = flag.String("s", " ", "separator")

// pointers are a key to the flag package
// this uses a program's cmd line args
// to set values of variables in a program
func main() {

	// flag.Parse() is called before the
	// flags are used to update the flag
	// variables from their default values
	flag.Parse()

	// non-flag args, i.e flag.Args() are
	// available as a slice of strings
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

// go run pointers1.go some words are
// go run pointers1.go -s & some words are
// go run pointers1.go -n some words are
