package main

import "fmt"

func main() {
	var brand string

	// prints the string quoted-form like this ""

	fmt.Printf("%q\n", brand)

	brand = "Google"
	fmt.Printf("%q\n", brand)
}
