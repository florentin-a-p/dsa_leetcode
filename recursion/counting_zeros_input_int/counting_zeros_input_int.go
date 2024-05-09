package main

import (
	"fmt"
	"math"
)

func countZeros(n int) int {
	// Handle negative numbers by taking the absolute value
	n = int(math.Abs(float64(n)))
	var countZero int
	fmt.Println("first call: ", n)

	if n < 10 { //stopping point, last digit
		if n == 0 {
			countZero = 1 //what the last step looks like
		} else {
			countZero = 0 //what the last step looks like
		}
	} else if n >= 10 { //when to call the method
		// Check if the last digit is zero
		countZero = 0
		if n%10 == 0 {
			countZero = 1
		}
		countZero = countZero + countZeros(n/10) //what we do with the return value
	}
	fmt.Println("before return: ", n)
	return countZero
}

func main() {
	number := 30204
	result := countZeros(number)
	fmt.Printf("The number of zeros in %d is %d\n", number, result)
}
