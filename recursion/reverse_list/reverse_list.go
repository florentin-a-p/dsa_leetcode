package main

import (
	"fmt"
)

// reverseList recursively processes the list by moving the first element to the end.
func reverseList(numList []int) []int {
	fmt.Println("first call: ", numList)

	if len(numList) == 0 { //stopping point
		numList = []int{} //what the last step looks like
	} else if len(numList) > 0 { //when to call the method
		first := numList[0]
		numList = append(reverseList(numList[1:]), first) //what to do with the return
	}
	fmt.Println("before return: ", numList)
	return numList
}

func main() {
	numList := []int{3, 5, 10, 2}
	fmt.Println("Original list:", numList)

	// Process the list
	processedList := reverseList(numList)
	fmt.Println("Processed list:", processedList)
}
