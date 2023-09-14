package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("List of All languages:", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB") // we can delete like this in slices too
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}
	for _, value := range languages {
		fmt.Printf("For Key v, value is %v\n", value)
	}

	// %v is used for value
}
