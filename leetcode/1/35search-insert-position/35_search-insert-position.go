/*
link: https://leetcode.com/problems/search-insert-position/

参考： https://www.cnblogs.com/grandyang/p/6854825.html

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.



Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4

*/

package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 0))
}

// 时间复杂度O(n)的解法
func searchInsert1(nums []int, target int) int {
	l := len(nums)
	if 0 == l {
		return 0
	}

	left := 0
	// right := l-1
	for nums[left] < target {
		left++
		if left > l-1 {
			break
		}
	}

	return left

}

/*
//时间复杂度log(n)的解法，不断收缩，right表示大于等于target的第一个位置

Runtime: 0 ms, faster than 100.00% of Go online submissions for Search Insert Position.
Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Search Insert Position.
*/
func searchInsert(nums []int, target int) int {
	l := len(nums)
	if 0 == l {
		return 0
	}

	left := 0
	right := l // 注意，这里是l,不是l-1. 因为插入后，元素个数加1
	// right := l-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right

}
