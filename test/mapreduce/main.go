package main

import "fmt"

func subMap(c chan int, i int) {
	c <- i
}

func main() {
	var n int = 8
	c := make(chan int, n)

	//map
	for i := 0; i < n; i++ {
		go subMap(c, i)
	}

	//reduce
	var count int = 0
	for i := 0; i < n; i++ {
		count += <-c
	}

	fmt.Println(count)
}
