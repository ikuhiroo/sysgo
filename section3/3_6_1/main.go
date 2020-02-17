// テキスト解析
// テキスト構文解析器
// バイト列か文字列
// ReadString('\n'): 戻り値に改行コードを含む
// Scan: 戻り値に改行コードを含まない / 分割文字が削除
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `Hello world
My name is Ikuhiroo.
I love coding.`

var source1 = `123 1.234 1.0e4 test`

func main() {
	fmt.Println("bufio.Reader を用いる方法==========================")
	// io.Reader インタフェースを満たすstrings.reader でラップ
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%v\n", line)
		if err == io.EOF {
			break
		}
	}

	fmt.Println("\nbufio.Scanner を用いる方法==========================")
	// デフォルトは改行区切り
	scanner := bufio.NewScanner(strings.NewReader(source))
	// 単語区切りに変更
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("%v\n", scanner.Text())
	}

	fmt.Println("\nデータ型を指定して解析するを用いる方法==========================")
	reader1 := strings.NewReader(source1)

	var i int
	var f, g float64
	var s string

	// io.Reader のデータを整数や浮動小数点数に変換
	// １つ目の引数にio.Readerを渡し，それ以降に変数のポインタを渡すと，その変数にデータが書き込まれる
	// スペース区切りが前提
	fmt.Fscan(reader1, &i, &f, &g, &s)
	fmt.Printf("i=%v f=%v g=%v s=%v\n", i, f, g, s)
}
