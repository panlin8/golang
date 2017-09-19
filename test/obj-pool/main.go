package main

import (
	"fmt"
	"log"
	"reflect"
	_ "runtime"
	"sync"
)

func main() {
	var pool = &sync.Pool{New: func() interface{} { return "hello beijing" }}

	//into pool
	pool.Put("hello world")

	//call gc
	//runtime.GC()

	temp := pool.Get()

	t := reflect.TypeOf(temp)

	v := reflect.ValueOf(temp)

	fmt.Println(t)
	fmt.Println(v)

	//out pool
	log.Printf("%v\n", temp)

	// if pool is NULL, call New create obj
	log.Printf("%s\n", pool.Get())
}
