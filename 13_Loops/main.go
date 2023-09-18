package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d:=0; d< len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{ // Here i gives index not value
	// 	fmt.Println(days[i])
	// }

	// for each
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++ // Control goes to rogueValue = 6
			continue
			// break
		}
		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at Mohith")

}
