package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate our pizza")

	reader := bufio.NewReader(os.Stdin) // reference of input
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for reading", input)

	// numRating, err := strconv.ParseFloat(input, 64) // '\n' comes
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // Trims the spaces

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating", numRating+1)
	}
}
