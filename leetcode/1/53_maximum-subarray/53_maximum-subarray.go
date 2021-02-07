/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	// max, start, end := maxSubArray2(nums)

	// fmt.Println("max, start, end", max, start, end)

	max := maxSubArray(nums)

	fmt.Println("max ", max)
}

// 返回连续最大子数组的和，以及对应子数组的起始和结束位置
func maxSubArray2(nums []int) (int, int, int) {
	fmt.Println("nums:", nums)
	len := len(nums)
	if 0 == len {
		return 0, 0, 0
	}

	current_max := nums[0]
	current_start := 0
	current_end := 0

	sum_start := 0
	sum_end := 0

	sum_current := nums[0]
	// start := 0
	// end := 0

	for i := 1; i < len; i++ {

		if nums[i] > sum_current+nums[i] {
			sum_current = nums[i]
			sum_start = i
			sum_end = i
		} else {
			sum_current = sum_current + nums[i]
			sum_end = i
		}
		if sum_current > current_max {
			current_max = sum_current
			current_start = sum_start
			current_end = sum_end
		}
	}

	return current_max, current_start, current_end

}

// 仅返回连续最大子数组的和 时间复杂度o(n)
func maxSubArray1(nums []int) int {
	fmt.Println("nums:", nums)
	len := len(nums)
	if 0 == len {
		return 0
	}

	current_max := nums[0]

	sum_current := nums[0]

	for i := 1; i < len; i++ {

		if nums[i] > sum_current+nums[i] {
			sum_current = nums[i]
		} else {
			sum_current = sum_current + nums[i]
		}
		if sum_current > current_max {
			current_max = sum_current
		}
	}

	return current_max

}

// 分治法
func maxSubArray(nums []int) int {
	l := len(nums)

	var helper func([]int, int, int) int

	helper = func(nums []int, left int, right int) int {
		if left == right {
			return nums[left]
		}
		mid := (left + right) / 2

		// 计算包括mid的跨子区间的最大值
		mid_sum := nums[mid]
		max := mid_sum
		l := mid - 1
		r := mid + 1
		for r <= right {
			mid_sum = nums[r] + mid_sum
			if mid_sum > max {
				max = mid_sum
			}
			r++
		}
		mid_sum = max
		for l >= left {
			mid_sum = nums[l] + mid_sum
			if mid_sum > max {
				max = mid_sum
			}
			l--
		}

		// 计算左区间最大值
		if mid-1 >= left {
			left_max := helper(nums, left, mid-1)
			if max < left_max {
				max = left_max
			}
		}

		// 计算右区间最大值
		if mid+1 <= right {
			right_max := helper(nums, mid+1, right)
			if max < right_max {
				max = right_max
			}

		}
		return max
	}

	return helper(nums, 0, l-1)

}
