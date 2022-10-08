// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// recursive: T O(n), M O(n)
	if head == nil {
		return nil
	}

	newHead := head
	fmt.Println("head.Val", head.Val)
	fmt.Println("head.Next", head.Next)
	fmt.Println("newHead.Val", newHead.Val)
	fmt.Println("newHead.Next", newHead.Next)
	if head.Next != nil {
		fmt.Println("Passing", head.Next, "to recursion")
		newHead = reverseList(head.Next)
		fmt.Println("Returned newHead.Val", newHead.Val)
		fmt.Println("head.Val", head.Val)
		head.Next.Next = head
	}
	head.Next = nil

	fmt.Println("Returning newHead.Val", newHead.Val)
	return newHead
}

func main() {
	node3 := ListNode{3, nil}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}

	fmt.Println(reverseList(&node1))
}
