/*

link:https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

编写一个程序，找到两个单链表相交的起始节点。

如下面的两个链表：



在节点 c1 开始相交。



示例 1：



输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。


示例 2：



输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*

思路1： 先对两个链表进行反转，然后找到最后一个相同的节点. 最后要反转回来. bug: 但是因为两个链表有重复部分，对一个对反转会影响另外一个，本思路行不通
思路2: 遍历两个linklist, 得到两个长度的差值t，将较长的链表前移t步后与另外一个链表逐个比较，知道遇到相等节点
思路3: 将link2接到link1的尾部，形成环  再参考leetcode142,找到环上的第一个元素返回，最后再复原link2的尾部

这里按照思路2去实现，思路2提交结果如下：
Runtime: 32 ms, faster than 99.60% of Go online submissions for Intersection of Two Linked Lists.
Memory Usage: 8.6 MB, less than 15.60% of Go online submissions for Intersection of Two Linked Lists.

*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	len1 := 0
	len2 := 0

	l1 := headA
	l2 := headB

	for l1 != nil {
		len1++
		l1 = l1.Next
	}

	for l2 != nil {
		len2++
		l2 = l2.Next
	}

	l1 = headA
	l2 = headB

	if len1 < len2 {
		step := len2 - len1
		for step > 0 {
			l2 = l2.Next
			step--
		}
	} else if len1 > len2 {
		step := len1 - len2
		for step > 0 {
			l1 = l1.Next
			step--
		}
	}

	for l1 != nil && l2 != nil {
		if l1 == l2 {
			return l1
		} else {
			l1 = l1.Next
			l2 = l2.Next
		}
	}

	return nil

}

func printList(head *ListNode) {
	h := head
	for h != nil {
		fmt.Println("h.val:", h.Val)
		h = h.Next
	}
}
