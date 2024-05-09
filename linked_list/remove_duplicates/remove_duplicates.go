package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	cacheNode := linkedList
	currentNode := linkedList

	for currentNode.Next != nil {
		if currentNode.Value > cacheNode.Value {
			cacheNode.Next = currentNode
			cacheNode = currentNode
		}
		currentNode = currentNode.Next
	}

	if currentNode.Next == nil {
		if currentNode.Value == cacheNode.Value {
			fmt.Println("cache: ", cacheNode)
			fmt.Println("currentNode: ", currentNode)
			//cacheNode.Next = nil
			cacheNode.Next = currentNode.Next
		}
	}

	return linkedList
}

func main() {
	// Array of values to be added to the linked list
	values := []int{1, 1, 3, 4, 4, 4, 5, 6, 6}

	// Create the head of the linked list
	var head *LinkedList
	var current *LinkedList

	// Loop through each value and add it to the linked list
	for _, value := range values {
		newNode := &LinkedList{Value: value}
		if head == nil {
			// Set the first node as the head
			head = newNode
			current = head
		} else {
			// Append new node to the list
			current.Next = newNode
			current = newNode
		}
	}

	// Printing the linked list to verify it's constructed properly
	current = head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}

	RemoveDuplicatesFromLinkedList(head)
}
