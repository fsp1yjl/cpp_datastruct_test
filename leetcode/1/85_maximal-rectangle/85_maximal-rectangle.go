/*

给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。


思路，将矩阵转化为柱形图，然后求柱形图的最大矩形面积. 求柱形图的最大矩形面积使用leetcode 84的解法


提交结果：
Runtime: 4 ms, faster than 69.89% of Go online submissions for Maximal Rectangle.
Memory Usage: 4.3 MB, less than 58.06% of Go online submissions for Maximal Rectangle.
*/

package l85

func maximalRectangle(matrix [][]byte) int {

	m := len(matrix)
	if 0 == m {
		return 0
	}
	n := len(matrix[0])

	if 0 == n {
		return 0
	}
	left := 0
	maxArea := 0
	for left < n {
		heights := make([]int, m)
		min_sum := n - left
		// fmt.Println("----------LEFT:", left)
		for i := 0; i < m; i++ {
			heights[i] = n - left

			sum := 0
			for j := left; j < n; j++ {
				if '1' == matrix[i][j] {
					sum++
					continue
				} else {
					if sum < min_sum {
						min_sum = sum
					}
					heights[i] = sum
					break
				}
			}

		}

		// fmt.Println("heights:", heights)

		tmpMaxArea := largestRectangleArea(heights)
		if tmpMaxArea > maxArea {
			maxArea = tmpMaxArea
		}

		// fmt.Println("tmparea:", tmpMaxArea)
		left += min_sum + 1
	}

	return maxArea
}

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
