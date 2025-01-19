package main

import "fmt"

func main() {
	fmt.Println("Loop in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for loop
	fmt.Println("For loop")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println()

	// for loop with range
	fmt.Println("For loop with range")
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}
	fmt.Println()

	rougheValue := 1
	for rougheValue < 10 {
		if rougheValue == 3 {
			rougheValue++
			goto lco

		}
		if rougheValue == 5 {
			rougheValue++
			continue
		}
		if rougheValue == 7 {
			break
		}
		fmt.Println(rougheValue)
		rougheValue++
	}

lco:
	fmt.Println("Jumping to LearnCodeOnline")

}
