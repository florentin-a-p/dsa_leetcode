package main

import "fmt"

func countZeros(array []int) int {
	var countZero int
	fmt.Println("first call: ", array, " countZero: ", countZero)

	if len(array) == 0 { //stopping point
		countZero = 0 //what the last step looks like
	} else if len(array) > 0 { //when we call the method
		first := array[0]
		if first == 0 {
			countZero = countZero + 1
		}
		countZero = countZero + countZeros(array[1:])
	}
	fmt.Println("before return: ", array, " countZero: ", countZero)
	return countZero
}

func main() {
	numberArray := []int{3, 0, 2, 0, 4}

	answer := countZeros(numberArray)
	fmt.Println("the count of zeros: ", answer)

	//for i := 0; i < len(numberArray); i++ {
	//	if numberArray[i] == 0 {
	//		countZero++
	//	}
	//}
	//fmt.Println("countZero ", countZero)
}
