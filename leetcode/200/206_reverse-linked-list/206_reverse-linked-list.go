/*
link :https://leetcode.com/problems/reverse-linked-list/

链表反转， 要求提供递归和非递归两种方法

*/

package main

import "fmt"

func main() {
	nums := []int{1, 2}
	head := genList(nums)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	reverseList(head)

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

//递归方式
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next

	newHead := reverseList(next)
	next.Next = head
	head.Next = nil // bug1, 这里需要将头的next置nil, 防止形成环

	return newHead
}

//非递归方式

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	cur := head.Next

	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	head.Next = nil //防止成环
	return pre

}
