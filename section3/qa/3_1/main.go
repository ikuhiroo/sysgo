// ファイルのコピー
// コマンドラインオプション
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// パース処理
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	reader, _ := os.Open(args[0])
	defer reader.Close()
	writer, _ := os.Create(args[1])
	defer writer.Close()
	io.Copy(writer, reader)
}
