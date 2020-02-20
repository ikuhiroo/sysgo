// TCP ソケットを使った HTTP サーバー
// TCP の機能（net.Conn）
// 低レベルの機能(net/http 以下を用いない)を使ってHTTPサーバーを書く
// HTTP/1.0 相当の送受信
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	// server is running at localhost:8888
	// Accept &{[%!V(uint8=127) %!V(uint8=0) %!V(uint8=0) %!V(uint8=1)] %!V(int=52179) %!V(string=)}
	// GET / HTTP/1.1
	// Host: localhost:8888
	// Accept:
	// User-Agent: curl/7.64.1"

	// ソケットを開いて待ち続ける
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("server is running at localhost:8888")
	// 一度で終了しないために Accept() を何度も呼び出す
	for {
		// 接続確立
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// 1 リクエスト処理中に他のリクエストの Accept() が行えるように
		// Goroutine を使って非同期にレスポンスを処理する
		go func() {
			// conn を使った読み書き
			fmt.Printf("Accept %V\n", conn.RemoteAddr())

			// リクエストを読み込む
			// HTTP リクエストのヘッダー、メソッド、パスなどの情報を切り出す
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			// デバッグ用: 読み込んだリクエストを取り出す
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))

			// レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}

	fmt.Println("vim-go")
}
