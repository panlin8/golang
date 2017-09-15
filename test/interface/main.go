package main

import (
	"flag"
	"fmt"
	"time"
)

var timeout = flag.Duration("timeout", 1*time.Second, "timeout time")

func main() {
	v1 := make(chan int, 5)
	v2 := make(chan int, 5)

	v2 <- 10
	v2 <- 15

	v1 = v2

	fmt.Println("vim-go", <-v1, <-v1)

	flag.Parse()

	fmt.Println("vim-go", *timeout)

	v3 := make(chan int)

	go func(v chan int) {
		v <- 10

	}(v3)

	fmt.Println("end", <-v3)

}
