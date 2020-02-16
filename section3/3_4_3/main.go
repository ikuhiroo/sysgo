// ソケット通信(アプリケーション層(HTTP)とトランスポート層(TCP / UDP)のみを利用する通信)
// HTTPやTCPのレベルで決められているルールに従って通信
// RFC (Request For Comments)
// RFC 793: TCP に関すること
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	fmt.Println("[Client 側]")
	fmt.Println("============================セッション成立============================")
	// net.Dial は net.Conn (通信のコネクションを表すインタフェース)を返す
	// io.Writer と io.Reader のハイブリッド
	// conn の実体は net.TCPConn 構造体のポインタ
	// サーバーの 80 番ポートに対してソケットを使ったプロセス間通信
	// "www.yahoo.co.jp:80"
	// "www.google.co.jp:80"
	conn, err := net.Dial("tcp", "www.google.co.jp:80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running at www.google.co.jp:80")

	fmt.Println("============================リクエストの作成============================")
	// HTTP/1.0 version
	// \n：Unix系OS全般、Mac OS X
	// \r\n：Windows系OS
	conn.Write([]byte("GET / HTTP/1.0\nHost: www.google.co.jp\n\n"))

	// HTTP のレスポンスをパースする
	// bufio.Reader でラップした net.Conn を渡す
	// bufio.NewReader() はHTTPのヘッダーやボディに分解できる
	// http.Response 構造体のオブジェクトが返される
	fmt.Println("============================レスポンスの作成============================")
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println("============================ヘッダーの表示============================")
	fmt.Println(res.Header)
	fmt.Println("============================ボディの表示============================")
	io.Copy(os.Stdout, res.Body)
}
