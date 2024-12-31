package main

import (
	"fmt"
	"math"
)

/*
for 初期化文; 条件式; 後処理 {
    // 繰り返したい処理
}

_は無視。
*/

func loop() {
	numbers := []int{10, 20, 30, 40, 50}

	for _, value := range numbers {
		fmt.Printf(" Value: %d\n", value)
	}
}

func mapkey() {
	capitals := map[string]string{
		"Japan":   "Tokyo",
		"USA":     "Washington, D.C.",
		"Germany": "Berlin",
	}

	for key, value := range capitals {
		fmt.Printf("Country: %s, Capital: %s\n", key, value)
	}
}

func pow() {
	v := math.Pow(3, 2)
	fmt.Println(v)
}

func nabe() {
	for day := 1; day <= 31; day++ {
		if day%10 == 3 || day/10 == 3 {
			fmt.Println(day, "〜")
		}
	}
}
