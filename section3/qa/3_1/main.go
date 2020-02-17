// ファイルのコピー
// コマンドラインオプション
// !go run main.go old.txt new.txt
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

	reader, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	writer, err := os.Create(args[1])
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	io.Copy(writer, reader)
}
