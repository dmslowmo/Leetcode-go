package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
        Left *TreeNode
        Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	temp := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = temp
	return root
}

func main() {
	left := TreeNode{1, nil, nil}
	right := TreeNode{2, nil, nil}
	root := TreeNode{3, &left, &right}
	inverted := invertTree(&root)
	fmt.Println(inverted.Left.Val)
	fmt.Println(inverted.Right.Val)

}
