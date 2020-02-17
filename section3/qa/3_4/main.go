// zip ファイルを WebServer からダウンロード
package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Content-Type ヘッダーを使ってファイルの種類が zip ファイルであることをブラウザに教える
	// 必須ではないが，ファイル名も指定できる
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	// WebServer で zip ファイルを作成
	source := map[string]string{
		"Hello": "World",
	}
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	encoder := json.NewEncoder(io.MultiWriter(gzipWriter, os.Stdout))
	encoder.SetIndent("", "   ")
	encoder.Encode(source)

	// アクセスしたら zip ダウンロード
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
