package main

import "fmt"

func main() {
	fmt.Println("Array study of Golang")

	var fruit [4]string

	fruit[0] = "Apple"
	fruit[1] = "Banana"
	fruit[3] = "Orange"

	fmt.Println("Fruit Array: ", fruit)
	fmt.Println("Fruit Array Length: ", len(fruit))

	var veg = [3]string{"Carrot", "Beans", "Potato"}
	fmt.Println("Veg Array: ", len(veg))

}
