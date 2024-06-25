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
	incrCall()
}

// it is safe for a function to return
// the address of a local variable

var x = f()

func f() *int {
	y := 1
	return &y

}

func getResult() {

	fmt.Println(f() == f())
}

// pointers contain addresses, hence
// passing a pointer argument to a
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
