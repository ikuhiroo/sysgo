// バイナリ解析用の io.Reader 関連機能
// 実際には，文字列分割で io.SectionReader を使うことはまずない
package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// strings.Reader は strings.NewReader で作成
	// strings.Reader(io.Read を満たす) でラップ
	reader := strings.NewReader("Example of problem\n")
	// io.SectionReader(io.ReadAt を満たす) に渡す
	sectionReader := io.NewSectionReader(reader, 14, 7)
	// 標準出力
	io.Copy(os.Stdout, sectionReader)
}
