package mymath

import (
	"fmt"
)

type ForTest struct {
	num int
}

func (this *ForTest) Loops() {
	for i := 0; i < 10000; i++ {
		this.num++
	}
}

func init() {
	fmt.Println("mymath : init")
}

func Add(x, y int) int {
	return x + y
}

func Max(x, y int) int {
	var max int = 0

	if x > y {
		max = x
	} else {
		max = y
	}

	return max
}
