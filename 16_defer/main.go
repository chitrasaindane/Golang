package main

import "fmt"

func main() {
	defer println("deferred call 1")
	println("main function")
	defer println("deferred call 2")
	defer println("deferred call 3")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
