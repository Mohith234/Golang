package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=react&paymentid=182hebd"

func main() {
	fmt.Println("Handling URL's in go")

	// parsing
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type is %T\n", qparams) // key value types

	fmt.Println(qparams["coursename"]) // react

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsofURL := &url.URL{ // we need to pass a reference
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcsss",
		RawPath: "user-hitesh",
	}

	anotherURL := partsofURL.String()
	fmt.Println(anotherURL)
}
