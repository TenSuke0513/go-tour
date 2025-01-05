package main

import "fmt"

type User struct {
	name string
	age  int
	from string
}

func smk() {
	m := map[string]User{
		"sotooka": {name: "yutaro", age: 29, from: "jiyugaoka"},
		"yazawa":  {name: "masataka", age: 29, from: "kunitachi"},
	}
	fmt.Println(m)
}
