package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in Go")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Vlaue of dice is", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, You can Open")
	case 2:
		fmt.Println("You can Move 2 spots")
	case 3:
		fmt.Println("You can Move 3 spots")
	case 4:
		fmt.Println("You can Move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can Move 5 spots")
	case 6:
		fmt.Println("You can Move 6 spots and run again")
	default:
		fmt.Println("What was that?")
	}
}
