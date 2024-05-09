package main

import "fmt"

// should be in ascending order only

func isIntArrayInOrder(array []int) bool {
	var isInOrder bool

	if len(array) == 1 { //stopping point
		isInOrder = true
	} else {
		first := array[0]
		if first > array[1] {
			isInOrder = false //stopping point
		} else { //when to call the method recursively
			isInOrder = isIntArrayInOrder(array[1:]) //what to do with the return value
		}
	}

	return isInOrder
}

func main() {
	array := []int{1, 2, 3, 4, 3}
	answer := isIntArrayInOrder(array)
	fmt.Print("array is in order? ", answer)
}
