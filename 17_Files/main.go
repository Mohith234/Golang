package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Go")
	content := "This needs to go in a file - mohithgadireddy"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length is :", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err) // Avoid more code

	fmt.Println("Txt data inside the file is: ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
