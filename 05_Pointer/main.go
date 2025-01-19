package main

import "fmt"

func main() {
	fmt.Println("Pointer study of Golang")

	var ptr *int
	fmt.Println("Pointer Value: ", ptr)

	mynum := 10
	var pointr = &mynum
	fmt.Println("Pointer Val ( poitr): ", pointr)
	fmt.Println("Pointer  (*pointr): ", *pointr)

	*pointr = *pointr * 2
	fmt.Println("Pointer New Value: ", mynum)
}
