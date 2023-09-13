package main

import "fmt"

const LoginToken string = "adaskljdh" // as 'L' is capital => Public

func main() {
	var username string = "Mohith"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45553352
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	// float32 data type gives only 5 digits
	// float64 data type gives more decimal values

	// default values
	var anotherVariable int
	fmt.Println(anotherVariable)

	// implicit types
	var website = "www.google.com" // no specific data type
	fmt.Println(website)           // implicitly gives string data type

	// no var style
	numberofUser := 300000
	fmt.Println(numberofUser)

	// warlus operator can only be used inside of the method. Cannot be used outside of the method
}
