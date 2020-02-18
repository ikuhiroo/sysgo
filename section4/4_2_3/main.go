// 素数計算をチャネルで返す
package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	// バッファなしチャネル
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 10000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		// チャネルの終了
		// これがないとdeadlock
		close(result)
	}()
	return result
}

func main() {
	// チャネルの作成
	pn := primeNumber()

	// for 変数 := ranfe チャネルで受信
	// 値が来るたびに for ループが回る，個数が未定の動的配列
	// チャネルがcloseされるまで
	for n := range pn {
		fmt.Println(n)
	}
}
