package main

/*
:=（短い変数宣言）:

新しい変数の宣言と初期化を同時に行う。
関数内でのみ使用可能。
型推論が行われる。
=（代入）:

既に宣言された変数に新しい値を代入する。
関数内外で使用可能。

*/

import (
	"fmt"
)

func hen() {
	var saka int = 1
	fmt.Println(saka)
}

func henhen() {
	moto := 2
	fmt.Println(moto)
}
