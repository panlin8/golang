package main

import (
	"fmt"
	"golib/go-rpc/data" //$GOPATH/ path
	"net"
	"net/http"
	"net/rpc"
)

type Arith int

func (p *Arith) Add(args data.Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	arith := new(Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("create listen failed!")
		return
	}

	http.Serve(l, nil)
}
