package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
这里只需要判断是否存在，而不用给出所有解的方式，直接递归处理即可

另外的思路是可以进行一次广度优先搜索，对每层进行加权，这样最后对叶子节点层即是sum

提交结果：
Runtime: 4 ms, faster than 96.17% of Go online submissions for Path Sum.
Memory Usage: 4.8 MB, less than 100.00% of Go online submissions for Path Sum.
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		} else {
			return false
		}
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
