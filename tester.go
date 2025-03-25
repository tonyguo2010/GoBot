package main

import (
	"fmt"
	"libs"
)

func main() {
	// s := "Welcome,to,GeeksforGeeks"
	// fmt.Println("", s)

	// result := strings.SplitAfter(s, ",")
	// fmt.Println("Result:", result[1])
	operations := libs.LoadOperations("PCL.json")

	fmt.Println(operations)
}
