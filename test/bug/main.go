package main

import "fmt"

func Afunc() {

	errch := make(chan string, 1)
	resch := make(chan int, 1)

	go func(errch chan string, resch chan int) {
		defer close(errch)
		defer close(resch)
		errch <- string("panlin")
		//fmt.Println(len(errch))
	}(errch, resch)

	select {
	case <-errch:
	case <-resch:
		fmt.Println(len(errch))
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxx")
	}
}

func main() {
	testch := make(chan string, 1)
	close(testch)
	fmt.Println(<-testch)

	for i := 0; i < 1000000; i++ {
		Afunc()
	}
}
