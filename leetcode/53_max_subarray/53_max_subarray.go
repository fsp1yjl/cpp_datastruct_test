package main

import "fmt"

/*
Given an integer array nums,
find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

*/
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

// 仅返回连续最大子数组的和
func maxSubArray(nums []int) int {
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
