package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists(left, right)
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
	list1Node3 := &ListNode{5, nil}
	list1Node2 := &ListNode{4, list1Node3}
	list1Node1 := &ListNode{1, list1Node2}

	list2Node3 := &ListNode{4, nil}
	list2Node2 := &ListNode{3, list2Node3}
	list2Node1 := &ListNode{1, list2Node2}

	list3Node2 := &ListNode{6, nil}
	list3Node1 := &ListNode{2, list3Node2}

	lists := []*ListNode{list1Node1, list2Node1, list3Node1}

	mergeKLists(lists)

}
