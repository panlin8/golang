package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func PrintGoInfo() {
	fmt.Println("go version: ", runtime.Version())
}

func PrintCpuInfo() {
	//
	fmt.Printf("CPU NUM: %d\n", runtime.NumCPU())
	//
	fmt.Printf("GOROOT: %s\n", runtime.GOROOT())
	//
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
}

func PrintFunInfo(layer int) bool {
	// get call pc
	pc, _, _, ok := runtime.Caller(layer)
	if !ok {
		fmt.Println("caller is failed!")
		return false
	}

	// create type Func
	funinfo := runtime.FuncForPC(pc)
	file, line := funinfo.FileLine(funinfo.Entry())
	fmt.Printf("file : %s\nline : %d\n", file, line)
	fmt.Printf("fun name : %s\n", funinfo.Name())

	return true
}

func SetEnv() {
	// open GC debug
	os.Setenv("GOGCTRACE", "1")
}

func UnSetEnv() {
	os.Unsetenv("GOGCTRACE")
}

func display(num int) {
	fmt.Println("goroutine num: ", runtime.NumGoroutine())

	for i := 1; i < num; i++ {
		fmt.Println(i)
		//runtime.Gosched()
	}

	wg.Done()

	runtime.Goexit()

	fmt.Println("Goruntine exit")
}

// system func
func init() {
	fmt.Println("init start : ------------------------------------------------------------------------------------------")
	//
	SetEnv()

	//
	PrintGoInfo()

	// print cpu info
	PrintCpuInfo()

	fmt.Println("init start : ------------------------------------------------------------------------------------------")
}

func clean() {
	//
	UnSetEnv()
}

func main() {
	//print fun info
	if ok := PrintFunInfo(1); !ok {
		return
	}

	cpunum := runtime.NumCPU()
	wg.Add(cpunum)

	//run GC
	runtime.GC()

	// set max cpu
	runtime.GOMAXPROCS(cpunum)

	go display(10)
	go display(10)

	wg.Wait()
}
