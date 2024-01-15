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
		ori := path.Join(dir, k, `run.go`)
		to := path.Join(dir, k, `run.go.tmp`)
		func() {
			file, err := os.Open(ori)
			if nil != err {
				panic(err)
			}
			defer file.Close()

			tmp, err := os.Create(to)
			if nil != err {
				panic(err)
			}
			defer tmp.Close()
			scanner := bufio.NewScanner(file)
			lines := 0
			for scanner.Scan() {
				lines++
				line := scanner.Text()
				if 1 == lines {
					line = fmt.Sprintf("package leetcode_%v\n", k)
				}
				if _, err := tmp.WriteString(line); nil != err {
					panic(err)
				}
			}
			if err := scanner.Err(); nil != err {
				panic(err)
			}
		}()
		os.Rename(to, ori)
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
