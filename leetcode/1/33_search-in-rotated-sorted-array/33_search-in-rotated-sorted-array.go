/*

link: https://leetcode.com/problems/search-in-rotated-sorted-array/
You are given an integer array nums sorted in ascending order, and an integer target.

Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

If target is found in the array return its index, otherwise, return -1.



Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
Example 3:

Input: nums = [1], target = 0
Output: -1





*/

package main

import "fmt"

func main() {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}

	nums := []int{1}

	fmt.Println(search(nums, 0))
}

/*
思路，先找到发生反转的位置，然后进行二分搜索

提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Search in Rotated Sorted Array.
*/
func search(nums []int, target int) int {
	//先判断是否发生过部分翻转

	l := len(nums)

	if 0 == l {
		return -1
	}

	rotateFlag := false
	index := -1
	pre := nums[0]
	for i := 1; i < l; i++ {
		if nums[i] < pre {
			rotateFlag = true
			index = i
			break
		} else {
			pre = nums[i]
		}
	}

	if !rotateFlag {
		return binarySearch(nums, 0, l-1, target)
	} else {
		right_min := nums[index]
		right_max := nums[l-1]
		left_min := nums[0]
		left_max := nums[index-1]

		if target >= left_min && target <= left_max {
			return binarySearch(nums, 0, index-1, target)
		} else if target >= right_min && target <= right_max {
			return binarySearch(nums, index, l-1, target)
		} else {
			return -1
		}
	}

}

func binarySearch(nums []int, left int, right int, target int) int {
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
