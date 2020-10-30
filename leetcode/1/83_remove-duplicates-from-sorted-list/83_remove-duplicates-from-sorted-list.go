/*
link: https://leetcode.com/problems/remove-duplicates-from-sorted-list/

Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3

*/

package main

import "fmt"

func main() {

	l1 := []int{1, 2, 2, 2, 3, 4, 5, 6, 7}

	link1 := genList(l1)

	h := deleteDuplicates(link1)

	tail := h
	for tail != nil {
		fmt.Println(tail.Val)
		tail = tail.Next
	}
}

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

/*

提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	cur := head.Next

	for cur != nil {
		if cur.Val == pre.Val {
			// 这里注意，如果已经出现了重复数字，下次检测只移动cur
			next := cur.Next
			pre.Next = next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return head

}
