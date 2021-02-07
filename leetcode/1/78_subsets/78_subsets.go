/*

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{0, 3, 5, 7, 9}))
}

/*

提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Subsets.
Memory Usage: 2.3 MB, less than 93.31% of Go online submissions for Subsets.


注意： 这里在递归函数里，传递和处理pre数组时，进行一次深拷贝，并限定len,cap
*/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var helper func([]int, int, []int)

	helper = func(nums []int, left int, pre []int) {
		prefix := make([]int, len(pre), len(pre)) //bug1: 这里即指定len,又制定cap，防止发生错误的修改
		copy(prefix, pre)
		l := len(nums)
		if l-1 == left {
			tmp1 := prefix
			tmp2 := append(prefix, nums[left])
			res = append(res, tmp1)
			res = append(res, tmp2)
		} else {
			helper(nums, left+1, prefix)
			helper(nums, left+1, append(prefix, nums[left]))
		}
	}

	helper(nums, 0, []int{})
	// fmt.Println("len:", len(res))
	return res
}
