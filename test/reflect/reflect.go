package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age  int
}

func (p person) getName() {
	fmt.Println(p.Name)
}

func (p person) getAge() {
	fmt.Println(p.Age)
}

type test1 interface {
	getName()
}

type test2 interface {
	getName()
	getAge()
}

func main() {
	var t test1 = person{Name: "panlin", Age: 25}
	t.getName()
	fmt.Println(reflect.TypeOf(t))

	//
	b := t.(person)
	b.getName()
	b.getAge()
	fmt.Println(reflect.TypeOf(b))

	//type assert one
	x := t.(test2)
	x.getName()
	x.getAge()
	fmt.Println(reflect.TypeOf(x))

	//fmt.Println(reflect.ValueOf(3))
}
