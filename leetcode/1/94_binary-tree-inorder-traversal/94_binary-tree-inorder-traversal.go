/*
二叉树的中序遍历

*/
package l94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归解法：
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	if root.Left != nil {
		l := inorderTraversal(root.Left)
		res = append(res, l...)
	}
	res = append(res, root.Val)

	if root.Right != nil {
		r := inorderTraversal(root.Right)
		res = append(res, r...)
	}

	return res
}

/*
非递归解法
*/

func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	p := root
	stack := make([]*TreeNode, 0)
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		l := len(stack)
		if l != 0 {
			pop := stack[l-1]
			res = append(res, pop.Val)
			stack = stack[:l-1]
			p = pop.Right
		}
	}

	return res

}
