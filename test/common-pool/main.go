package main

import (
	"fmt"
	"github.com/jolestar/go-commons-pool"
	"log"
	"math"
	"reflect"
	"time"
)

type MyPoolObject struct {
	Name string
}

type MyObjectFactory struct {
}

func (f *MyObjectFactory) MakeObject() (*pool.PooledObject, error) {
	log.Println("MakeObject")
	return pool.NewPooledObject(&MyPoolObject{Name: "panlin"}), nil
}

func (f *MyObjectFactory) DestroyObject(object *pool.PooledObject) error {
	//do destroy
	log.Println("DestroyObject")
	return nil
}

func (f *MyObjectFactory) ValidateObject(object *pool.PooledObject) bool {
	//do validate
	log.Println("ValidateObject")
	return false
}

func (f *MyObjectFactory) ActivateObject(object *pool.PooledObject) error {
	//do activate
	log.Println("ActivateObject")
	return nil
}

func (f *MyObjectFactory) PassivateObject(object *pool.PooledObject) error {
	//do passivate
	log.Println("PassivateObject")
	return nil
}

func PoolInfo(pool *pool.ObjectPool) {
	fmt.Printf(" GetNumIdle: %d\n", pool.GetNumIdle())
	fmt.Printf(" GetNumActive: %d\n", pool.GetNumActive())
	fmt.Printf(" GetDestroyedCount: %d\n", pool.GetDestroyedCount())
	fmt.Printf(" GetDestroyedByBorrowValidationCount: %d\n", pool.GetDestroyedByBorrowValidationCount())
}

var config *pool.ObjectPoolConfig = &pool.ObjectPoolConfig{
	MaxTotal:                       10,
	MaxIdle:                        10,
	MinIdle:                        0,
	Lifo:                           true,
	MaxWaitMillis:                  int64(-1),
	MinEvictableIdleTimeMillis:     int64(1000 * 60 * 30),
	SoftMinEvictableIdleTimeMillis: int64(math.MaxInt64),
	NumTestsPerEvictionRun:         3,
	TestOnCreate:                   false,
	TestOnBorrow:                   false,
	TestOnReturn:                   false,
	TestWhileIdle:                  true,
	TimeBetweenEvictionRunsMillis:  int64(1), //int64(-1)
	BlockWhenExhausted:             true,
	EvictionPolicyName:             "github.com/jolestar/go-commons-pool/DefaultEvictionPolicy"}

func main() {
	opool := pool.NewObjectPool(new(MyObjectFactory), config)

	for i := 0; i < 10; i++ {
		err := opool.AddObject()
		if err != nil {
			log.Printf("addobject failed: %s\n", err)
		}
		fmt.Printf("i: %d\n", i)
	}

	obj, _ := opool.BorrowObject()

	fmt.Println(reflect.TypeOf(obj))
	fmt.Println(obj.(*MyPoolObject).Name)

	for i := 0; i < 5; i++ {
		PoolInfo(opool)
		time.Sleep(1)
	}

	opool.ReturnObject(obj)
}
