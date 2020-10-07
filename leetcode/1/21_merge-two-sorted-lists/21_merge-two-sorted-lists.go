/*
link: https://leetcode.com/problems/merge-two-sorted-lists/
难度： easy

Merge two sorted linked lists and return it as a new sorted list.
The new list should be made by splicing together the nodes of the first two lists.



Example 1:


Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: l1 = [], l2 = []
Output: []
Example 3:

Input: l1 = [], l2 = [0]
Output: [0]
*/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*

 提交结果：
 Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.6 MB, less than 24.67% of Go online submissions for Merge Two Sorted Lists.

*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	head = &ListNode{
		Val:  0,
		Next: nil,
	}
	tail := head

	left := l1
	right := l2
	for left != nil && right != nil {
		if left.Val <= right.Val {
			node := &ListNode{
				Val:  left.Val,
				Next: nil,
			}
			tail.Next = node
			tail = tail.Next
			left = left.Next
		} else {
			node := &ListNode{
				Val:  right.Val,
				Next: nil,
			}
			tail.Next = node
			tail = tail.Next
			right = right.Next
		}
	}

	for left != nil {
		node := &ListNode{
			Val:  left.Val,
			Next: nil,
		}
		tail.Next = node
		tail = tail.Next
		left = left.Next
	}

	for right != nil {
		node := &ListNode{
			Val:  right.Val,
			Next: nil,
		}
		tail.Next = node
		tail = tail.Next
		right = right.Next
	}

	if head.Next != nil {
		return head.Next
	} else {
		return nil
	}

}
