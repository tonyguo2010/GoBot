package main

import (
	"flag"
	"libs"
)

var (
	start_point string
)

func main() {
	flag.StringVar(&start_point, "url", "https://banner.centennialcollege.ca", "start URL")
	flag.Parse()
	libs.Main(start_point)
}
