package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

type Btree struct {
	Data   interface{}
	Lchild *Btree
	Rchild *Btree
	Dtype  reflect.Type
}

func NewBtree(dtype reflect.Type) *Btree {
	tree := Btree{Dtype: dtype}
	return &tree
}

func (tree *Btree) Insert(data interface{}) {
	if tree.Dtype != reflect.TypeOf(data) {
		fmt.Println("insert data type is error!")
		return
	}

	//node := &Btree{Data: data, Dtype: tree.Dtype}

	if tree.Lchild != nil {

	} else {

	}
}

func getSha256Code(s string) {
	hash := sha256.New()
	hash.Write([]byte(s))
	fmt.Printf("%x\n", hash.Sum(nil))
}

func main() {
	getSha256Code("panlin")
}
