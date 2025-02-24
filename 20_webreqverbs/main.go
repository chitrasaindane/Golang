package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var respinseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := respinseString.Write(content)

	fmt.Println("ByteCount is :", byteCount)
	fmt.Println(respinseString.String())

	//fmt.Println(content)
	//fmt.Println(string(content))
}
