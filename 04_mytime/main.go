package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("TIME study of Golang")

	presentTime := time.Now()
	fmt.Println("Present Time: ", presentTime)

	fmt.Println("Present Date & Time: ", presentTime.Format("2006-01-02 15:04:05"))

	fmt.Println("Present Date & Day: ", presentTime.Format("2006-01-02 Mon"))

	createDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Create Date: ", createDate)
	fmt.Println("Create Date & Time: ", createDate.Format("2006-01-02 Mon"))
}
