/*
link: https://leetcode.com/problems/minimum-depth-of-binary-tree/


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
// 提交结果：
Runtime: 240 ms, faster than 48.33% of Go online submissions for Minimum Depth of Binary Tree.
Memory Usage: 24.1 MB, less than 10.00% of Go online submissions for Minimum Depth of Binary Tree.

*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	min_l := minDepth(root.Left)
	min_r := minDepth(root.Right)

	// bug1 需要处理左子树和右子树为空的情况
	if root.Left == nil {
		return 1 + min_r
	}
	if root.Right == nil {
		return 1 + min_l
	}

	if min_l < min_r {
		return 1 + min_l
	} else {
		return 1 + min_r
	}
}
