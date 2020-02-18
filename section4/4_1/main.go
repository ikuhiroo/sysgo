// go をつけるだけで並列処理
// FIFO 型
package main

import (
	"fmt"
	"time"
)

// goroutine が呼ぶ関数
func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}
func main() {
	fmt.Println("start sub()")
	// goroutine を使って関数を実行
	// go sub()
	// time.Sleep(2 * time.Second)

	// インラインで無名関数を作ってその場で goroutine で実行
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
}
