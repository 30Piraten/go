package main

import (
	"fmt"
	"unsafe"
)

// the new function, is only a
// syntactic convenience
func main() {

	// each call returns a distinct variable
	// with a unique address
	p := new(int)
	r := new(int)

	// there is an exception to this
	// with zero-sized variables whose
	// type bear no information such as a
	// struct{} or [0]int.
	// --> this happens because they don't require
	// any storage space, so the compiler con optimize
	// their memory usage by giving them the same address

	// define two zero-sized variables
	var a struct{}
	var b struct{}

	// check if both variables
	// have the same address
	fmt.Println("a:", unsafe.Pointer(&a))
	fmt.Println("b:", unsafe.Pointer(&b))
	fmt.Println("result:", unsafe.Pointer(&a) == unsafe.Pointer(&b))

	// define two zero-sized variables
	var x [0]int
	var y [0]int

	// check if both variables
	// have the same address
	fmt.Println("x:", unsafe.Pointer(&x))
	fmt.Println("y:", unsafe.Pointer(&y))
	fmt.Println("result:", unsafe.Pointer(&x) == unsafe.Pointer(&y))

	fmt.Println(p == r)
}
