// テスト用の適当なサイズのファイルを作成
package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, _ := os.Create("rand.txt")
	defer file.Close()
	io.CopyN(file, rand.Reader, 1024)

	// buffer := make([]byte, 1024)
	// writer, _ := os.Create("log.txt")
	// defer writer.Close()
	// io.CopyBuffer(writer, rand.Reader, buffer)
}
