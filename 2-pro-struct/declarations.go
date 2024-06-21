// Four major Declarations:
// var, const, func, type

package main

import "fmt"

const boilingF = 122.5

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point is = %gF or %gC\n", f, c)
}
