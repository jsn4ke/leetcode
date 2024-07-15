package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
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
	nw := path.Join(dir, `../../internal`, *name)
	err = os.MkdirAll(nw, 0755)
	if nil != err {
		panic(err)
	}
	arr := strings.Split(*name, "/")
	f, err := os.Create(path.Join(nw, fmt.Sprintf("%v.go", arr[len(arr)-1])))
	if nil != err {
		panic(err)
	}
	f.WriteString(fmt.Sprintf("package leetcode_%v\n", strings.Join(arr, "_")))
	f.Close()
}
