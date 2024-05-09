package main

import (
	"fmt"
)

func sumNumbers(sum int, first int, numList []int) int {
	if len(numList) > 0 { //stopping condition
		first := numList[0]
		sum = sumNumbers(sum, first, numList[1:])
	}
	sum = sum + first
	return sum
}

func main() {
	//numList := []int{3, 5, 10, 2}
	numList := []int{-1, 2, 3, 4}

	sum := 0
	first := 0

	// Process the list
	sumAllNumbers := sumNumbers(sum, first, numList)
	fmt.Println("Sum of all numbers:", sumAllNumbers)

}
