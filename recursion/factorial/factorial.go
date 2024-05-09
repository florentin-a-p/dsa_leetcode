package main

import "fmt"

func factorial(n int) int {
	var result int
	fmt.Println("first call: ", n, " result: ", result)

	if n <= 0 { // base case or stopping point
		result = 1
	} else if n > 0 { // when do we call the method recursively
		result = n * factorial(n-1) // what we do with the return from the call
	}
	//fmt.Println("result is ", result, " from n ", n)
	fmt.Println("before return: ", n, " result: ", result)
	return result
}

func main() {
	n := 7
	answer := factorial(n)
	fmt.Print("the factorial of ", n, " is ", answer)
}
