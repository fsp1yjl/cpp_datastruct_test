/*

link: https://leetcode.com/problems/container-with-most-water/

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.


Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49

这里不推荐使用暴力解法1

*/

package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	area := maxArea(height)
	fmt.Println("area:", area)
}

/*
// 解法一 暴力解法： 依次求出宽度为1 到 len-1的长方形面积。 时间复杂度O(N*N)

Runtime: 708 ms, faster than 5.35% of Go online submissions for Container With Most Water.
Memory Usage: 5.8 MB, less than 81.69% of Go online submissions for Container With Most Water.
*/
func maxArea1(height []int) int {
	l := len(height)

	width := l - 1

	max := 0

	for i := 1; i <= width; i++ {
		for j := 0; j < l-i; j++ {
			left := j
			right := j + i

			lower := height[left]
			if height[right] < height[left] {
				lower = height[right]
			}

			v := lower * i
			if v > max {
				max = v
			}
		}
	}

	return max
}

/*
解法2:
初步左右指针分别指向0和len-1, 计算当前面积，并收缩高度较短的一侧的指针，因为这样才有增大面积的可能性

运行结果：
Runtime: 12 ms, faster than 91.13% of Go online submissions for Container With Most Water.
Memory Usage: 5.8 MB, less than 52.40% of Go online submissions for Container With Most Water.
*/
func maxArea(height []int) int {
	l := len(height)
	left := 0
	right := l - 1
	max := 0
	for left < right {
		cur := 0
		if height[left] < height[right] {
			cur = height[left] * (right - left)
			left++
		} else {
			cur = height[right] * (right - left)
			right--
		}
		if cur > max {
			max = cur
		}
	}

	return max
}
