/*

link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”



Example 1:


Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
Example 2:


Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.


二叉树的最近公共祖先

*/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归解法， 提交结果：
Runtime: 12 ms, faster than 75.71% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
Memory Usage: 7.1 MB, less than 94.29% of Go online submissions for Lowest Common Ancestor of a Binary Tree.

*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)

	r := lowestCommonAncestor(root.Right, p, q)

	//如果左右递归均有捕获，这此时的root为分叉的开始，直接返回root
	if l != nil && r != nil {
		return root
	}

	// 如果仅r不为空，则表示问题被缩小到右子树
	if l == nil && r != nil {
		return r
	}

	// 问题被缩小到左子树
	if l != nil && r == nil {
		return l
	}

	return nil

}