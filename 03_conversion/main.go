package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to Pizza App"
	fmt.Println(welcome)
	fmt.Print("Enter Rating between 1 to 5 : ")
	reader := bufio.NewReader(os.Stdin)

	rating, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating", rating)

	//COnversion of string to float
	fmt.Println("After conversion")
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Rating is %f\n", numRating+1)
		fmt.Printf("Type of rating is %T\n", numRating)
	}
}
