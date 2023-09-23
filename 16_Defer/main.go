package main

import "fmt"

func main() {
	defer fmt.Println("World") // delay the execution of a function
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// Multiple defer statements follow LIFO - Reverse order

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
