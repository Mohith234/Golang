package main

import "fmt"

func main() {
	fmt.Println("Arrays is less used in Go")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList)) // gives actual length of the array not the number of elements present

	var vegList = [6]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg List is:", len(vegList))

}
