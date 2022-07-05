package main

import "fmt"

type node struct {
	Data int
	Next *node
}

type linkedList struct {
	Head   *node
	Length int
}

func (l *linkedList) prepend(n *node) {
	currentHead := l.Head
	l.Head = n
	l.Head.Next = currentHead
	l.Length++
	fmt.Println("Inserted", n.Data, "into linked list")
}

func (l *linkedList) delete(key int) {
	if l.Length == 0 {
		return
	}

	fmt.Println("Searching data to delete(HEAD):", l.Head.Data)
	if l.Head.Data == key {
		fmt.Println("Found", key, "in Head, removing...")
		l.Head = l.Head.Next
		l.Length--
		return
	}

	currentNode := l.Head
	for currentNode != nil {
		fmt.Println("Searching data to delete:", currentNode.Next.Data)

		if currentNode.Next.Data == key {
			fmt.Println("Found", key, "from linked list, removing...")
			currentNode.Next = currentNode.Next.Next
			l.Length--
			return
		} else {
			if currentNode.Next.Next == nil {
				fmt.Println(key, "is not in the linked list")
				return
			}
			currentNode = currentNode.Next
		}
	}
}

func (l linkedList) printLinkedListData() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println("Printing data:", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func main() {
	myList := &linkedList{}
	node1 := &node{Data: 1}
	node2 := &node{Data: 2}
	node3 := &node{Data: 3}
	node4 := &node{Data: 4}
	node5 := &node{Data: 5}
	node6 := &node{Data: 6}
	node7 := &node{Data: 7}
	node8 := &node{Data: 100}

	myList.prepend(node8)
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)
	fmt.Println()
	myList.printLinkedListData()
	myList.delete(7)
	myList.delete(100)
	myList.delete(333)
	myList.delete(4)
	fmt.Println()
	fmt.Println("After deleting")
	myList.printLinkedListData()
}
