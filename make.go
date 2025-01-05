package main

import "fmt"

/*
スライスの長さと容量って？
スライスの長さ(len): スライスに含まれている要素の数
スライスの容量(cap): スライスの基となる配列の、スライス開始位置から終端までの要素数。容量は「スライスの開始位置から、基となる配列の最後までの要素数」
*/

func printSlice(s []int) {
	s = []int{2, 3, 5, 7, 11, 13}

	// Slice the slice to give it zero length.
	s = s[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// Extend its length.
	s = s[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// Drop its first two values.
	s = s[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}

/*
make とは:

スライス、マップ、チャネルを「使える状態」にするための道具。
どんなときに使う？:

スライスやマップを作っただけでは使えないとき。
初心者向けアナロジー:

make は「箱の組み立てキット」。
スライスやマップは「箱の中に入れる準備」をしないとエラーになる。
*/

func making() {
	a := make([]int, 5)
	fmt.Println(a)
}

type classmate struct {
	fname string
	lname string
}

var rugbybu = []classmate{
	{fname: "ryosuke", lname: "sakamoto"},
	{fname: "ryoma", lname: "huruich"},
	{fname: "kunpei", lname: "oonishi"},
	{fname: "masataka", lname: "yazawa"},
}

func makefor() {
	for _, name := range rugbybu {
		fmt.Println(name.fname)
	}
}
