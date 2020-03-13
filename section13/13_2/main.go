// goroutine との情報共有
package main

import (
	"fmt"
	"time"
)

func sub1(c int) {
	fmt.Println("shared by argument", c*c)
}

func main() {
	// 引数渡し
	go sub1(10)

	// クロージャー(環境をキャプチャする匿名関数)のキャプチャ渡し
	c := 20
	go func() {
		// 関数内にキャプチャされる値を隠れ引数として渡す
		fmt.Println("shared by argument", c*c)
	}()
	time.Sleep(time.Second)
}
