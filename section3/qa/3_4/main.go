// zip ファイルを WebServer からダウンロード
package main

import (
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
	file, _ := os.Open("log.zip")
	defer file.Close()

	// アクセスしたら zip ダウンロード
	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
