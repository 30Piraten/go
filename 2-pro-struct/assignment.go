package main

import "fmt"

func main() {

	// get the results for both func
	resultGcd := gcd(19, 16)
	resultFib := fib(10)

	// print result to console
	fmt.Println("The GCD is:", resultGcd)
	fmt.Println("The 5th number of the fib is:", resultFib)
}

// Greatest common divisor of 2 integers
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// The nth Fibonacci number
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
