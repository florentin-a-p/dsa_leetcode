package main

import "fmt"

func sumNumbers2(numList []int) int {
	var sum int
	fmt.Println("first call: ", numList, " sum: ", sum)

	if len(numList) == 0 { //stopping condition
		sum = 0 //what the last step looks like
	} else if len(numList) > 0 { //when to call the method
		first := numList[0]

		//fmt.Println("before calling next: ", numList, " sum: ", sum)
		sum = first + sumNumbers2(numList[1:]) //what to do with the return value
	}
	fmt.Println("before return: ", numList, " sum: ", sum)
	return sum
}

func main() {
	numList := []int{-1, 2, 3, 4}
	sumAllNumbers := sumNumbers2(numList)
	fmt.Println("Sum of all numbers:", sumAllNumbers)
}
