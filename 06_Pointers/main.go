package main

import "fmt"

func main() {
	fmt.Println("Hi")

	// var ptr *int
	// fmt.Println("Value of pointer", ptr) // default value is nil

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of pointer", ptr)
	fmt.Println("Pointer value is", *ptr)
	// Pointer is directe reference to the memory location
	// *ptr => value stored in the pointer

	*ptr = *ptr + 2
	fmt.Println("New Value is", myNumber)
}

// New Keyword

//
