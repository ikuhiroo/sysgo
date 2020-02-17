// PNG ファイル(バイナリhォーマット)を分析する
package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	// エンディアン変換(ネットワークバイトオーダーに変換)
	// 変換結果を格納する変数のポインタ
	binary.Read(chunk, binary.BigEndian, &length)

	// 4 バイトのバッファを作成
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	// チャンクを格納する配列
	var chunks []io.Reader

	// 先頭の 8 バイトがシグニチャ（固定のバイト列）
	// 最初の 8 バイトをとばす
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		// io.Reader
		// エンディアン変換(ネットワークバイトオーダーに変換)
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		// io.SectionReader
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		// 次のチャンクの先頭に移動
		// チャンク名（4 バイト）+ データ長 + CRC（ 4 バイト）先に移動
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	fmt.Println(chunks)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
