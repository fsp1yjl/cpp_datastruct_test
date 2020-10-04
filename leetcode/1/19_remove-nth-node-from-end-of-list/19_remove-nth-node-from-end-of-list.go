/*
link: https://leetcode.com/problems/remove-nth-node-from-end-of-list/

Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?

思考：
首先想到的是先对链表进行反转，然后遍历反转后对连标
但是follow up的要求是最多进行依次反转
想到是否可以使用快慢指针进行处理，慢指针表示要删除节点的前一个，快指针为慢指针向后移动n的节点
最终快指针为最后的时候， 慢指针的一下即为等待删除的元素
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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
Memory Usage: 2.2 MB, less than 99.79% of Go online submissions for Remove Nth Node From End of List.
Next challenges:
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left := head
	right := head

	for i := 0; i < n; i++ {
		if right != nil {
			right = right.Next
		} else {
			right = nil
		}
	}

	// 如果首次right即为nil, 表示传入的n超过了链表长度
	// bug1, 要考虑链表表头元素被删除的情况处理
	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	// next := left.Next.Next

	left.Next = left.Next.Next

	return head
}
