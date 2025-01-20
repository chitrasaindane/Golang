package main

import "fmt"

func greeter() {
	fmt.Println("Hello from Golang")
}

func add(val1 int, val2 int) int {
	return val1 + val2

}

func proadder(val ...int) (int, string) {
	total := 0
	for _, v := range val {
		total += v
	}
	return total, "Success"
}
func main() {
	fmt.Println("Functions in Go")
	greeter()

	result := add(10, 20)
	fmt.Println("Result:", result)

	prores, msg := proadder(10, 20, 30, 40, 50)
	fmt.Println("Proadder:", prores)
	fmt.Println("Message:", msg)
}
