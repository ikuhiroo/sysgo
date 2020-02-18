// OS からのシグナルをチャネルで受け取る例
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// サイズが 1 より大きいチャネルを作成
	signals := make(chan os.Signal, 1)
	// SIGINT (Ctrl+C) を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルが来るまで待つ
	fmt.Println("Waiting SIGINT (Ctrl+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}
