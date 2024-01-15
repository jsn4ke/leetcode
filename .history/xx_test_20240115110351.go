package leatcode

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestXxx(t *testing.T) {
	dir, err := os.Getwd()

	if nil != err {
		panic(err)
	}
	dired := map[string]struct{}{}
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			fmt.Println(err)
			return err
		}
		if info.IsDir() {

		}
	})
}
