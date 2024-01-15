package leatcode

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	dir, err := os.Getwd()

	if nil != err {
		panic(err)
	}
	dirs := getFolder()
	files := getFiles()

}

func getFolder() map[string]struct{} {
	dir, err := os.Getwd()

	if nil != err {
		panic(err)
	}
	dired := map[string]struct{}{}
	exp, err := regexp.Compile(`^\d+_[a-zA-Z]+$`)
	if nil != err {
		panic(err)
	}
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			fmt.Println(err)
			return err
		}
		if !exp.MatchString(info.Name()) {
			return nil
		}
		if info.IsDir() {
			dired[info.Name()] = struct{}{}
		}
		return nil
	})
	return dired
}

func getFiles() map[string]fs.FileInfo {
	dir, err := os.Getwd()

	if nil != err {
		panic(err)
	}
	mp := map[string]fs.FileInfo{}
	exp, err := regexp.Compile(`^\d+_[a-zA-Z]+\.go$`)
	if nil != err {
		panic(err)
	}
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if nil != err {
			fmt.Println(err)
			return err
		}
		if !exp.MatchString(info.Name()) {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		mp[strings.Split(info.Name(), `.`)[0]] = info
		return nil
	})
	return mp
}
