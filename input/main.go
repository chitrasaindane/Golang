package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "User input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Rating: ")

	//comma ok || err Syntax

	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", rating)
	fmt.Printf("Type of rating is %T\n", rating)
}
