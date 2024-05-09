package main

import "fmt"

func searchElement(n int, array []int) bool {
	var found bool
	fmt.Println("first call: ", array, " found: ", found)

	if len(array) < 1 { //stopping point
		found = false //what the last step looks like
	} else {
		first := array[0]
		if n == first { //stopping point
			found = true
		} else { //when we call recursively
			found = searchElement(n, array[1:])
		}
	}
	fmt.Println("before return: ", array, " found: ", found)
	return found
}

func main() {
	testArray := []int{2, 3, 6, 6}
	isFound := searchElement(6, testArray)
	fmt.Println("result is: ", isFound)
}
