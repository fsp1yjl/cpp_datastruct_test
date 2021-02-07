package l84

import (
	"fmt"
)

// 暴力方法，求出每个高度对应的最大面积
/*
Runtime: 564 ms, faster than 7.86% of Go online submissions for Largest Rectangle in Histogram.
Memory Usage: 4.7 MB, less than 87.86% of Go online submissions for Largest Rectangle in Histogram.

*/
func largestRectangleArea1(heights []int) int {
	l := len(heights)

	max := 0
	if l == 0 {
		return max
	}

	for i := 0; i < l; i++ {
		left := i
		right := i
		h := heights[i]
		for left >= 1 {

			if heights[left-1] >= h {
				left--
			} else {
				break
			}
		}

		for right <= l-2 {

			if heights[right+1] >= h {
				right++
			} else {
				break
			}
		}

		area := (right - left + 1) * h
		if area > max {
			max = area
		}
	}
	fmt.Println("max:", max)
	return max

}

/*

二分法，先求出包括mid的最大面积，与(left,mid-1), (mid+1, right)两子区间的最大面积做比较
因为最大面积要么在左右子区间，要么为包括mid的左右延伸区域
提交结果：
Runtime: 12 ms, faster than 34.29% of Go online submissions for Largest Rectangle in Histogram.
Memory Usage: 4.5 MB, less than 88.57% of Go online submissions for Largest Rectangle in Histogram.
*/
func largestRectangleArea(heights []int) int {
	l := len(heights)

	max := 0
	if l == 0 {
		return max
	}

	var helper func(int, int) int

	helper = func(left int, right int) int {
		if left == right {
			return heights[left]
		}
		mid := (left + right) / 2

		h := heights[mid]

		//step1: 先计算中间面积
		max := h
		l := mid
		r := mid
		min_height := h
		for l > left && r < right {
			if heights[l-1] < heights[r+1] {
				r++
				if heights[r] < min_height {
					min_height = heights[r]
				}
			} else {
				l--
				if heights[l] < min_height {
					min_height = heights[l]
				}
			}
			tmp := (r - l + 1) * min_height
			if tmp > max {
				max = tmp
			}

		}
		for l > left {
			l--
			if heights[l] < min_height {
				min_height = heights[l]
			}
			tmp := (r - l + 1) * min_height
			if tmp > max {
				max = tmp
			}
		}

		for r < right {
			r++
			if heights[r] < min_height {
				min_height = heights[r]
			}
			tmp := (r - l + 1) * min_height
			if tmp > max {
				max = tmp
			}
		}

		// fmt.Println("mid area:", max, mid)

		// step2 计算左边最大面积
		if mid-1 >= left {
			areaLeft := helper(left, mid-1)
			if areaLeft > max {
				max = areaLeft
			}
		}

		// step3 计算右边最大面积
		if mid+1 <= right {
			areaRight := helper(mid+1, right)
			if areaRight > max {
				max = areaRight
			}
		}

		return max
	}

	return helper(0, l-1)

}
