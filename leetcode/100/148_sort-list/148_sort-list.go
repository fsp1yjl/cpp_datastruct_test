/*
link : https://leetcode.com/problems/sort-list/


Given the head of a linked list, return the list after sorting it in ascending order.

Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
 要求时间复杂度 nlog(n)

Example 1:


Input: head = [4,2,1,3]
Output: [1,2,3,4]
*/

package main

func main() {
	l := genList([]int{4, 2, 1, 3, 5})
	sortList(l)

}

type ListNode struct {
	Val  int
	Next *ListNode
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

/*
提交结果：
Runtime: 92 ms, faster than 20.77% of Go online submissions for Sort List.
Memory Usage: 8.3 MB, less than 10.26% of Go online submissions for Sort List.

*/
func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	//快慢指针找到中点， 对左右排序，然后进行归并

	quick := head
	slow := head
	for quick != nil && quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}

	// fmt.Println("middle:", slow.Val)
	middle := slow
	left := head //注意，这里已经对链表进行链截取
	right := middle.Next
	middle.Next = nil

	var sortedHead *ListNode
	var sortedTail *ListNode

	if right == nil {
		return head
	} else {
		l := sortList(left)
		r := sortList(right)

		head := l
		for head != nil {
			// fmt.Println("ll:", head.Val)
			head = head.Next
		}

		head = r
		for head != nil {
			// fmt.Println("rr:", head.Val)
			head = head.Next
		}

		//进行一次归并
		for l != nil && r != nil {
			n := &ListNode{}
			if l.Val < r.Val {
				n.Val = l.Val
				l = l.Next

			} else {
				n.Val = r.Val
				r = r.Next
			}

			if sortedHead == nil {
				sortedHead = n
			}

			if sortedTail == nil {
				sortedTail = n
			} else {
				sortedTail.Next = n
				sortedTail = n
			}
		}

		for l != nil {
			n := &ListNode{
				Val: l.Val,
			}
			l = l.Next
			sortedTail.Next = n
			sortedTail = n
		}

		for r != nil {
			n := &ListNode{
				Val: r.Val,
			}
			r = r.Next
			sortedTail.Next = n
			sortedTail = n
		}

		return sortedHead
	}
}
