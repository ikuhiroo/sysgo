// CopyN (指定したサイズ分だけ書き出す) の実装
// io.CopyN(dest io.Writer, src io.Reader, length int)
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// io.Copy() と本章で紹介した構造体1つ使う
func CopyN(dst io.Writer, src io.Reader, length int64) (written int64, err error) {
	// バッファーの準備
	lReader := io.LimitReader(src, length)
	written, err = io.Copy(dst, lReader)

	if written == length {
		return length, nil
	}
	if written < length && err == nil {
		panic(err)
	}
	return
}

func main() {
	var length int64 = 5

	fmt.Println("=========================ファイルから読み込む場合")
	reader, _ := os.Open("log.txt")
	defer reader.Close()
	CopyN(os.Stdout, reader, length)

	fmt.Println("\n=========================文字列を読み込む場合")
	// reader の使い回しはできない？？
	r := strings.NewReader("Hello there ! This is Ikuhiroo !")
	CopyN(os.Stdout, r, length)
}
