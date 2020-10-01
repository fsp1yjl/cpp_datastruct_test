/*
link:
https://leetcode.com/problems/3sum-closest/
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.



Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).


Constraints:

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4

*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}

/*

提交结果：
Runtime: 4 ms, faster than 97.75% of Go online submissions for 3Sum Closest.
Memory Usage: 2.7 MB, less than 68.54% of Go online submissions for 3Sum Closest.

思路：先对数组排序，然后确定最左侧后，从后续数组中选择mid和right
*/
func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}

	//数组排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < l-2; i++ {
		mid := i + 1
		right := l - 1
		for mid < right {
			sum := nums[i] + nums[mid] + nums[right]
			if sum > target { // 如果和大于target，则移动右侧，去缩小sum; 否则移动mid
				right--
			} else {
				mid++
			}

			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
			}
		}

	}

	return res

}
