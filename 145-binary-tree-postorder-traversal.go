// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, arr)
	postorder(root.Right, arr)
	*arr = append(*arr, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	arr := []int{}
	postorder(root, &arr)
	return arr
}

func main() {
	node3 := TreeNode{Val: 3, Left: nil, Right: nil}
	node2 := TreeNode{Val: 2, Left: &node3, Right: nil}
	node1 := TreeNode{Val: 1, Left: nil, Right: &node2}

	fmt.Println(postorderTraversal(&node1))
}
