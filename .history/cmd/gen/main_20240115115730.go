package main

import (
	"flag"
	"os"
	"path"
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
	nw := path.Join(dir, `../../`, *name)
	os.Mkdir(nw)
}
