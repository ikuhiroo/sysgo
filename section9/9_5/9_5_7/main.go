package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var imageSuffix = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
	".tiff": true,
	".eps":  true,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(`Find images Usage: %s [path to find]`, os.Args[0])
		return
	}
	root := os.Args[1]
	// 深さ優先探索
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				// _build という名前のトラバースをスキップする
				if info.Name() == "_build" {
					return filepath.SkipDir
				}
				return nil
			}
			// 拡張子を取り出す
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if imageSuffix[ext] {
				rel, err := filepath.Rel(root, path)
				if err != nil {
					return nil
				}
				fmt.Printf("%s\n", rel)
			}
			return nil
		})
	if err != nil {
		fmt.Println(1, err)
	}
}
