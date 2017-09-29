package main

import "fmt"

func fun1() (result int) {
	result = 0
	defer func() {
		fmt.Printf("defer result: %d\n", result)
		result++
	}()

	return 100
}

func fun2() int {
	result := 0
	defer func() {
		fmt.Printf("defer result: %d\n", result)
		result++
	}()

	return result
}

func fun3(input int) (result int) {

	fmt.Printf("input: 0x%x\nresult: 0x%x\n", &input, &result)

	return result
}

func fun4() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	fmt.Printf("func1 : %d\n", fun1())
	fmt.Printf("func2 : %d\n", fun2())
	fmt.Printf("func3 : %d\n", fun3(1))
	fmt.Printf("func4 : %d\n", fun4())
}
