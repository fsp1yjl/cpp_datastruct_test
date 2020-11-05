/*
判断一个链表是否有环

使用快慢指针

参考：
https://www.youtube.com/watch?v=bxCb37nLXWM

假设有环，（2t-p+1)%q , (t-p+1)%q 分别表示t次后，落在环内节点信息。
  p表示非环内节点个数， t表示走的步数，q表示环内节点个数，当t为q的倍数的时候，一定成立。
  即当存在环的时候，使用快慢指针，经过q步后，一定会出现碰头的情况

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
Runtime: 4 ms, faster than 96.69% of Go online submissions for Linked List Cycle.
Memory Usage: 3.8 MB, less than 19.46% of Go online submissions for Linked List Cycle.

*/

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head.Next
	quick := head.Next.Next

	for quick != nil && quick.Next != nil {
		if slow == quick {
			return true
		} else {
			slow = slow.Next
			quick = quick.Next.Next
		}
	}

	return false

}