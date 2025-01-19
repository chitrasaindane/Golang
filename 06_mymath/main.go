package main

import (
	"crypto/rand" // For cryptographic random numbers
	"fmt"
	"math/big"
	mrand "math/rand"
	"time"
)

func main() {
	fmt.Println("Maths in golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4.5

	fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	// Random number using math/rand
	mrand.Seed(time.Now().UnixNano())
	fmt.Println(mrand.Intn(5) + 1)

	// Random number using crypto/rand
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Crypto random number:", myRandomNum)
}
