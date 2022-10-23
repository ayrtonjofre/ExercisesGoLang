//EXERCISE
//Print the documentation of runtime.NumCPU function in the command line
//Print also its source code using in the command line
//HINT
//You should use correct go doc tools

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
}

//Command Documentation: go doc runtime NumCPU
//Command Source Code: go doc -src runtime NumCPU
