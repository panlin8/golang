package main

import (
	"fmt"
	"golib/go-rpc/data"
	"net/rpc"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: ./client 127.0.0.1")
	}

	seraddr := args[1]

	client, err := rpc.DialHTTP("tcp", seraddr+":1234")
	if err != nil {
		fmt.Println("DialHTTP create failed!")
		return
	}

	param := data.Args{5, 8}
	var reply int

	err = client.Call("Arith.Add", param, &reply)
	if err != nil {
		fmt.Println("rpc call failed! -> ", err)
		return
	}

	fmt.Printf("reply = %d\n", reply)
}
