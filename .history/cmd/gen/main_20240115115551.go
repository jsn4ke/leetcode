package main

import "flag"

var (
	name = flag.String("name", "", "")
)

func main() {
	flag.Parse()
}
