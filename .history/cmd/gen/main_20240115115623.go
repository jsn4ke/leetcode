package main

import (
	"flag"
	"os"
)

var (
	name = flag.String("name", "", "")
)

func main() {
	flag.Parse()
	dir, err := os.Getwd()
	if nil != err {
		panic(err)
	}

}
