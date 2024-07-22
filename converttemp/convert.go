package converttemp

// convert Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// convert Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// // convert Celsius to Kelvin
// func CToK(c Celsius) Kelvin {
// 	return Kelvin(c + 273)
// }

// // convert Fahrenheit to Kelvin
// func FToK(f Fahrenheit) Kelvin {
// 	return CToK(FToC(f))
// }
