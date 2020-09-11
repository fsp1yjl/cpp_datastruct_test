/*
link: https://leetcode.com/problems/house-robber-iii/
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

Example 1:

Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
Example 2:

Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.

*/

package main

import "fmt"

func main() {
	var r *TreeNode
	r = buildOne(3)
	r.Left = buildOne(4)
	r.Right = buildOne(5)
	r.Left.Right = buildOne(1)

	r.Left.Right = buildOne(3)
	r.Right.Right = buildOne(1)

	fmt.Println("r:", rob(r))

}

func buildOne(val int) *TreeNode {
	r := &TreeNode{}
	r.Val = val

	return r
}

// leetcode提交时struct已经定义好，不需要显示定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
//思路： 这里的思路是将rob(root)转化为 root 参与计算和 root不参与计算两个可能去计算

结果：
Runtime: 1172 ms, faster than 12.61% of Go online submissions for House Robber III.
Memory Usage: 5.2 MB, less than 59.46% of Go online submissions for House Robber III.

*/
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := root.Left
	r := root.Right

	var l_l *TreeNode
	var l_r *TreeNode
	var r_l *TreeNode
	var r_r *TreeNode

	if l != nil {
		l_l = l.Left
		l_r = l.Right
	}

	if r != nil {
		r_l = r.Left
		r_r = r.Right
	}

	s1 := root.Val + rob(l_l) + rob(l_r) + rob(r_l) + rob(r_r)
	s2 := rob(l) + rob(r)

	if s1 > s2 {
		return s1
	} else {
		return s2
	}
}
