/*
link:
https://leetcode.com/problems/4sum/

Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2, 2, -4}
	fmt.Println(fourSum(nums, 0))
}

/*
 和3sum 类似，多一层循环

 注意去除重复解的方法

*/
func fourSum(nums []int, target int) [][]int {

	res := [][]int{}

	l := len(nums)

	if l < 4 {
		return res
	}
	// 数组排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// fmt.Println(nums)

	for i := 0; i < l-3; i++ {
		// fmt.Println("here i:", i, nums[i])
		if i > 0 && nums[i] == nums[i-1] { // 四个元素均用同样的方式过滤重复解
			continue
		}
		for j := i + 1; j < l-2; j++ {
			// fmt.Println("here j:", j, nums[j])

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			t := target - nums[i] - nums[j]
			left := j + 1
			right := l - 1
			for left < right {
				// fmt.Println("left:, right:", left, right)
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right < l-1 && nums[right] == nums[right+1] {
					right--
					continue
				}

				if t > nums[left]+nums[right] {
					left++
				} else if t < nums[left]+nums[right] {
					right--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				}
			}
		}
	}

	return res

}
