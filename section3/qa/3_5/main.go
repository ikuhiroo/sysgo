// CopyN (指定したサイズ分だけ書き出す) の実装
// io.CopyN(dest io.Writer, src io.Reader, length int)
package main

import (
	"io"
	"os"
)

// io.Copy() と本章で紹介した構造体1つ使う
func CopyN(dst io.Writer, src io.Reader, length int) (written int64, err error) {
	// バッファーの準備
	lReader := io.LimitReader(src, int64(length))
	return io.Copy(dst, lReader)
}

func main() {
	var length = 4

	reader, _ := os.Open("log.txt")
	writer, _ := os.Create("new.txt")
	CopyN(writer, reader, length)
}
