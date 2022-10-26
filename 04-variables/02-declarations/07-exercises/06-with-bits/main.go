// ---------------------------------------------------------
// EXERCISE: Declare with bits
//
//  1. Declare a few variables using the following types
//    int
//    int8
//    int16
//    int32
//    int64
//    float32
//    float64
//    complex64
//    complex128
//    bool
//    string
//    rune
//    byte
//
// 2. Observe their output
//
// 3. After you've done, check out the solution
//    and read the comments there
//
// EXPECTED OUTPUT
//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
//  ""
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	var (
		age         int
		height      int8
		brightness  int16
		contrast    int32
		collor      int64
		temperature float32
		sensitive   float64
		distance    complex64
		large       complex128
		isOff       bool
		phrase      string
		whatever    rune
		data        byte
	)

	fmt.Println(age, height, brightness, contrast, collor, temperature, sensitive, distance, large, isOff, phrase, whatever, data)
}
