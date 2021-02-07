/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。



进阶：

你可以不使用代码库中的排序函数来解决这道题吗？
你能想出一个仅使用常数空间的一趟扫描算法吗？


示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	nums := []int{2, 0, 2}
	sortColors(nums)
}

/*
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
Memory Usage: 2.1 MB, less than 66.40% of Go online submissions for Sort Colors.

注意i向后移动的条件，
如果nums[i] == 1 ，i++
如果nums[i] == 0 且i == left ，i++(如果发生了swap,不移动i)
如果nums[i] == 2 且i == right, i++ (如果发生swap, 不移动i)
*/
func sortColors(nums []int) {
	l := len(nums)
	if l == 1 {
		return
	}

	left := 0
	right := l - 1
	for i := 0; i < l; {
		if i > right {
			break
		}

		if 0 == nums[i] {
			nums[i], nums[left] = nums[left], nums[i]
			// bug1, 当i与left相等时，需要将i前移，
			if i == left {
				i++
			}
			left++
			continue
		}
		if 2 == nums[i] {
			nums[i], nums[right] = nums[right], nums[i]
			if i == right {
				i++
			}
			right--
			continue
		}

		if 1 == nums[i] {
			i++
		}
	}

	fmt.Println(nums)
}
