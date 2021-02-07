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

import (
	"container/heap"
	"fmt"
)

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
func mergeKLists_1(lists []*ListNode) *ListNode {
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

/*
网上看到对构建最小堆去解题堆思路：
link: https://segmentfault.com/a/1190000023327265
https://ieevee.com/tech/2018/01/29/go-heap.html



*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var fhead = &ListNode{}
	h := &IntHeap{}
	heap.Init(h)

	for i := range lists {
		if lists[i] == nil {
			continue
		}
		heap.Push(h, node{Val: lists[i].Val, ListIdx: i})
		lists[i] = lists[i].Next
	}

	curr := fhead
	// container/heap 会自动帮你做 heapify
	// 每次从小顶堆 pop 之后要从原来的 list 中再取出一个 node
	// 如果原来的 list 为空就继续 pop
	// 这也就是为什么不直接存 ListNode.Val 到 heap 里
	// 为了保存当前从堆 pop 出来的值来自哪个 list 专门声明了 node 类型
	// 另外，可能你会对为什么一定要从 pop 出来的值原来的 list 再取一个 node
	// 原因在于，本地的 merge 逻辑就是随时都在堆里放当前所有 list 的最小值
	// 每次都从原 list 取一个值放进堆里，这个值被 push 进去之后可能不会成为堆顶
	// 但是成为一旦原 list 的下一个值可能成为堆顶或者可能在其被 push 进堆之前成为堆顶
	// 就会造成最终 merge 排序的乱序
	// 所以最稳健和简单的做法就是每次都从原 list 找下一个值 push 到堆里
	for h.Len() > 0 {
		n := heap.Pop(h).(node)
		curr.Next = &ListNode{Val: n.Val}
		if lists[n.ListIdx] != nil {
			heap.Push(h, node{
				Val:     lists[n.ListIdx].Val,
				ListIdx: n.ListIdx})
			lists[n.ListIdx] = lists[n.ListIdx].Next
		}
		curr = curr.Next
	}
	return fhead.Next
}

type node struct {
	Val     int
	ListIdx int
}

type IntHeap []node

func (h IntHeap) Len() int { return len(h) }

// Less 递增表示构成小顶堆
func (h IntHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(node))
}

func (h *IntHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}
