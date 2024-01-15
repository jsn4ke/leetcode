package leatcode

import (
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
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {

	})
}
