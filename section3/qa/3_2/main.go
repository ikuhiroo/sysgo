package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	// バッファー準備
	buffer := make([]byte, 1)
	writer, _ := os.Create("log.txt")
	defer writer.Close()
	// ランダム文字列を書き込む
	io.CopyBuffer(writer, rand.Reader, buffer)
}
