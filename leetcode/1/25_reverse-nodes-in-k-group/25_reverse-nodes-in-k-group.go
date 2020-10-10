/*
link: https://leetcode.com/problems/reverse-nodes-in-k-group/


Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

Follow up:

Could you solve the problem in O(1) extra memory space?
You may not alter the values in the list's nodes, only nodes itself may be changed.
 

Example 1:


Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
Example 2:


Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
Example 3:

Input: head = [1,2,3,4,5], k = 1
Output: [1,2,3,4,5]
Example 4:

Input: head = [1], k = 1
Output: [1]

*/


package main
import "fmt"

func main() {

	l1 := []int{1, 2, 3, 4, 5, 6, 7}

	link1 := genList(l1)

	// tail := link1
	// for tail != nil {
	// 	fmt.Println(tail.Val)
	// 	tail = tail.Next
	// }
	h := reverseKGroup(link1, 4)

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
Runtime: 4 ms, faster than 98.84% of Go online submissions for Reverse Nodes in k-Group.
Memory Usage: 3.7 MB, less than 6.36% of Go online submissions for Reverse Nodes in k-Group.

思路： 将问题转化为递归， head指向子问题的解， 后面的进行反转
*/

func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 1 {
		return head 
	}
	
	tail1 := head 
	n := 1
	for n < k {
		if tail1.Next == nil {
			tail1 = nil 
			break
		} else  {
			n++
			tail1 = tail1.Next
		}
	}

	// fmt.Println("tail1:", tail1.Val)
	
	// 如果tail1为空，表示链表长度小于k
	if tail1 == nil {
		return head 
	}


	// kHead := head 
	// kTail := tail1 

	return reverk(head, k)


}

func reverk(head *ListNode, k int) *ListNode {

	tail1 := head 
	n := 1

	for n < k {
		if tail1.Next == nil {
			tail1 = nil 
			break
		} else  {
			n++
			tail1 = tail1.Next
		}
	}

	

	if tail1 == nil {
		return head 
	} else {
		nextHead := tail1.Next
		n = 1
		pre := head 
		cur := pre.Next
		
		for n < k {
			next := cur.Next
			cur.Next = pre 

			pre = cur 
			cur = next 
			n++
		}
		
		if nextHead != nil {
			head.Next = reverk(nextHead, k)
		} else {
			head.Next = nil 
		}
		return tail1 

	}
}

	




