package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func main() {
	l1Node3 := &ListNode{4, nil}
	l1Node2 := &ListNode{2, l1Node3}
	l1Node1 := &ListNode{1, l1Node2}

	l2Node3 := &ListNode{4, nil}
	l2Node2 := &ListNode{3, l2Node3}
	l2Node1 := &ListNode{1, l2Node2}

	mergeTwoLists(l1Node1, l2Node1)
}
