package main

import (
	"bufio" // packages
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for the pizza:")

	// comma ok syntax || comma err syntax
	// instead of try catch
	input, _ := reader.ReadString('\n') // read until new line

	// If err is present instead of _ we store it in err
	// input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T", input)
}
