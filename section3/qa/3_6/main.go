// ストリーム総集
// 制約
// 使って良いのは io パッケージな内容 + 基本文法
// io.Pipe() を使う場合は，ブロッキングを防ぐために goroutine を使う必要がある
// 文字列リテラルを使用してはいけない
package main

import (
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
	// io.SectionReader と io.LimitReader を使えば必要な文字を切り出す
	lReader_1 := io.LimitReader(computer, 1)
	lReader_2 := io.LimitReader(system, 1)
	lReader_3 := io.NewSectionReader(programming, 5, 1)
	lReader_4 := io.NewSectionReader(programming, 8, 1)
	lReader_5 := io.NewSectionReader(programming, 8, 1)

	stream = io.MultiReader(lReader_3, lReader_2, lReader_1, lReader_4, lReader_5)
	// 「ASCII」の文字列を出力させる
	io.Copy(os.Stdout, stream)
}
