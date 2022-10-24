/*Create a new library
In it, create a function that returns the Go version
Create a command and import your library
Call your function that returns Go version
Run your program
It should print the current Go version on your system.*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
}
