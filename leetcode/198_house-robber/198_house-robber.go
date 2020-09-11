/*
link:  https://leetcode.com/problems/house-robber/


You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.



Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.


Constraints:

0 <= nums.length <= 100
0 <= nums[i] <= 400
*/

package main

import "fmt"

func main () {
	nums := []int{2,7,9,3,1}
	fmt.Println(rob(nums))
}

/*

提交运行结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
Memory Usage: 2.1 MB, less than 22.70% of Go online submissions for House Robber.
*/

func rob(nums []int) int {
	l := len(nums) 
	if 0 == l {
		return 0
	}

	if 1 == l {
		return nums[0]
	}

	dp := make([][]int, l)
	
	for i:=0; i < l; i++ {
		// dp[i][0]存放 0-i区间不偷nums[i]的最大收益
		// dp[i][1]存放 0-i区间偷nums[i]的最大收益
		//依次求出dp[i][0]   0<=i < l, 最终结果为max(dp[l-1][0], dp[l-1][1])
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i:=1 ; i < l; i++ {
		// 计算dp[i][0]
		s1 := dp[i-1][1]
		s2 := dp[i-1][0]
		if s1 > s2 {
			dp[i][0]  = s1
		} else {
			dp[i][0]  = s2
		}


		//计算dp[i][1]
		if 1 == i {
			dp[i][1] = nums[i]
		} else {
			if dp[i-2][0] > dp[i-2][1] {
				dp[i][1] = nums[i] + dp[i-2][0]
			} else {
				dp[i][1] = nums[i] + dp[i-2][1]
			}
		}
	}
	if dp[l-1][0] > dp[l-1][1] {
		return dp[l-1][0]
	} else {
		return dp[l-1][1]
	}
}
