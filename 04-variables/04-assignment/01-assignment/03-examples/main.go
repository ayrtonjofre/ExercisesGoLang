package main

import "fmt"

func main() {
	var (
		name   string
		age    int
		famous bool
	)

	//Example1
	name = "Newton"
	age = 80
	famous = true
	fmt.Println(name, age, famous)

	//Example2

	name = "Somebody"
	age = 20
	famous = false
	fmt.Println(name, age, famous)

	//Example3
	// EXERCISE: Why this doesn't work? Think about it.
	// save the previous value of the variable
	// to a new variable

	var prevName string
	prevName = name

	name = "Einstein"

	fmt.Println("Previous Name: ", prevName)
	fmt.Println("Current Name: ", name)

}
