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
	"math/rand"
)

/*
パッケージのインポート

*/

func main() {
	fmt.Println("my favorite number is", rand.Intn(10))
}