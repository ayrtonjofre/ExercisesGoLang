// ---------------------------------------------------------
// EXERCISE: Multiple Short Declare
//
//  Declare two variables using multiple short declaration
//
// EXPECTED OUTPUT
//  14 true
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	age, live := 14, true
	fmt.Println(age, live)
}
