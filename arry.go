package main

import "fmt"

func arry() {
	var a [2]string
	a[0] = "can"
	a[1] = "tama"
	fmt.Println(a)

	hour := []int{1, 2, 3, 4, 5}
	hour = append(hour, 6)
	fmt.Println(hour)

	ss := hour[1:4]
	fmt.Println(ss)

	type kugayama struct {
		fname string
		lname string
	}

	qkumi := []kugayama{
		{fname: "koyama", lname: "masato"},
		{fname: "yamamoto", lname: "daiki"},
	}

	qkumi = append(qkumi, kugayama{fname: "huruichi", lname: "ryoma"})
	qkumi[0] = kugayama{fname: "kashiwagi", lname: "shugo"}
	fmt.Println(qkumi)

	igobu := []struct {
		age  int
		name string
	}{
		{18, "tanaka"},
		{19, "yamashita"},
	}
	fmt.Println(igobu)
}
