/*
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5

 有序数组变二叉搜索树

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
递归处理，每次取中间元素作为根节点，并处理左侧和右侧

提交结果：
Runtime: 100 ms, faster than 87.18% of Go online submissions for Convert Sorted Array to Binary Search Tree.
Memory Usage: 238.6 MB, less than 5.41% of Go online submissions for Convert Sorted Array to Binary Search Tree.
*/
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)

	if 0 == l {
		return nil
	}

	return helper(nums, 0, l-1)

}

func helper(nums []int, start int, end int) *TreeNode {

	// if start > end {
	// 	return nil
	// }

	mid := (start + end) / 2

	node := &TreeNode{
		Val: nums[mid],
	}

	if start == end {
		return node
	}

	if mid-1 >= start {
		node.Left = helper(nums, start, mid-1)
	}
	if mid+1 <= end {
		node.Right = helper(nums, mid+1, end)
	}

	return node

}
