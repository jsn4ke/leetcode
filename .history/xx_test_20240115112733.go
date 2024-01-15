package leatcode

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path"
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
	for k := range dirs {
		k := k
		func() {
			file, err := os.Open(path.Join(dir, k, `run.go`))
			if nil != err {
				panic(err)
			}
			defer file.Close()
			tmp, err := os.Open(path.Join(dir, k, `run.go.tmp`))
			if nil != err {
				panic(err)
			}
			defer tmp.Close()
			// 创建一个带缓冲的读取器
			reader := bufio.NewReader(file)

			// 写入新的第一行内容
			_, err = tmp.WriteString(fmt.Sprintf("package leetcode_%v", k) + "\n")
			if err != nil {
				panic(err)
			}
			// 读取并写入剩余的内容
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				_, err = tempFile.WriteString(line)
				if err != nil {
					return err
				}
			}
		}()

	}
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
