/*
link: https://leetcode.com/problems/house-robber-ii/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
             because they are adjacent houses.
Example 2:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
			 Total amount you can rob = 1 + 3 = 4.


*/

// 分析， 同198题类似的，但是这里除了相邻不能同时取得外， 另外加了一个限制，就是数组变成一个环
// 因为有个环，可以将问题转化为三种情况， 仅偷左边，仅偷右边， 左右都不偷三种情况
// 子文题可以使用198的方法

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

/*
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber II.
Memory Usage: 2.2 MB, less than 17.89% of Go online submissions for House Robber II.
*/
func rob(nums []int) int {
	l := len(nums)
	if 0 == l {
		return 0
	}
	if 1 == l {
		return nums[0]
	}

	s_left := nums[0] + rob_1(2, l-2, nums)    // 偷0，不偷l-1
	s_right := nums[l-1] + rob_1(1, l-3, nums) //偷l-1 , 不偷0
	s_double := rob_1(1, l-2, nums)            // 0和l-1均不偷

	s := s_left
	if s_right > s {
		s = s_right
	}

	if s_double > s {
		s = s_double
	}

	return s

}

func rob_1(start int, end int, nums []int) int {
	lenght := len(nums)
	l := end - start + 1
	if start > lenght-1 || end < 0 || l <= 0 {
		return 0
	}

	if 1 == l {
		return nums[start]
	}

	dp := make([][]int, l)

	for i := 0; i < l; i++ {
		// dp[i][0]存放 0-i区间不偷nums[i]的最大收益
		// dp[i][1]存放 0-i区间偷nums[i]的最大收益
		//依次求出dp[i][0]   0<=i < l, 最终结果为max(dp[l-1][0], dp[l-1][1])
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nums[start]

	for i := 1; i < l; i++ {
		// 计算dp[i][0]
		s1 := dp[i-1][1]
		s2 := dp[i-1][0]
		if s1 > s2 {
			dp[i][0] = s1
		} else {
			dp[i][0] = s2
		}

		//计算dp[i][1]
		if 1 == i {
			dp[i][1] = nums[i+start]
		} else {
			if dp[i-2][0] > dp[i-2][1] {
				dp[i][1] = nums[i+start] + dp[i-2][0]
			} else {
				dp[i][1] = nums[i+start] + dp[i-2][1]
			}
		}
	}
	if dp[l-1][0] > dp[l-1][1] {
		return dp[l-1][0]
	} else {
		return dp[l-1][1]
	}
}
