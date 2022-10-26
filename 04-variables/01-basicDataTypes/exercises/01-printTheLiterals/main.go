// ---------------------------------------------------------
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool constants
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	prints()
}

func prints() {
	fmt.Println(-100, -50, 0, 50, 100)
	fmt.Println(.50, 3.14)
	fmt.Println(true, false)
	fmt.Println("Ayrton")
	fmt.Println("Añanduba")
}
