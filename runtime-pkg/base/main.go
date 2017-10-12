package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func display(num int) {

	for i := 1; i < num; i++ {
		fmt.Println(i)
	}

	wg.Done()
}

func main() {
	//
	cpunum := runtime.NumCPU()
	fmt.Printf("cpu-num: %d\n", cpunum)

	wg.Add(cpunum)

	//
	goroot := runtime.GOROOT()
	fmt.Printf("GOROOT: %s\n", goroot)

	//
	fmt.Printf("GOOS: %s\n", runtime.GOOS)

	// set max cpu
	runtime.GOMAXPROCS(cpunum)

	go display(10)
	go display(10)

	wg.Wait()
}
