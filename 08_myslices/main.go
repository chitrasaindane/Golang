package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")

	// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
	// Arrays are fixed in size, but slices can grow and shrink.
	// declaration of array and slices : var arr [5]int, var slice []int

	var fruit = []string{"apple", "banana", "mango", "orange", "grapes"}
	fmt.Printf("Fruit: %T\n", fruit)

	fruit = append(fruit, "kiwi")
	fmt.Println(fruit)

	fruit = append(fruit[2:]) // (start index:- 2, end index: end of slice)
	fmt.Println(fruit)

	fruit = append(fruit[1:3]) // (start index:1, end index: 3) 3 is exclusive
	fmt.Println(fruit)

	fruit = append(fruit[:2]) // (start index:0, end index: 2) 2 is exclusive
	fmt.Println(fruit)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 234 // this will give error as index out of range
	highScores = append(highScores, 234, 676, 97986) // this will work
	fmt.Println(len(highScores))
	fmt.Println(highScores)

	sort.Ints(highScores) // sort the slice
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove value from slice by index
	var courses = []string{"Python", "Javascript", "Java", "Go", "C++", "Ruby"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
