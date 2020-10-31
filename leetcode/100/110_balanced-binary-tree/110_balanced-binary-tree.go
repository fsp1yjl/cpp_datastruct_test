/*

判断一棵二叉树是否是二叉平衡树

思路： 根据定义，判断左右子树的最大深度相差是否超过1, 最大深度算法直接采用leetcode 104
并对子树做同样对处理、

*/

package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
提交结果：
Runtime: 8 ms, faster than 77.21% of Go online submissions for Balanced Binary Tree.
Memory Usage: 5.7 MB, less than 100.00% of Go online submissions for Balanced Binary Tree.

*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	depth_l := maxDepth(root.Left)
	depth_r := maxDepth(root.Right)

	sub := depth_l - depth_r

	if sub < 0 {
		sub = 0 - sub
	}

	if sub > 1 {
		return false
	} else {
		// 递归判断左 右子树是否平衡
		fmt.Println("herre")
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}

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
