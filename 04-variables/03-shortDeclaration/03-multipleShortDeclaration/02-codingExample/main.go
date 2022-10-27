package main

import "fmt"

func main() {
	name, lastname := "Ayrton", "Jofre"
	fmt.Println(name, lastname)
	age, mbirthday := 26, "October"
	fmt.Println(age, mbirthday)
	on, off := true, false
	fmt.Println(on, off)

	// there's no limit
	// however, more declarations that you declare
	// more unreadable it becomes...

	degree, ratio, heat := 4.33, 30.5, 20.
	fmt.Println(degree, ratio, heat)
	nFiles, valid, msg := 10, true, "hello"
	fmt.Println(nFiles, valid, msg)

}
