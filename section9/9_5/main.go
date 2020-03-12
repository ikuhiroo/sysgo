package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// パス連結
	fmt.Printf("Temp File Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))

	// パス分割
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
	// 配列化
	fragments := strings.Split(os.Getenv("GOPATH"), string(filepath.Separator))
	fmt.Println(fragments)

	// 複数のパスからなる文字列を分解
	for _, path := range filepath.SplitList(os.Getenv("GOPATH")) {
		execpath := filepath.Join(path, "bin")
		result, err := os.Stat(execpath)
		if !os.IsNotExist(err) {
			fmt.Println(execpath)
			return
		}
		fmt.Println(result)
	}

	// パスをそのままクリーンにする
	fmt.Println(filepath.Clean("/Users/ikuhiro/Desktop/sysgo/../sysgo/section9/9_5/main.go"))
	// path/path.go
	// パスを絶対パスに整形
	abspath, _ := filepath.Abs("main.go")
	fmt.Println(abspath)
	// /usr/local/go/src/path/filepath/path_unix.go
	// パスを相対パスに整形
	relpath, _ := filepath.Rel("/Users/ikuhiro/Desktop/sysgo", "/Users/ikuhiro/Desktop/sysgo/section9/9_5/main.go")
	fmt.Println(relpath)
	// path/filepath/path.go
}
