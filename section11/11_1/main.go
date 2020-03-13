package main

import (
	"fmt"
	"os"
)

func main() {
	path, _ := os.Executable()
	fmt.Printf(" 実行ファイル名: %s\n", os.Args[0])
	fmt.Printf(" 実行ファイルパス: %s\n", path)
	fmt.Printf(" プロセスID: %d\n", os.Getpid())
	fmt.Printf(" 親プロセスID: %d\n", os.Getppid())
}
