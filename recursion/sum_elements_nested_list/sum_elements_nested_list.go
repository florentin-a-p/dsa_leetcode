package main

import "fmt"

func isTypeInt(obj interface{}) bool {
	_, ok := obj.(int)
	if ok {
		return true
	} else {
		return false
	}
}

func sumElementsNestedList(array []interface{}) int {
	var sum int
	fmt.Println("first call: ", array, " sum: ", sum)
	if len(array) == 0 { //stopping point
		sum = 0
	} else if len(array) > 0 {
		first := array[0]
		if isTypeInt(first) {
			sum = sum + first.(int)
		} else {
			firstObj, _ := first.([]interface{})
			sum = sum + sumElementsNestedList(firstObj)
		}
		sum = sum + sumElementsNestedList(array[1:])
	}
	fmt.Println("before return: ", array, " sum: ", sum)
	return sum
}

//func sumElementsNestedList(array []interface{}) int {
//	var sum int
//	if len(array) == 0 { // Base case: empty slice
//		return 0
//	}
//
//	// Process the first element of the slice
//	first := array[0]
//	switch v := first.(type) {
//	case int:
//		sum += v // Directly add if it's an integer
//	case []interface{}:
//		sum += sumElementsNestedList(v) // Recursively sum if it's a nested list
//	default:
//		// Handle other unexpected types
//		fmt.Println("Unsupported type found")
//	}
//
//	// Recursively process the rest of the slice
//	sum += sumElementsNestedList(array[1:])
//	return sum
//}

func main() {
	array := []interface{}{
		1,
		2,
		[]interface{}{3, 4},
		5,
		6}

	answer := sumElementsNestedList(array)
	fmt.Println("sum of all elements in nested list is: ", answer)

}

//loop over the array
//check if the element is of type array -> yg ini, anggep aja udh keimplement, biar bs fokus di recursion
//yes -> call recursion function
//no -> add the element to the existing sum
