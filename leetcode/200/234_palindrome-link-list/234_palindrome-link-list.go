/*
leetcode 9 的升级版问题
link: https://leetcode.com/problems/palindrome-linked-list/

Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true
Follow up:
Could you do it in O(n) time and O(1) space?


分析： 题目要求是时间复杂度O(n), 空间复杂度O(1)

运行结果：
Runtime: 12 ms, faster than 93.35% of Go online submissions for Palindrome Linked List.
Memory Usage: 6 MB, less than 92.02% of Go online submissions for Palindrome Linked List.
Next challenges:
*/

package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
	}

	a := &ListNode{
		Val: 2,
	}

	a.Next = &ListNode{
		Val: 2,
	}
	a.Next.Next = &ListNode{
		Val: 1,
	}

	head.Next = a

	fmt.Println(isPalindrome(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 接下来考虑使用快慢指针，找到链表的中点，对中点之后对节点进行反转操作。
func isPalindrome(head *ListNode) bool {

	// bug1: 要考虑链表为空的情况
	if head == nil {
		return true
	}
	//step1 :使用快慢双指针，进行链表后半部分翻转
	slow := head
	quick := head
	rightLast := head
	leftLast := head
	for quick != nil {
		// fmt.Println("222")

		if quick.Next == nil {
			rightLast = quick
			quick = nil

			// fmt.Println("llllll")
		} else {
			rightLast = quick.Next
			quick = quick.Next.Next
		}
		leftLast = slow
		slow = slow.Next
		// fmt.Println("leftlast:", leftLast.Val)
		// fmt.Println("rigthlast:", rightLast.Val)
	}
	// fmt.Println("leftlast:", leftLast.Val)
	// fmt.Println("rigthlast:", rightLast.Val)

	tmp := leftLast.Next
	// bug2 要考虑leftLast.Next为空造成tmp.Next报错的情况
	tmpNext := tmp
	if tmp != nil {
		tmpNext = tmp.Next
	}

	tmpPre := leftLast
	for tmp != nil {

		tmp.Next = tmpPre
		tmpPre = tmp
		tmp = tmpNext
		if tmpNext != nil {
			tmpNext = tmpNext.Next

		} else {
			tmpNext = nil
		}

	}

	// bug3 要考虑leftLast为最后一个元素对情况， 否则会造成环状链表
	if leftLast.Next != nil {
		leftLast.Next.Next = nil
		leftLast.Next = rightLast
	}

	l := head
	r := leftLast.Next

	for r != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
		// fmt.Println("ddd")
	}

	return true
}

/*

// 最开始的想法是把linklist的值按大端和小端分别计算存储，然后比对
// 但是这个方案有一个问题是，节点过多的时候，会造成数字溢出
func isPalindrome_1(head *ListNode) bool {
	littleEnd := 0 // 低字节放在内存低位
	bigEnd := 0    // 低字节放内存高位

	bigWeight := 1
	node := head
	for node != nil {
		littleEnd = littleEnd*10 + node.Val
		bigEnd = bigEnd + node.Val*bigWeight

		bigWeight *= 10

		node = node.Next
	}

	fmt.Println("little:", littleEnd)

	fmt.Println("BIG:", bigEnd)
	if bigEnd == littleEnd {
		return true
	} else {
		return false
	}
}

*/
