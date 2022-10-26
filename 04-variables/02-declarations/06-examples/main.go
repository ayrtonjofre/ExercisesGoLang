package main

import "fmt"

func main() {
	//Nomes de variaveis sao sensitiveCase:
	//Myage, MyAge, MYAGE, sao variaveis diferentes

	//Caso de uso:

	//Quando usar um Parallel Declaration?

	//NOT-GOOD:

	//var myAge int
	//var yourAge int

	//SO-SO

	//var (
	//	myAge int
	//	yourAge int
	//)

	//BETTER:

	var myAge, yourAge int
	fmt.Println(myAge, yourAge)

	var temperature float64
	fmt.Println(temperature)

	var succes bool
	fmt.Println(succes)

	var language string
	fmt.Println(language)
}
