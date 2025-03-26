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
	js := libs.LoadJsonFromFile("timing.json")
	fmt.Println(js)

	fmt.Println(js["main_loop"])
	// fmt.Println(js["pre_input"])
}
