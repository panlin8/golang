package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string `json:"name", bson:"name"`
	Age  int    `json:"age", bson:"age"`
}

func main() {
	//mongodb://127.0.0.1:27017 : local db
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//database: panlin, collection: col
	col := session.DB("panlin").C("col")

	// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)

	//insert
	err = col.Insert(&Person{Name: "liyuan", Age: 26})
	if err != nil {
		fmt.Println("[ERROR] insert liyuan is failed!")
		return
	}

	//find one
	result := Person{}
	err = col.Find(bson.M{"name": "liyuan"}).One(&result)
	if err != nil {
		fmt.Println("[ERROR] name: panlin is not exist!")
		return
	}

	//find Iter
	iter := col.Find(bson.M{"name": "panlin"}).Iter()
	for iter.Next(&result) {
		if err != nil {
			fmt.Println("[ERROR] name: panlin is not exist!")
			return
		}

		fmt.Println("panlin age:", result.Age)
	}

	//remove
	info, err := col.RemoveAll(bson.M{"name": "liyuan"})
	if err != nil {
		fmt.Println("[ERROR] remove all failed!")
		return
	}

	fmt.Printf("info: %#v\n", info)

	//update
	info, err = col.UpdateAll(bson.M{"name": "panxinwei"}, bson.M{"$set": bson.M{"name": "panlin"}})
	if err != nil {
		fmt.Println("[ERROR] update all failed!")
		return
	}

	fmt.Printf("info: %#v\n", info)

	info, err = col.RemoveAll(bson.M{"usr_name": "liyuan"})
	if err != nil {
		fmt.Println("[ERROR] remove all failed!")
		return
	}

	fmt.Printf("info: %#v\n", info)
}
