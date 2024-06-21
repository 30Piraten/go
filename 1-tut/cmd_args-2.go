package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	// if the amount of data from the cmd
	// is too large, a simpler & more efficient
	// solution would be to use the Join function
	// from the 'string' package

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(len(s))

}
