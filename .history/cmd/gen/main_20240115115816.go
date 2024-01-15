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
	err = os.Mkdir(nw, 0755)
	if nil != err {
		panic(err)
	}
	f, err := os.Create(path.Join(nw, `run.go`))
	if nil != err {
		panic(err)
	}
	f.Close()
}
