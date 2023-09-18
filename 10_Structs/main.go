package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	// No inheritance in golang; No super or parent
	// What you see is code.

	mohith := User{"Mohith", "mohith@go.dev", true, 16}
	fmt.Println(mohith)
	fmt.Printf("Mohith details are: %+v\n", mohith) // Details for structs
	fmt.Printf("Name is: %v\n", mohith.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// first capital letter means it can be used anywhere
