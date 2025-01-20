package main

import "fmt"

func main() {
	fmt.Println("Methods in Go")

	chitra := User{"Chitra", "chitra@gmail.com", 25, true}

	fmt.Println("User Status:-", chitra.GetStatus())

	fmt.Println("User NewAge:-", chitra.NewAge())

}

// defining a struct
type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

// methods are functions that are associated with a specific type. They enable you to define behavior for types, similar to how methods work in object-oriented programming. A method has a receiver argument, which is a special parameter used to indicate the type the method is associated with.
func (u User) GetStatus() string {
	if u.Status {
		return "Active"
	}
	return "Inactive"
}

func (u User) NewAge() int {
	u.Age = u.Age + 1
	return u.Age
}
