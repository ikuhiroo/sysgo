// 標準入力
// io.Reader メモリ確保 /  メモリを確保しない
package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		buffer := make([]byte, 5)
		// Read() にはタイムアウトの仕組みがない
		size, err := os.Stdin.Read(buffer)
		if err != nil {
			fmt.Print("EOF")
			break
		}
		fmt.Printf("size=%d input=%s \n", size, string(buffer))
	}
}
