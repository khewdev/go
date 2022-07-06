/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if !(root.Val < right && root.Val > left) {
		return false
	}

	return validate(root.Left, left, root.Val) && validate(root.Right, root.Val, right)
}

func main() {
	node1 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2}
	node2.Left = node1
	node2.Right = node3

	fmt.Println(isValidBST(node2))
}
