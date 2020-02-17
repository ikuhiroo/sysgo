package main

import (
	"archive/zip"
	"os"
)

func main() {
	file, _ := os.Create("log.zip")

	// io.ZipWriter: zip ファイル書き込み用の構造体，io.Reader も io.Writer も満たさない
	// zip ファイルの読み書きに構造体そのものではなくインタフェースを使う
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// newfile.txt -> log.zip
	// io.Writer が返る
	zipWriter.Create("newfile.txt")

	// 実際のファイルではなく，文字列 strings.Reader を使って zip ファイルを作成する
	// ラップして渡す
}
