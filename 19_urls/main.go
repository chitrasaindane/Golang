package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://www.something.com:8041/MyService ?name=John&age=25"

func main() {
	fmt.Println("Welcome to handling URLs in Go")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of the query params is %T\n", qparams)

	fmt.Println(qparams["name"])

	for _, val := range qparams {
		fmt.Println("Paaram is: ", val)
	}

	partsofUrl := &url.URL{
		Scheme:   "http",
		Host:     "www.newurl.com",
		Path:     "/MyService",
		RawQuery: "name=John&age=25"}

	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)
}
