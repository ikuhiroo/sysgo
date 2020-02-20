// TCP ソケットを使った HTTP クライアント
package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	// HTTP/1.0 200 OK
	// Connection: close

	// Hello World

	// 接続確立
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(
		"GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}

	request.Write(conn)
	// RFC に従ったパース処理
	response, err := http.ReadResponse(
		bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
