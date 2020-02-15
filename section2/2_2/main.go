package main

import (
	"fmt"
)

// インターフェース の定義
// インターフェース: 変数をひとまとめにする
type Talker interface {
	Talk()
}

// Talker インターフェースを満たす構造体
// フィールド(複数のデータを内部に持てる)
// Greeter を表現するのに必要なデータが定義
// クラスがなく，構造体しかない
// 構造体がメソッドを持てる
type Greeter struct {
	name string
}

// 構造体にメソッドを定義
// Greeter 構造体はインターフェースを満たす
// -> インターフェースで宣言された全てのメソッドがデータ型に対して定義される
// レシーバー (g Greeter)
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	//""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
	//  interface と struct の互換性のチェック
	//""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
	// 1. interface の型を持つ変数を宣言
	var talker Talker
	// 2. 構造体のインスタンスのポインタをインタフェース型の変数に代入する
	talker = &Greeter{"ikuhiroo"}
	// アクセス方法
	talker.Talk()
}
