package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	//No inheritance ,no super, no parent
	chitra := User{"Chitra", "chitra@gmail.com", 21, true}
	fmt.Println(chitra)

	//Accessing a single field
	fmt.Println("User Name:", chitra.Name)
	fmt.Printf("Chitra Details: %+v\n", chitra)

	fmt.Printf("Name is %v and Email is %v\n", chitra.Name, chitra.Email)

}

// defining a struct
type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
