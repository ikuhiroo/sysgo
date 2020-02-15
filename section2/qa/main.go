// io.Writer は入出力に共通する仕様を満たすメソッドを持つ型
// ファイルの入出力関連，バッファ，インターネットなど
package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// JSON 化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	// http.ResponseWriter を指定して JSON エンコーダーで変換
	// os.Stdout に出力するとログが残る
	// io.MultiWriter
	// JSON の文字列変換
	// gzip 圧縮を行いながら圧縮前の出力を標準出力にも出す

}

func main() {
	file, err := os.Create("qa_output.txt")
	if err != nil {
		panic(err)
	}

	// Q2.1: フォーマットしてデータを io.Writer に書き出す
	fmt.Fprintf(file, "%d\n", time.Now())

	// Q2.2: csv 出力 encoding/csv Flush() メソッドが必須
	writer := csv.NewWriter(file)
	writer.Write([]string{"Alice", "20"})
	writer.Write([]string{"Bob", "21"})
	writer.Write([]string{"Carol", "22"})
	writer.Flush()

	// Q2.3: gzip された JSON 出力しながら，標準出力にログを出力
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
