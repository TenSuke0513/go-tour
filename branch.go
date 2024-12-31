package main

import "fmt"

/*
if switchの使い分けの目安
if を使うべき場合
条件式が複雑で、論理演算（&&, ||）を使う場合。
動的な値や関数の結果を条件に使用したい場合。
switch を使うべき場合
単一の値に基づく分岐が多い場合（特にケースが多い場合）。
条件が明確で可読性を重視したい場合。
デフォルトの分岐（default）が必要な場合。
*/

func iiii() {
	x := 10

	if x < 5 {
		fmt.Println("x は 5 未満です")
	} else if x >= 5 && x < 10 {
		fmt.Println("x は 5 以上 10 未満です")
	} else {
		fmt.Println("x は 10 以上です")
	}
}

func ssss() {
	x := 10

	switch {
	case x < 5:
		fmt.Println("x は 5 未満です")
	case x >= 5 && x < 10:
		fmt.Println("x は 5 以上 10 未満です")
	default:
		fmt.Println("x は 10 以上です")
	}
}
