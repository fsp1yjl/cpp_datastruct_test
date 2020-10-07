/*
link: https://leetcode.com/problems/merge-k-sorted-lists/


You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.



Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []

*/

package main

import "fmt"

func genList(nums []int) *ListNode {
	l := len(nums)
	if l == 0 {
		return nil
	}

	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}

	tail := head

	for i := 1; i < l; i++ {
		node := &ListNode{
			Val:  nums[i],
			Next: nil,
		}

		tail.Next = node
		tail = tail.Next
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := []int{1, 2, 3}
	l2 := []int{1, 4, 5}

	l3 := []int{1, 3, 9}

	link1 := genList(l1)
	link2 := genList(l2)
	link3 := genList(l3)

	lists := []*ListNode{link1, link2, link3}

	l := mergeKLists(lists)

	tail := l
	for tail != nil {
		fmt.Println(tail.Val)
		tail = tail.Next
	}

}

/*
这里并不是一个优雅的解法，思路是将数组进行多次2个链表的归并排序
结果也只是勉强通过

结果：
Runtime: 1992 ms, faster than 5.46% of Go online submissions for Merge k Sorted Lists.
Memory Usage: 6.8 MB, less than 5.05% of Go online submissions for Merge k Sorted Lists.

*/
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)

	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}

	merged := lists[0]
	second := 1

	for second < l {
		// fmt.Println("ddd")
		merged = mergeTwoLists(merged, lists[second])
		second++
	}

	return merged

}

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
