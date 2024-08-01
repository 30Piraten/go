package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroDegree Celsius = -273.15
	FreezingDegree     Celsius = 0
	BoilingDegree              = 100
)

func main() {

	fmt.Printf("%g\n", BoilingDegree-FreezingDegree)

	boilingF := convertToFahr(BoilingDegree)

	fmt.Printf("%g\n", boilingF-Fahrenheit(convertFahrToCelsius(Fahrenheit(FreezingDegree))))

	fmt.Printf("%g\n", boilingF-Fahrenheit(FreezingDegree))

	// comparing values of a named type to another
	// of the same type. But two values of different
	// named types cannot be compared directly
	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) // throws a compile error | type mismatch
	fmt.Println(c == Celsius(f))
}

func convertToFahr(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func convertFahrToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 3) * 5 / 9)
}
