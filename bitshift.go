package main

import "fmt"

/*
ビットフラグを使う時と普通に数字を割り当てることの違い
普通に数字を割り当てる場合:

単純な状態管理が必要なとき。
状態が非常に多くないとき。
シンプルさを重視したい場合。
ビットフラグを使う場合:

状態が複数あり、同時に複数の状態を持つ必要がある場合。
高速な検索や効率的なデータ管理が求められる場合。
状態が増える可能性がある場合。
*/

// フラグを定義
const (
	Read  = 1 << 0 // 0001 = 1 (読み取り権限)
	Write = 1 << 1 // 0010 = 2 (書き込み権限)
	Exec  = 1 << 2 // 0100 = 4 (実行権限)
)

func bitshift() {
	// 権限を持つユーザーの例
	var userPermissions int = 0 // 初期状態は権限なし

	// 権限を追加
	userPermissions |= Read                    // 読み取り権限を追加
	userPermissions |= Write                   // 書き込み権限を追加
	fmt.Printf("現在の権限: %b\n", userPermissions) // 出力: 11 (2進数)

	// 権限の確認
	if userPermissions&Exec != 0 {
		fmt.Println("実行権限があります")
	} else {
		fmt.Println("実行権限がありません")
	}

	// 権限の削除
	userPermissions &^= Write                  // 書き込み権限を削除
	fmt.Printf("現在の権限: %b\n", userPermissions) // 出力: 1 (2進数)

	// 権限の追加と確認
	userPermissions |= Exec
	if userPermissions&Exec != 0 {
		fmt.Println("実行権限が追加されました")
	}
	fmt.Printf("最終的な権限: %b\n", userPermissions) // 出力: 101 (2進数)
}
