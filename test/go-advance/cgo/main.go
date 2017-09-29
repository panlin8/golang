package main

import (
	"fmt"
	"golib/test/go-advance/cgo/add"
)

func main() {
	ret := add.Add(2, 5)
	fmt.Println("add : ", ret)
}
