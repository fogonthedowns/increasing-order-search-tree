package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	ans := &TreeNode{}
	prev = ans

	inOrder(root)

	return ans.Right
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)

	// switch
	prev.Right = root
	prev = root
	prev.Left = nil

	inOrder(root.Right)
	return
}
