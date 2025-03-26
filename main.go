package main

import (
	"flag"
	"libs"
)

var (
	script_name string
)

func main() {
	flag.StringVar(&script_name, "script", "sample.json", "The script to execute")
	flag.Parse()
	libs.Script_path = script_name
	libs.Main()
}
