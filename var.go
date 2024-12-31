package main

import (
	"fmt"
)

/*
<varとconstの違い>
varは可変constは変更不可
Q: useStateは可変とは言わないのか？
A: var と const の違いにおける「可変か否か」というのは、「型が可変かどうか」ではなく、「参照そのものが変わるかどうか」 ということです。
ReactのuseState で const を使っている場合に「値が変わっているように見える」のは、この違いに由来しています。

*/

var i, j = 1, 2

func varvar() {
	var c, python, java = true, false, "no"
	fmt.Println(i, j, c, python, java)
}
