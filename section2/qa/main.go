// io.Writer は入出力に共通する仕様を満たすメソッドを持つ型
// https://www.write-ahead-log.net/entry/2017/08/06/025738
package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 必要なHTTPヘッダを送信する
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	// JSON 化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	// NewEncoder / Marshall / ...
	// io.MultiWriter: gzipWriter を受け取り，書き込まれたデータを加工して，別の io.Writer(標準出力) に書き出す
	encoder := json.NewEncoder(io.MultiWriter(gzipWriter, os.Stdout))
	encoder.SetIndent("", "   ")
	encoder.Encode(source)

	// 圧縮せずに通信する
	// bytes, _ := json.Marshal(source)
	// io.WriteString(w, string(bytes))
}

func main() {
	// // Q2.1: フォーマットしてデータを io.Writer に書き出す
	// fmt.Fprintf(file, "%d\n", time.Now())

	// // Q2.2: csv 出力 encoding/csv Flush() メソッドが必須
	// writer := csv.NewWriter(file)
	// writer.Write([]string{"Alice", "20"})
	// writer.Write([]string{"Bob", "21"})
	// writer.Write([]string{"Carol", "22"})
	// writer.Flush()

	// Q2.3: gzip された JSON 出力しながら，標準出力にログを出力
	// 通信に使用するデータ量を制限する
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
