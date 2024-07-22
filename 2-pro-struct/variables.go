package main

import (
	"fmt"
	"os"
)

func main() {

	// variable declaration
	var user_account string = "this-user"
	var id int = 23

	var people = ""   // defined
	var cousin string // defined

	data := "userkeys"
	var junge int64 = "harvey" // value not well-defined

	// multiple variable declaration
	var x, y, z int
	var j, k, l string
	r, s, t := 0, true, "user"
	var a, b, c = true, 12.4, "user"

	// := used for calls to functions
	f, err := os.Open(infile)
	if err != nil {
		return err
	}
	f.Close()

	// := must declare new variables
	// this will throw a compile error: no new variables
	// best bet, use an assignment operator instead
	f, err := os.Create(outfile)

	fmt.Printf("The user's account is:%s | %d\n", user_account, id)

}
