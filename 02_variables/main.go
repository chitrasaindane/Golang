package main

import "fmt"

const LoginToken string = "asdjfalsdfjalsdf"

func main() {
	var username string = "Chitra"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isloggedin bool = true
	fmt.Println("Is logged in: ", isloggedin)
	fmt.Printf("Variable is of type: %T \n", isloggedin)

	var smallint uint8 = 255
	fmt.Println("Small int: ", smallint)
	fmt.Printf("Variable is of type: %T \n", smallint)

	var smallfloat float32 = 3.1499867857466
	fmt.Println("Small float: ", smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	var smallcomplex complex64 = 3.1499867857466 + 3.1499867857466i
	fmt.Println("Small complex: ", smallcomplex)
	fmt.Printf("Variable is of type: %T \n", smallcomplex)

	//default values and some aliases

	var defaultint int
	fmt.Println("Default int: ", defaultint)
	fmt.Printf("Variable is of type: %T \n", defaultint)

	var defaultstring string
	fmt.Println("Default string: ", defaultstring)
	fmt.Printf("Variable is of type: %T \n", defaultstring)

	//implicit type

	var website = "golang.org"
	fmt.Println("Website: ", website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style

	numberOfUsers := 3000000.0
	fmt.Println("Number of users: ", numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	fmt.Println("LoginToken :", LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
