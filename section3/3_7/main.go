// io.Reader / io.Writer でストリーム(pipe)を自由に操る
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("io.MultiReader===================================")
	header := bytes.NewBufferString("-------Header--------\n")
	content := bytes.NewBufferString("Example MultiReader\n")
	footer := bytes.NewBufferString("-------footer--------\n")
	// io.Reader の全ての入力がつながっている
	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)

	fmt.Println("\nio.TeeReader===================================")
	var buffer bytes.Buffer
	reader = bytes.NewBufferString("io.TeeReader samples !!!")
	// 読み込まれた内容を別の io.Writer に書き出す
	teeReader := io.TeeReader(reader, &buffer)
	// 終端記号まで読み込む
	_, _ = ioutil.ReadAll(teeReader)
	// バッファに残っている
	fmt.Println(buffer.String())

	fmt.Println("\nio.Pipe===================================")
	// 同期的なデータのやりとりのみ

}
