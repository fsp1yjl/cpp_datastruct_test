/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 7, 6, 3, 5, 1}
	fmt.Println(combinationSum(nums, 9))
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var helper func([]int, int, []int)
	helper = func(candidates []int, target int, pre []int) {
		// fmt.Println("before:", target, "-----,pre:", pre)

		l := len(candidates)
		tmp := pre
		first := candidates[0]
		count := 0

		for first*count <= target {
			if count > 0 {
				tmp = append(tmp, first)
			}
			if first*count == target {
				// bug1, 这里需要对tmp进行一次深拷贝
				dd := make([]int, 0)
				dd = append(dd, tmp...)

				res = append(res, dd)
				break
			}
			if l > 1 && (target-first*count) > first {
				// candidates为1，表示无后续可选元素
				helper(candidates[1:], target-first*count, tmp)
			}
			count++
		}
	}

	// fmt.Println("len:", candidates)
	helper(candidates, target, []int{})
	return res

}
