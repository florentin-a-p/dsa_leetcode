package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList
	counter := 1

	for currentNode.Next != nil {
		counter = counter + 1
		currentNode = currentNode.Next
	}

	counter = counter + 1
	fmt.Println("counter: ", counter)

	middlePoint := counter / 2
	if counter%2 == 0 {
		middlePoint = middlePoint - 1
	} else {
		middlePoint = middlePoint + 1
	}

	currentNode = linkedList

	for counter = 1; counter < middlePoint; counter++ {
		currentNode = currentNode.Next
	}

	return currentNode
}

//traverse through the linked list to get the length, counter starts from 1
//divide the length by 2 to get middle point, if it's even subtract 1, if it's odd add 1
//re-traverse the linekd list again until the middle point
//return
