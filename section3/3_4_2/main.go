// インタフェース間のデータコピー
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// ファイルの読み込み
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File open error: ", err)
		return
	}
	defer file.Close()

	fmt.Println("標準出力==============================")
	// write: os.Stdout, read: file
	io.Copy(os.Stdout, file)

	// EOF を示しているので読み取り位置を変更する必要がある
	_, err = file.Seek(0, os.SEEK_SET)
	if err != nil {
		fmt.Println("File seek error: ", err)
		return
	}

	// ファイルの作成
	newfile, err := os.Create("copy.txt")
	if err != nil {
		fmt.Println("File open error: ", err)
		return
	}
	defer newfile.Close()

	// ファイルをコピー
	io.Copy(newfile, file)

	// EOF から 先頭に移動
	_, err = newfile.Seek(0, os.SEEK_SET)
	if err != nil {
		fmt.Println("File seek error: ", err)
		return
	}

	fmt.Println("標準出力==============================")
	io.Copy(os.Stdout, newfile)
}
