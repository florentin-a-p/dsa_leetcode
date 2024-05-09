package main

import "fmt"

type SpecialArray []interface{}

//// ProductSum computes the sum of a "special" array, where sub-arrays increase the sum multiplier
//func ProductSum(array SpecialArray, multiplier int) int {
//	var sum int
//	if len(array) == 0 { //stopping point
//		sum = 0
//	} else if len(array) > 0 {
//		item := array[0]
//
//		switch value := item.(type) {
//		case int:
//			sum += value
//		case SpecialArray:
//			multiplier++
//			sum += multiplier * ProductSum(value, multiplier)
//			multiplier--
//		default:
//			panic(fmt.Sprintf("unexpected type %T in the array", item))
//		}
//		sum += ProductSum(array[1:], multiplier)
//	}
//	return sum
//}

func ProductSum(array SpecialArray) int {
	return ProductSumHelper(array, 1)
}

func ProductSumHelper(array SpecialArray, multiplier int) int {
	var sum int
	if len(array) == 0 { //stopping point
		sum = 0
	} else if len(array) > 0 {
		item := array[0]

		switch value := item.(type) {
		case int:
			sum += value
		case SpecialArray:
			multiplier++
			sum += multiplier * ProductSumHelper(value, multiplier)
			multiplier--
		default:
			panic(fmt.Sprintf("unexpected type %T in the array", item))
		}
		sum += ProductSumHelper(array[1:], multiplier)
	}
	return sum
}

func main() {
	testArray := SpecialArray{
		5,
		2,
		SpecialArray{7, -1},
		3,
		SpecialArray{6,
			SpecialArray{-13, 8},
			4},
	}
	result := ProductSum(testArray)
	fmt.Println("Product Sum:", result)
}
