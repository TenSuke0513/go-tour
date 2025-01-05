package main

/*
Goにおける main.go の扱い
main パッケージのルール

Goでは実行可能なプログラムは 1つの main パッケージ しか持てない。
main パッケージには必ず1つの main() 関数が必要。
同じディレクトリに複数の main.go ファイルを作ると、エラーが発生する。
<go build: conflicting main packages in main.go and other_main.go>

ディレクトリを分けて、それぞれmain.goを作るか、main.goと同じ階層に別ファイルで関数作って、それをmainで呼び出す的な感じ。
*/

import (
	"fmt"
)

/*
goには関数名とか変数名(識別子っていうらしい)の頭が大文字か小文字でスコープが変わる。

*/

func main() {
	fmt.Println("テストプリント")
	testPrint()

	fmt.Println("hikihen")
	a, b := swap("中田", "田中")
	fmt.Println(a, b)
	x, y := split(10)
	fmt.Println(x, y)

	varvar()
	loop()
	mapkey()
	fmt.Println(sqrt(5))
	kata()
	hen()
	henhen()
	pow()
	nabe()
	ryoten()
	rugbyman()
	arry()
	making()
	makefor()
	smk()
}

/*
vscodeショートかっと関連


調べるメモ
*/
