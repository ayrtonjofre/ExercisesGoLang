// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//
//  Fix the code by using the conversion expression.
// a, b := 10, 5.5
// fmt.Println(a + b)
// EXPECTED OUTPUT
//  15.5
// ---------------------------------------------------------
package main

import "fmt"

func main() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}

// change the datatype of variable "a" to float64 declaring the type before the variabable in the println
