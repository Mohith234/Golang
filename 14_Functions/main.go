package main

import "fmt"

func main() { // main is the entrypoint function
	fmt.Println("Welcome to Functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is : ", result)

	// proRes, _ := proAdder(2, 5, 6, 7)
	proRes, myMessage := proAdder(2, 5, 6, 7)
	fmt.Println("Pro Result is : ", proRes)
	fmt.Println("Pro message is : ", myMessage)
}

func adder(val1 int, val2 int) int { // function signatures - return type of the function
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro"
}

func greeter() {
	fmt.Println("Hello world")
}

// we cannot define a function directly inside another funciton
// we can just call it inside another function
