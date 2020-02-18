// タイマー
// time.After(duration)を使って，決まった時間を測るタイマーを作成
package main

import (
	"fmt"
	"time"
)

func main() {
	// バッファーなし
	// c1 := make(chan string)

	// select {
	// case msg1 := <-c1:
	// fmt.Println("Message 1", msg1)
	// case <-time.After(3 * time.Second):
	// fmt.Println("timeout")
	// }

	duration := 3 * time.Second
	<-time.After(duration)

	fmt.Printf("timeout")
}
