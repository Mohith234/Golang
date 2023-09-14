package main

import "fmt"

func main() {
	fmt.Println("If else in Go")
	loginCount := 23
	var result string

	if loginCount < 10 { // here { should not go into the next line
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exact"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Special Syntax - Assigning num on the go
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {

	// }
}
