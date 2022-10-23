package main

// File Scope
import "fmt"

// File Scope
const ok = true

// Package Scope
func main() { //Comeca o bloco Scope

	hello := "Hello"
	fmt.Println(hello, ok)
} // Block scope ends
