package main

import "fmt"

func main() {

	t := 1
	r := &t // r,of type *int, points to t
	*r = 12 // equals t = 12

	// pointers can be comparable to
	// each other: two pointers are equal
	// if and only if they point to the
	// same variable or both are nil

	var x, y int
	fmt.Println("Pointers:", &x == &x, &x == &y, &x == nil)

	// the zero value of a pointer of
	// any type is nil. This equates to
	// true, if r points to a variable
	if r != nil {
		fmt.Printf("%t\n", r != nil)
	}
	// g := &7 // you cant' do this!

	// fmt.Println(*r)
	fmt.Printf("%d\n", *r)
	fmt.Printf("%d\n", t)

	getResult()
}

// <----------------------------------->

var x = f()

func f() *int {
	y := 1
	return &y

}

func getResult() {

	fmt.Println(f() == f())
}
