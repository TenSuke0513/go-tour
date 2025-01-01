package main

import (
	"fmt"
)

/*
type~ってのは、オブジェクトの型を作るイメージ
*/

type Users struct {
	sex     string
	age     int
	feature string
	job     string
}

var ryosuke = Users{sex: "men", age: 28, feature: "rugby", job: "code"}

func ryoten() {
	fmt.Println(ryosuke.feature)
}

func rugbyman() {
	users := []Users{
		{sex: "女", age: 32, feature: "rugby", job: "pro"},
		{sex: "men", age: 30, feature: "rugby", job: ""},
		{sex: "women"},
	}
	for _, user := range users {
		//fmt.Println(i)
		fmt.Println(user.sex)
	}
}
