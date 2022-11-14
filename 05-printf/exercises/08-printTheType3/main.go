// ---------------------------------------------------------
// EXERCISE: Print the Type #3
//
//  Print the type and value of "hello" using fmt.Printf
//
// EXPECTED OUTPUT
// 	Type of hello is string
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	fmt.Printf("The type of %s is %[1]T\n", "hello")
}
