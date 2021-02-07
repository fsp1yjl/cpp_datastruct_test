/*

link : https://leetcode-cn.com/problems/linked-list-cycle-ii/

题目说明：
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

进阶：

你是否可以使用 O(1) 空间解决此题？




*/

package main

import "fmt"

func main() {
	l := genCircleList([]int{3,2,0, -4}, 1)

	fmt.Println(detectCycle(l))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func genCircleList(nums []int, c int) *ListNode {
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

	circle := head 
	for c > 0 {
		circle = circle.Next
		c--
	}
	tail.Next = circle

	return head
}

/*
提交结果：
Runtime: 4 ms, faster than 96.45% of Go online submissions for Linked List Cycle II.
Memory Usage: 3.8 MB, less than 31.91% of Go online submissions for Linked List Cycle II.

*/
func detectCycle(head *ListNode) *ListNode {
	//先通过快慢指针判断是否有环，如果有，先确定环的长度
	//假设存在环，环前元素个数为p, 换上元素个数为q, 则相遇时快走2t步，慢走t步，
	// （t-p+1）% q 和 （2t-p+1)%q  相等，此时当t=q时，解成立
	
	if head == nil {
		return nil
	}

	quick := head 
	slow := head 
	circleFlag := false
	cnt := 0
	for quick.Next !=nil && quick.Next.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		cnt++
		if  quick == slow {
			circleFlag = true 
			//表示存在环，此时cnt即表示环上元素数量
			break
		}
	}

	if !circleFlag  {
		return nil
	}

	// fmt.Println("cnt:",cnt)
	
	// 确定了环上都元素个数为cnt后， 利用一个相差cnt步的指针对，第一个相等的节点就是环的第一个节点
	slow = head 
	quick = head 
	p := 0
	for cnt > 0 {
		quick = quick.Next
		cnt--
	}
	for slow != quick {
		p++
		slow = slow.Next
		quick = quick.Next
	}

	return slow
}