/*

给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差

示例：

输入：{1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
输出： 3，即数值对(11, 8)
提示：

1 <= a.length, b.length <= 100000
-2147483648 <= a[i], b[i] <= 2147483647
正确结果在区间[-2147483648, 2147483647]内

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/smallest-difference-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 15, 11, 2}
	b := []int{23, 127, 235, 19, 8, 9, 12}

	smallestDifference(a, b)
}

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	pre_flag := ""
	pre_left := a[0]
	pre_right := b[0]

	left := 0
	right := 0
	if a[0] == b[0] {
		return 0
	}
	min := 0
	if a[0] < b[0] {
		pre_flag = "left"
		min = b[0] - a[0]
		left++
	}
	if a[0] > b[0] {
		pre_flag = "right"
		min = a[0] - b[0]
		right++
	}

	l1 := len(a)
	l2 := len(b)

	// fmt.Println("l1::", l1)
	for left < l1 && right < l2 {
		if a[left] == b[right] {
			return 0
		}

		if a[left] < b[right] {
			//上一个被归并的元素在a中，表示连续两个a中元素小于b的当前元素
			if "left" == pre_flag {
				tmp_min := b[right] - a[left]
				if tmp_min < min {
					min = tmp_min
				}
				pre_left = a[left]
			}
			//表示上一个被归并的元素在b中，表示出现交替归并
			if "right" == pre_flag {
				tmp_min := a[left] - pre_right
				if tmp_min < min {
					min = tmp_min
				}
				pre_left = a[left]
			}
			left++
			pre_flag = "left"
			continue
		}

		// fmt.Println("left:", left)
		if a[left] > b[right] {

			//上一个被归并的元素在b中，表示连续两个b中元素小于a的当前元素
			if "right" == pre_flag {
				tmp_min := a[left] - b[right]
				if tmp_min < min {
					min = tmp_min
				}
				pre_right = b[right]
			}
			//表示上一个被归并的元素在a中，表示出现交替归并
			if "left" == pre_flag {
				tmp_min := b[right] - pre_left
				if tmp_min < min {
					min = tmp_min
				}
				pre_right = b[right]
			}
			right++
			pre_flag = "right"
			continue
		}

	}

	if left < l1 && "right" == pre_flag {
		tmp_min := a[left] - pre_right
		if tmp_min < min {
			min = tmp_min
		}
	}

	if right < l2 && "left" == pre_flag {
		tmp_min := b[right] - pre_left
		if tmp_min < min {
			min = tmp_min
		}
	}

	fmt.Println(min)

	return min
}
