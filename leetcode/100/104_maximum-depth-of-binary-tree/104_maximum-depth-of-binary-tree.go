/*
link: https://leetcode.com/problems/maximum-depth-of-binary-tree/

Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.

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
// 求二叉树的最大深度，这里直接使用递归的方式处理
提交结果：
Runtime: 4 ms, faster than 91.45% of Go online submissions for Maximum Depth of Binary Tree.
Memory Usage: 4.4 MB, less than 12.66% of Go online submissions for Maximum Depth of Binary Tree.
Next challenges:

*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth_l := maxDepth(root.Left)
	depth_r := maxDepth(root.Right)

	if depth_l > depth_r {
		return 1 + depth_l
	} else {
		return 1 + depth_r
	}
}