package main

import "fmt"

func binarySearch(array []int, num int) bool {
	var found bool
	fmt.Println("first call: ", array, " found: ", found)

	if len(array) == 1 { //1st stopping point, when the array is empty
		if num == array[0] {
			found = true
		} else {
			found = false
		}
	} else if len(array) > 1 {
		midPoint := len(array) / 2
		if num == array[midPoint] { //2nd stopping point, when found
			found = true
		} else if num > array[midPoint] { //recurse on right
			found = binarySearch(array[midPoint+1:], num)
		} else if num < array[midPoint] { //recurse on left
			found = binarySearch(array[:midPoint], num)
		}
	}
	fmt.Println("before return: ", array, " found: ", found)
	return found
}

func main() {
	numArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	answer := binarySearch(numArray, 14)
	fmt.Println("the binary search result is: ", answer)
}
