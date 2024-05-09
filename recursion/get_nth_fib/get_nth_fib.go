package main

import "fmt"

func GetNthFib(n int) int {
	var sum int
	fmt.Println("first call: ", n, " sum: ", sum)

	if n > 2 {
		sum = GetNthFib(n-1) + GetNthFib(n-2)
	} else if n == 2 {
		sum = 1
	}

	fmt.Println("before return: ", n, " sum: ", sum)
	return sum
}

func main() {
	total := GetNthFib(5)
	fmt.Println("sum of fibonacci", total)
}
