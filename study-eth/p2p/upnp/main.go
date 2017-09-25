package main

import (
	"fmt"
	elog "github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/nat"
	"log"
	"net/http"
	"os"
)

const (
	layer int = 4 //debug
)

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for key, value := range db {
		fmt.Printf("%s = %d\n", key, value)
	}
}

func main() {
	elog.Root().SetHandler(elog.LvlFilterHandler(elog.Lvl(layer), elog.StreamHandler(os.Stderr, elog.TerminalFormat(true))))

	var nat_mode nat.Interface = nat.Any()

	nat_quit := make(chan struct{})

	// close to end map
	defer close(nat_quit)

	go func() {
		fmt.Println("map start...")
		nat.Map(nat_mode, nat_quit, "tcp", 8888, 8888, "panlin-nat")
		fmt.Println("map end...")
	}()

	//http server
	db := database{"iphone": 6000, "xiaomi": 2000}
	http.HandleFunc("/list", db.list)

	fmt.Println("http start...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil)) //defalut httphandermux
}
