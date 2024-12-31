package main

import (
	"fmt"
	"math/cmplx"
)

/*
 */
var (
	ToBe   bool       = false
	MaxInt uint       = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func kata() {
	fmt.Printf("tyep %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("tyep %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("tyep %T Value: %v\n", z, z)
}
