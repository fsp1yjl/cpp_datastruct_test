/*
link: https://leetcode.com/problems/permutation-sequence/

The set [1, 2, 3, ..., n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.



Example 1:

Input: n = 3, k = 3
Output: "213"
Example 2:

Input: n = 4, k = 9
Output: "2314"
Example 3:

Input: n = 3, k = 1
Output: "123"


Constraints:

1 <= n <= 9  (这里不用考虑头部为0的情况了)
1 <= k <= n!

*/

package main

import (
	"strconv"
)

func main() {
	getPermutation(4, 9)
}

func getPermutation(n int, k int) string {

	nums := make([]int, 0)

	res := make([]int, 0)

	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	multi := 1
	for i := n; i > 1; i-- {
		multi *= i
	}

	weight := k - 1 // 这里最开始一定要对k做一个减1操作

	for i := n; i >= 1; i-- {
		// fmt.Println("multi, weight:", multi, weight)

		multi = multi / i

		curIndex := 0
		curIndex = weight / multi
		res = append(res, nums[curIndex])
		nums = append(nums[:curIndex], nums[curIndex+1:]...)
		weight = weight % multi

		// fmt.Println("ddddddd aftere:", nums)

	}
	str := ""

	for i := 0; i < len(res); i++ {
		str += strconv.Itoa(res[i])
	}

	// fmt.Println("ddddddd aftere:", str)

	return str
}
