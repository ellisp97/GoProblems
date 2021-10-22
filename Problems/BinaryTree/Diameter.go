package Problems

import "github.com/ellisp97/GoProblems/Utility"

var diameter int

func diameterOfBinaryTree(root *Utility.TreeNode) int {
	diameter = 0
	helperdiameterOfBinaryTree(root)
	return diameter
}

func helperdiameterOfBinaryTree(root *Utility.TreeNode) int {
	if root == nil {
		return -1
	}

	left := helperdiameterOfBinaryTree(root.Left)
	right := helperdiameterOfBinaryTree(root.Right)

	diameter = max(diameter, 2+left+right)
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
