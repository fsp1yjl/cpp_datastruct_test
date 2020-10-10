/*
link: https://leetcode-cn.com/problems/remove-element/

Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example 1:

Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums being 2.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,1,2,2,3,0,4,2], val = 2,

Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.

Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.

*/

package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))

	fmt.Println(nums)

}

/*
	原地删除特定元素，并返回剩余元素个数，原数组前面几位为为删除后的结果

	思路：
	每次发现有可删除元素，从数组最后寻找节点填补空缺


	提交结果：
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Remove Element.

*/
func removeElement(nums []int, val int) int {
	l := len(nums)

	res := l

	for i := 0; i < res; i++ { // 注意这里的终止条件，会随着删除元素而发生变化
		cur := nums[i]
		if cur == val {
			res--
			for j := res; j > i; j-- {
				if nums[j] != val {
					nums[i] = nums[j]
					break
				} else {
					res--
				}
			}
		}
	}

	return res
}
