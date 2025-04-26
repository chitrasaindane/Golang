package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"coursename"`     //json key name
	Price    int      `json:"price"`          //json key name
	Platform string   `json:"platform"`       //json key name
	Password string   `json:"-"`              //ignore this field in json
	Tags     []string `json:"tags,omitempty"` //omit empty field in json
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcocourse := []Course{
		{"ReactJS", 299, "learnCodeOnline.in", "password123", []string{"web-dev", "js"}},
		{"MERN", 199, "learnCodeOnline.in", "password123", []string{"web-dev", "js"}},
		{"Angular", 299, "learnCodeOnline.in", "password123", nil},
	}

	// package this data as JSON data
	finaljson, err := json.MarshalIndent(lcocourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finaljson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
{
"coursename": "ReactJS",
"price": 299,
"website": "learnCodeOnline.in",
"tags": ["web-dev", "js"]
}
 `)

	var lcocourse Course
	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Printf("json data is valid\n")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Printf("json data is not valid\n")
	}
	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is:%T \n", k, v, v)
	}
}
