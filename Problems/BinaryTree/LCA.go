package Problems

import (
	"fmt"

	"github.com/ellisp97/GoProblems/Utility"
)

/*

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”


Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.

*/
func LCA(root, p, q *Utility.TreeNode) *Utility.TreeNode {
	fmt.Println(root)

	if root == nil || root == p || root == q {
		return root
	}

	left := LCA(root.Left, p, q)
	right := LCA(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}

}

func LCA_BST(root, p, q *Utility.TreeNode) *Utility.TreeNode {

	if p.Val < root.Val && q.Val < root.Val {
		return LCA_BST(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return LCA_BST(root.Right, p, q)
	}

	return root
}
