/*
link: https://leetcode.com/problems/trapping-rain-water/

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6

*/

package main

import "fmt"

func main() {
	height := []int{4, 4, 4, 7, 1, 0}

	sum := trap(height)
	fmt.Println(sum)
}

/*
思路：
将数组切割为能盛水的多个区间，先找到区间的左侧，再确定区间的右侧。
左侧确定的方法：递增区间的最大数即为左侧
右侧确定的方法：
	左侧向后搜索，记录临时最大高度，
	如果高度小于左侧高度，则继续向右搜索，找到一个更接近左侧高度的元素
	如果遇到高度大于左侧的元素，则停止搜索，一个桶区间确认完成。进行这个桶区间盛水量计算
	之后可以从右侧开始，进行下一个桶区间的判断，并计算
运行结果：
Runtime: 4 ms, faster than 72.26% of Go online submissions for Trapping Rain Water.
Memory Usage: 2.8 MB, less than 99.70% of Go online submissions for Trapping Rain Water.

*/
func trap(height []int) int {
	l := len(height)

	// bug1 ：需要处理数组小于2或者为空的情况
	if l <= 2 {
		return 0
	}

	sum := 0
	leftIndex := 0
	curLeft := height[0]
	i := leftIndex
	for i < l-1 {
		if height[i] >= curLeft {
			curLeft = height[i]
			leftIndex = i // bug2: 不要忘记每次更新左桶索引
			i++
		} else {
			// leftFlag = true
			curRight := height[i]
			rightIndex := i
			min := height[i]         //这里的min用来记录桶区间的最小值，防止出现单调递减计算结果为负的情况
			for j := i; j < l; j++ { // 这个内侧循环是为了找到left右侧，最接近左桶板子或者大于左板子的第一个值

				if height[i] < min {
					min = height[i]
				}
				if height[j] > curRight {
					curRight = height[j]
					rightIndex = j
					//如果右侧大于左桶板子，直接退出循环； 否则继续向右寻找最接近左板的值
					if height[j] >= curLeft {
						break

					}
				}
			}

			h := curLeft
			if curRight < h {
				h = curRight
			}

			// fmt.Println("left:,right:", curLeft, curRight)
			// fmt.Println("left:,right: index", leftIndex, rightIndex)

			if h >= min {
				for k := leftIndex + 1; k < rightIndex; k++ {
					sum += h - height[k]
				}
			}

			leftIndex = rightIndex
			curLeft = height[rightIndex]
			i = rightIndex

		}

	}

	return sum
}
