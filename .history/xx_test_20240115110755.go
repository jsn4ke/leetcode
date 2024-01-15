package leatcode

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestXxx(t *testing.T) {
	dir, err := os.Getwd()

	if nil != err {
		panic(err)
	}
	dired := map[string]struct{}{}
	regexp.Compile(`$\d+_[a-zA-Z]+\.go$`)
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			fmt.Println(err)
			return err
		}
		fmt.Println(path, info.Name())
		if reg
		if info.IsDir() {
			dired[info.Name()] = struct{}{}
		}
		return nil
	})
}
