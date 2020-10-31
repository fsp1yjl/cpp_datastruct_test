/*
https://leetcode.com/problems/symmetric-tree/

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
 

But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3
 

Follow up: Solve it both recursively and iteratively.
*/



package main
func main() {

}
type TreeNode struct {
	    Val int
	     Left *TreeNode
	     Right *TreeNode
	 }


func isSymmetric(root *TreeNode) bool {
    if root == nil {
		return true 
	}
	if root.Left == nil && root.Right == nil {
		return true 
	}
	if root.Left ==nil || root.Right == nil {
		return false 
	}

	return isSymmetricRecursive(root.Left, root.Right)


}

//****递归解法
func isSymmetricRecursive(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true 
	}
	if left == nil || right == nil {
		return false 
	}

	if left.Val != right.Val {
		return false 
	}

	return isSymmetricRecursive(left.Left, right.Right) && isSymmetricRecursive(left.Right, right.Left)
}

//** 非递归解法
func isSymmetricIter(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true 
	}
	if left == nil || right == nil {
		return false 
	}

	// 这里l,r为两个队列，先进先出
	l := make([]*TreeNode,0)
	r := make([]*TreeNode, 0)

	l = append(l, left)
	r = append(r, right)

	for len(l) != 0 && len(r) !=0 {
		l1 := l[0]
		r1 := r[0]

		if (l1 == nil && r1 != nil) || (l1!=nil && r1 == nil) {
			return false 
		}
        // bug1 这里需要考虑节点为nil的情况
		if l1 == nil && r1 == nil {
			l = l[1:]
			r = r[1:]
			continue
		}

		if l1.Val != r1.Val {
			return false
		} else {
			l = l[1:]
			r = r[1:]
		}
		
		// 这里为nil的子节点也会入队列
		// if l1.Left != nil {
			l = append(l, l1.Left)
		// }
		// if l1.Right != nil {
			l = append(l, l1.Right)
		// }

		// if r1.Right != nil {
			r = append(r, r1.Right)
		// }
		// if r1.Left != nil {
			r = append(r, r1.Left)
		// }

		if len(l) != len(r) {
			return false 
		}
	}

	return true

}



