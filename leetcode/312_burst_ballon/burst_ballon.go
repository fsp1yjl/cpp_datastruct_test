/*

youtube : https://www.youtube.com/watch?v=z3hu2Be92UA

link:https://leetcode.com/problems/burst-balloons/

Given n balloons, indexed from 0 to n-1. Each balloon is painted with a number on it represented by array nums. You are asked to burst all the balloons. If the you burst balloon i you will get nums[left] * nums[i] * nums[right] coins. Here left and right are adjacent indices of i. After the burst, the left and right then becomes adjacent.

Find the maximum coins you can collect by bursting the balloons wisely.

Note:

You may imagine nums[-1] = nums[n] = 1. They are not real therefore you can not burst them.
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
Example:

Input: [3,1,5,8]
Output: 167
Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
             coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167



c[i][j] = max{c[i][k-1] + nums[i-1]*nums[k]*nums[j+1] + c[k+1][j]  |  i<=k <=j}
*/

package main

import "fmt"

func main() {

	nums := []int{3, 1, 5, 8}

	maxCoins(nums)
}

// 本题要和887扔鸡蛋一样，要用逆向思维去处理
func maxCoins(nums []int) int {
	n := len(nums)
	if 0 == n {
		return 0
	}
	nums = append(nums, 1)
	left := []int{1}
	nums = append(left, nums...)

	fmt.Println(nums)
	// 切片首尾+1 后，长度增2
	// length := n + 2

	dp := make([][]int, n+2)
	for i := 1; i < n+2; i++ {
		dp[i] = make([]int, n+2)
		dp[i][i] = nums[i]
	}

	// fmt.Println("dp:", dp)
	// 这里的循环考虑从数组长度为1的子问题开始，最终到达长度为n的问题
	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			for k := i; k <= j; k++ {
				tmp := dp[i][k-1] + dp[k+1][j] + nums[i-1]*nums[k]*nums[j+1]
				if tmp > dp[i][j] {
					dp[i][j] = tmp
				}
			}
		}
	}
	// fmt.Println("dp:", dp)
	// fmt.Println("dp---1-n:", dp[1][n])
	return dp[1][n]

}
