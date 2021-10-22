package Utility

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeStack struct {
	Stack
}

func (s *TreeNodeStack) Push(n *TreeNode) { s.Stack.Push(n) }
func (s *TreeNodeStack) Pop() *TreeNode   { return s.Stack.Pop().(*TreeNode) }
