package main

import (
	"flag"
	"fmt"
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
	err = os.Mkdir(nw, 0755)
	if nil != err {
		panic(err)
	}
	f, err := os.Create(path.Join(nw, fmt.Sprintf("%v.go", *name)))
	if nil != err {
		panic(err)
	}
	f.WriteString(fmt.Sprintf("package leetcode_%v\n", *name))
	f.Close()
}
