package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Go")
	content := "This needs to go in a file - learncoeonline.com"

	file, err := os.Create("./myinfo.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./myinfo.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is: ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
