// タイマー待ち
// チャネルをイベントの通知で使う例
package main

import "fmt"

func main() {
	fmt.Println("start sub()")

	// 終了を受けるためのチャネル(バッファなし)
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		// 終了の通知
		done <- true
	}()
	// データの受け取り
	<-done
	fmt.Println("all tasks are finished")

	// チャネルは開けっぱなしでいい
	// ガーベッジコレクションに回収されるので
}
