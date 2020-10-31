/*
link : https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]


*/

package main

import "fmt"

func main() {
	nums := []int{1}

	fmt.Println(nums[1:])
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：
用一个二维数组队列存放每一层对节点信息，每次取本层节点处理后出队列，并将下层节点入队列
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
Memory Usage: 2.9 MB, less than 30.00% of Go online submissions for Binary Tree Level Order Traversal II.

*/

func levelOrderBottom(root *TreeNode) [][]int {
	arr := make([][]int, 0)
	nodesArr := make([][]*TreeNode, 0)

	if root == nil {
		return arr
	}

	nodesArr = append(nodesArr, []*TreeNode{root})

	for len(nodesArr) != 0 {
		nodes := nodesArr[0]

		tmp := make([]int, 0)
		childNodes := make([]*TreeNode, 0)
		l := len(nodes)
		for i := 0; i < l; i++ {
			tmp = append(tmp, nodes[i].Val)
			if nodes[i].Left != nil {
				childNodes = append(childNodes, nodes[i].Left)
			}

			if nodes[i].Right != nil {
				childNodes = append(childNodes, nodes[i].Right)
			}
		}

		//老节点出队列，新节点入队列
		if len(childNodes) != 0 {
			nodesArr = append(nodesArr, childNodes)
		}
		nodesArr = nodesArr[1:]

		arr = append(arr, tmp)
	}

	// 最后对二维数组做一次反转
	newArr := make([][]int, 0)
	ll := len(arr) - 1

	for i := ll; i >= 0; i-- {
		tmp := arr[i]
		newArr = append(newArr, tmp)
	}

	return newArr

}
