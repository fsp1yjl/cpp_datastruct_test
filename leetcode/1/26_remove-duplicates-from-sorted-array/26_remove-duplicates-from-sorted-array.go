/*
link:
https://leetcode.com/problems/remove-duplicates-from-sorted-array


Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
*/

package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums))

	fmt.Println(nums)
}

/*
提交结果：
Runtime: 8 ms, faster than 89.65% of Go online submissions for Remove Duplicates from Sorted Array.
Memory Usage: 4.6 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted Array.

*/
func removeDuplicates(nums []int) int {
	l := len(nums)

	if l <= 1 {
		return l
	}
	index := 0
	preNum := nums[0]

	for i := 1; i < l; i++ {
		cur := nums[i]
		if cur != preNum {
			index++
			nums[index] = cur
			preNum = cur
		}
	}

	return index + 1
}
