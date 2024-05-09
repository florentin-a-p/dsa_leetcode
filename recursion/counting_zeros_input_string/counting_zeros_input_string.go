package main

import "fmt"

func countZerosInputString(array []rune) int {
	var countZero int
	fmt.Println("first call: ", array, " countZero: ", countZero)

	if len(array) == 0 { //stopping point
		countZero = 0 //what the last step looks like
	} else if len(array) > 0 { //when we call the method
		first := array[0]
		if string(first) == "0" {
			countZero = countZero + 1
		}
		countZero = countZero + countZerosInputString(array[1:])
	}
	fmt.Println("before return: ", array, " countZero: ", countZero)
	return countZero
}

func main() {
	numberString := "3002"

	//for i := 0; i < len(numberString); i++ {
	//	fmt.Println(string(numberString[i]))
	//}

	numberArray := []rune(numberString)

	answer := countZerosInputString(numberArray)
	fmt.Println("the count of zeros: ", answer)
}
