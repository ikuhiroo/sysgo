// ストリーム総集
// 制約
// 使って良いのは io パッケージな内容 + 基本文法
// io.Pipe() を使う場合は，ブロッキングを防ぐために goroutine を使う必要がある
// 文字列リテラルを使用してはいけない
package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	// ここに io パッケージを使ったコードを書く
	validation := []string{"A", "S", "C", "I", "I"}
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(io.MultiReader(computer, system, programming))

	for i := 0; i < len(validation); i++ {
		if strings.Contains(buffer.String(), validation[i]) {
			stream = append(stream, strings.NewReader(validation[i]))
		}
	}
	// 「ASCII」の文字列を出力させる
	io.Copy(os.Stdout, stream)
}
