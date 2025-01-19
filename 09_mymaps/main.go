package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	languages["go"] = "golang"

	fmt.Println("List of all Languages: ", languages)
	fmt.Println("Length of the map: ", len(languages))

	fmt.Println("Js short for:", languages["js"])

	delete(languages, "rb")
	fmt.Println("List of all Languages: ", languages)

	//loops in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("Value is %v\n", value)
	}
}
