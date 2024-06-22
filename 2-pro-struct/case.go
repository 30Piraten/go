package main

import "fmt"

const packageUser string = "Mein Name ist Marthe "

func main() {
	localUser := "Ru"

	result := packageUser + localUser

	fmt.Printf("%s\n", result)
}
