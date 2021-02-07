package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1}

	nums2 := []int{5, 6, 7, 8, 9, 10}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	k1 := (l1 + l2 + 1) / 2
	k2 := (l1 + l2 + 2) / 2

	if k1 == k2 {
		return float64(helper(nums1, 0, nums2, 0, k1))
	} else {
		sum := helper(nums1, 0, nums2, 0, k1) + helper(nums1, 0, nums2, 0, k2)
		return float64(sum) / 2
	}

}

func helper(nums1 []int, s1 int, nums2 []int, s2 int, k int) float64 {
	if s1 >= len(nums1) {
		//表示nums1已经无可利用元素,直接取nums2的第s2+k个元素返回
		return float64(nums2[s2+k-1])
	}
	if s2 >= len(nums2) {
		return float64(nums1[s1+k-1])
	}

	// 如果返回第1个元素，则需要单独处理，防止进入死循环
	if k == 1 {
		if nums1[s1] < nums2[s2] {
			return float64(nums1[s1])
		} else {
			return float64(nums2[s2])
		}
	}
	left := math.MaxInt64
	right := math.MaxInt64

	// 这里需要考虑向前移动k/2越界的情况
	if (s1 + k/2 - 1) < len(nums1) {
		left = nums1[s1+k/2-1]
	}
	if (s2 + k/2 - 1) < len(nums2) {
		right = nums2[s2+k/2-1]
	}

	if left < right {
		//表示nums1 有k/2个元素可以被放入中位数左侧
		return helper(nums1, s1+k/2, nums2, s2, k-k/2)
	} else {
		return helper(nums1, s1, nums2, s2+k/2, k-k/2)
	}

}
