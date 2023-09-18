package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	// No inheritance in golang; No super or parent
	// What you see is code.

	mohith := User{"Mohith", "mohith@go.dev", true, 1}
	fmt.Println(mohith)
	fmt.Printf("Mohith details are: %+v\n", mohith) //Details for structs
	fmt.Printf("Name is: %v\n", mohith.Name)
	mohith.GetStatus()
	mohith.NewMail()
	fmt.Printf("Mohith Email is: %v\n", mohith.Email) // Gives out older email not modified one because the email is passed as a copy to the NewMail() method
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int // not exportable
	// first capital letter means it can be used anywhere
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is : ", u.Email)
}

// Functions in classes known as methods - In programming
// Functions in structs known as methods - In Golang
