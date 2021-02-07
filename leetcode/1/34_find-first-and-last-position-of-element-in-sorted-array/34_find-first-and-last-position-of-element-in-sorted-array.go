/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}

	fmt.Println(searchRange(nums, 8))
}

/*

思路： 先二分查找找到等于target的一个点，然后双指针，左右移动，找到左右边界

提交结果：
Runtime: 8 ms, faster than 81.84% of Go online submissions for Find First and Last Position of Element in Sorted Array.
Memory Usage: 4.1 MB, less than 100.00% of Go online submissions for Find First and Last Position of Element in Sorted Array.



*/
func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l < 1 {
		return []int{-1, -1}
	}
	left := 0
	right := l - 1
	tt := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			tt = mid
			break
		}

		if nums[mid] > target {
			right = mid - 1
			continue
		}

		if nums[mid] < target {
			left = mid + 1
			continue
		}
	}

	if tt == -1 {
		return []int{-1, -1}
	}

	left = tt
	right = tt
	for left >= 1 {
		if target == nums[left-1] {
			left--
		} else {
			break
		}
	}

	for right < l-1 {
		if target == nums[right+1] {
			right++
		} else {
			break
		}
	}

	return []int{left, right}

}
