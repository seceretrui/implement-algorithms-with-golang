package main

import "fmt"

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(inorderTraversal(root))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}
