/*
link: https://leetcode.com/problems/longest-continuous-increasing-subsequence/

Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence (i.e. subarray). The subsequence must be strictly increasing.

A continuous increasing subsequence is defined by two indices l and r (l < r) such that it is [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].



Example 1:

Input: nums = [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
Even though [1,3,5,7] is an increasing subsequence, it is not continuous as elements 5 and 7 are separated by element
4.
Example 2:

Input: nums = [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2] with length 1. Note that it must be strictly
increasing.

*/

package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 7, 4, 5, 9, 10, 12}
	fmt.Println(findLengthOfLCIS(nums))
}

/*
easy 不解释：
Runtime: 8 ms, faster than 87.50% of Go online submissions for Longest Continuous Increasing Subsequence.
Memory Usage: 4.5 MB, less than 22.73% of Go online submissions for Longest Continuous Increasing Subsequence.

*/

func findLengthOfLCIS(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}

	max := 1
	tmp := 1

	pre := nums[0]

	for i := 1; i < l; i++ {
		if nums[i] > pre {
			tmp++
		} else {
			tmp = 1
		}
		if tmp > max {
			max = tmp
		}
		pre = nums[i]
	}
	return max
}
