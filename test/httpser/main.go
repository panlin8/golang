package main

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for key, value := range db {
		fmt.Printf("%s = %d\n", key, value)
	}
}

func main() {
	db := database{"iphone": 6000, "xiaomi": 2000}
	http.HandleFunc("/list", db.list)

	log.Fatal(http.ListenAndServe("localhost:8080", nil)) //defalut httphandermux
}
