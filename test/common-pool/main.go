package main

import (
	"fmt"
	"github.com/jolestar/go-commons-pool"
	"reflect"
)

type MyPoolObject struct {
	Name string
}

type MyObjectFactory struct {
}

func (f *MyObjectFactory) MakeObject() (*pool.PooledObject, error) {
	fmt.Println("MakeObject")
	return pool.NewPooledObject(&MyPoolObject{Name: "panlin"}), nil
}

func (f *MyObjectFactory) DestroyObject(object *pool.PooledObject) error {
	//do destroy
	fmt.Println("DestroyObject")
	return nil
}

func (f *MyObjectFactory) ValidateObject(object *pool.PooledObject) bool {
	//do validate
	fmt.Println("ValidateObject")
	return true
}

func (f *MyObjectFactory) ActivateObject(object *pool.PooledObject) error {
	//do activate
	fmt.Println("ActivateObject")
	return nil
}

func (f *MyObjectFactory) PassivateObject(object *pool.PooledObject) error {
	//do passivate
	fmt.Println("PassivateObject")
	return nil
}

func main() {
	pool := pool.NewObjectPoolWithDefaultConfig(new(MyObjectFactory))
	obj, _ := pool.BorrowObject()

	fmt.Println(reflect.TypeOf(obj))
	fmt.Println(obj.(*MyPoolObject).Name)

	pool.ReturnObject(obj)
}
