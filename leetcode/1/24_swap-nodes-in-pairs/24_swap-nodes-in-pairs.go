/*
link: https://leetcode.com/problems/swap-nodes-in-pairs/

Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes. Only nodes itself may be changed.

example1:
Input: head = [1,2,3,4]
Output: [2,1,4,3]

Example 2:
Input: head = []
Output: []

Example 3:
Input: head = [1]
Output: [1]

*/

package main

import "fmt"

func main() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7}

	link1 := genList(l1)

	h := swapPairs(link1)

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

tail1,tail2 双指针，每次移动两个距离
1->2->3->4->5, 则tail1依次为1，3，5， tail2为2，4

提交情况：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
Memory Usage: 2.1 MB, less than 74.26% of Go online submissions for Swap Nodes in Pairs.
Next challenges:
Reverse Nodes in k-Group


*/
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	head2 := head.Next

	tail1 := head
	tail2 := head.Next

	for tail2 != nil && tail2.Next != nil && tail2.Next.Next != nil {
		next2 := tail2.Next.Next

		next1 := tail1.Next.Next

		tail2.Next = tail1
		tail1.Next = next2
		tail1 = next1
		tail2 = next2

	}

	// fmt.Println("tail1:", tail1.Val)
	// fmt.Println("tail2:", tail2.Val)

	// 注意这里的最终情况，要么tail2为最后一个元素，要么tail2.next 为最后一个元素
	if tail2.Next == nil {
		tail2.Next = tail1
		tail1.Next = nil
	} else if tail2.Next.Next == nil {
		last := tail2.Next
		tail2.Next = tail1
		tail1.Next = last
	}

	return head2
}
