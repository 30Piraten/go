package main

import "fmt"

// universal block
const Bias bool = true

// var time string

func main() {
	// within this is the syntactic block

	if Bias {

		var time string = "the time is 12.45PM"
		fmt.Printf("%v", time)
	}

}
