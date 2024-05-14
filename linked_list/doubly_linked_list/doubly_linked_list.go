// Feel free to add methods and fields to the struct definitions.
package main

import "fmt"

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	// Write your code here.
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	// set the prev of the head as the node
	if ll.Head == nil {
		ll.Head = node
	} else {
		ll.Head.Prev = node
	}
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	// set the next of the tail as node
	if ll.Tail == nil {
		ll.Tail = node
	} else {
		ll.Tail.Next = node
	}
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	// traverse the linked list until you get to the node value with input node
	// store the node in cache
	// go the previous node, set the next to nodeToInsert
	// set the cache's prev to nodeToInsert
	currentNode := ll.Head

	for currentNode.Next != nil {
		if currentNode.Value == node.Value {
			cacheNode := currentNode
			currentNode.Prev.Next = nodeToInsert
			cacheNode.Prev = nodeToInsert
			break
		}

		currentNode = currentNode.Next
	}
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	// traverse the linked list until you get to the node value with input node
	// store the node in cache
	// go the next node, set the previous to nodeToInsert
	// set the cache's next to nodeToInsert
	currentNode := ll.Head

	for currentNode.Next != nil {
		if currentNode.Value == node.Value {
			cacheNode := currentNode
			currentNode.Next.Prev = nodeToInsert
			cacheNode.Next = nodeToInsert
			break
		}

		currentNode = currentNode.Next
	}
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	// set counter to 1
	// traverse until you get to that position, increment the counter
	// once you find the node  at the position, store it in cache
	// go to the next node, set the previous to nodeToInsert
	// set the cached node's next to nodeToInsert

	counter := 1
	currentNode := ll.Head

	for currentNode.Next != nil {
		if counter == position {
			cacheNode := currentNode
			currentNode.Prev.Next = nodeToInsert
			cacheNode.Prev = nodeToInsert
			break
		}

		counter = counter + 1
		currentNode = currentNode.Next
	}
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	// similar to remove duplicates
	// traverse the linked list, then remove the nodes (use cache)

	currentNode := ll.Head
	for currentNode.Next != nil {
		if currentNode.Value == value {
			cacheNode := currentNode.Prev
			currentNode.Next.Prev = cacheNode
			break
		}
		currentNode = currentNode.Next
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	// traverse the linked list, then remove the node (use cache)

	currentNode := ll.Head
	for currentNode.Next != nil {
		if currentNode.Value == node.Value {
			cacheNode := currentNode.Prev
			currentNode.Next.Prev = cacheNode
			break
		}
		currentNode = currentNode.Next
	}

}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	// traverse the linked list and return true once you find a node with that value
	currentNode := ll.Head
	for currentNode.Next != nil {
		if currentNode.Value == value {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func main() {
	myLinkedList := NewDoublyLinkedList()
	myFirstNode := &Node{
		Value: 1,
	}

	mySecondNode := &Node{
		Value: 2,
	}

	myThirdNode := &Node{
		Value: 3,
	}

	//myLinkedList.SetHead(myFirstNode)
	fmt.Println("====== initializing 1st and 2nd node=======")
	fmt.Println("myLinkedList", myLinkedList)
	fmt.Println("myFirstNode", myFirstNode)
	fmt.Println("mySecondNode", mySecondNode)

	//myFirstNode.Next = mySecondNode
	//mySecondNode.Prev = myFirstNode
	//fmt.Println("====== setting next of 1st node and prev of 2nd node =======")
	//fmt.Println("myLinkedList", myLinkedList)
	//fmt.Println("myFirstNode", myFirstNode)
	//fmt.Println("mySecondNode", mySecondNode)

	myLinkedList.SetHead(myFirstNode)
	fmt.Println("======= set head of linkedlist with 1st node ======")
	fmt.Println("myLinkedList", myLinkedList)
	fmt.Println("myFirstNode", myFirstNode)
	fmt.Println("mySecondNode", mySecondNode)

	myLinkedList.SetTail(mySecondNode)
	fmt.Println("======= set tail of linkedlist with 2nd node ======")
	fmt.Println("myLinkedList", myLinkedList)
	fmt.Println("myFirstNode", myFirstNode)
	fmt.Println("mySecondNode", mySecondNode)

	myLinkedList.SetTail(myThirdNode)
	fmt.Println("======= set tail of linkedlist with 3rd node ======")
	fmt.Println("myLinkedList", myLinkedList)
	fmt.Println("myFirstNode", myFirstNode)
	fmt.Println("mySecondNode", mySecondNode)
	fmt.Println("myThirdNode", myThirdNode)

	//myLinkedList.InsertAfter(myFirstNode, mySecondNode)
	//
	//fmt.Println("======= insert 2nd node after head ======")
	//fmt.Println("myLinkedList", myLinkedList)
	//fmt.Println("myFirstNode", myFirstNode)
	//fmt.Println("mySecondNode", mySecondNode)
	//
	//myLinkedList.InsertAfter(mySecondNode, myThirdNode)
	//
	//fmt.Println("======= insert 3rd node after 2nd node ======")
	//fmt.Println("myLinkedList", myLinkedList)
	//fmt.Println("myFirstNode", myFirstNode)
	//fmt.Println("mySecondNode", mySecondNode)
	//fmt.Println("myThirdNode", myThirdNode)
	//
	fmt.Println("======= traversing doubly linked list ======")

	currentNode := myLinkedList.Head
	for currentNode.Next != nil {
		fmt.Println("node value is: ", currentNode.Value, " node is: ", currentNode)
		currentNode = currentNode.Next
	}

	//myLinkedList.SetTail(mySecondNode)
	//fmt.Println("======= set tail of linkedlist with 2nd node ======")
	//fmt.Println("myLinkedList", myLinkedList)
	//fmt.Println("myFirstNode", myFirstNode)
	//fmt.Println("mySecondNode", mySecondNode)
}
