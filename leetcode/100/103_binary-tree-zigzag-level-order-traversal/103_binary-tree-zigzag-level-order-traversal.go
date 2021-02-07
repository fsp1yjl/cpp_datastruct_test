/*

link: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/


Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
*/

package main

import "fmt"

func main() {
	s1 := -56 + 33 + 583 + 231 - 330
	s2 := -231 - 121 - 942 + 5604 - 197
	s3 := -279 - 229 + 50 - 95
	s4 := -1000 + 7085 + 209 + 40 - 649 + 517 + 156
	s5 := -479 - 1832 + 272 - 356 + 762
	s6 := -507 - 9 + 707 + 134 - 84

	fmt.Println("sum:", s1+s2+s3+s4+s5+s6)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	// q1 := make([]int, 0)
	q2 := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	// q1 = append(q1, root.Val)
	q2 = append(q2, root)

	leftFlag := true

	for len(q2) != 0 {
		tmp := make([]*TreeNode, 0)
		q1 := make([]int, 0)
		l := len(q2)
		for i := 0; i < l; i++ {
			q1 = append(q1, q2[i].Val)
			left := q2[i].Left
			right := q2[i].Right

			if left != nil {
				tmp = append(tmp, left)
			}
			if right != nil {
				tmp = append(tmp, right)
			}

		}
		//对需要反转对行数据进行反转
		if !leftFlag {
			ll := len(q1)
			for i := 0; i < ll/2; i++ {
				// fmt.Println("qi,q---i:", q1[i], q1[ll-i-1])
				q1[i], q1[ll-i-1] = q1[ll-i-1], q1[i]
				// fmt.Println("after qi,q---i:", q1[i], q1[ll-i-1])
			}
		}

		leftFlag = !leftFlag
		res = append(res, q1)

		q2 = tmp
	}

	return res

}
