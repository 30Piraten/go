package main

import (
	"fmt"
)

// var x, or x points to &y
var x = f()

func main() {
	getResult()
	incrCall()
}

// y will still exist even after
// the call to f has returned, and
// the pointer x (x = f()), will still
// refer to it.
func f() *int {
	y := 1
	return &y
}

func getResult() {

	fmt.Println(f() == f())
}

// pointers contain addresses, hence
// passing a pointer argument to a function
// enables the function to update the
// variable that was indirectly passed
func increment(p *int) int {
	*p++ // increments what p points to, not change it
	return *p
}

func incrCall() {
	m := 3
	// result := increment(&m)
	increment(&m)              // side effect m = 4
	fmt.Println(increment(&m)) // 5 and m = 5
}
