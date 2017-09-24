package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	endata, err := rlp.EncodeToBytes([]uint{10, 20})
	if err != nil {
		fmt.Printf("unuse err:%v\n", err)
		return
	}

	var dedata []uint
	err = rlp.DecodeBytes(endata, &dedata)

	fmt.Println(dedata)
}
