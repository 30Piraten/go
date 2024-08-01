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

//<--------------------->

// Prints two Fahrenheit to Celsius conversions

func FahToCel() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fahToCelsius(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fahToCelsius(boilingF))
}

func fahToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
