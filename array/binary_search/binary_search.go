package main

import "fmt"

func binary_search(array []int, num int) bool {
	found := false
	arraySize := len(array)
	midPoint := (arraySize / 2)
	leftPoint := 0
	rightPoint := (arraySize - 1)

	for leftPoint <= rightPoint && !found {
		fmt.Println("loop, leftPoint: ", leftPoint, " rightPoint: ", rightPoint, " midPoint: ", midPoint)
		if num == array[midPoint] {
			found = true
		} else if num < array[midPoint] {
			//fmt.Println("num on left")
			rightPoint = midPoint - 1
			midPoint = (leftPoint + rightPoint) / 2
		} else if num > array[midPoint] {
			//fmt.Println("num on right")
			leftPoint = midPoint + 1
			midPoint = (leftPoint + rightPoint) / 2
		}
	}
	return found
}

func main() {
	numArray := []int{1, 2, 3, 4, 5}
	answer := binary_search(numArray, 2)
	fmt.Println("the binary search result is: ", answer)
}
