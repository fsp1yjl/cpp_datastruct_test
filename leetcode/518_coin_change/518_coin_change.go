/*
link: https://leetcode.com/problems/coin-change-2/



You are given coins of different denominations and a total amount of money. Write a function to compute the number of combinations that make up that amount. You may assume that you have infinite number of each kind of coin.



Example 1:

Input: amount = 5, coins = [1, 2, 5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
Example 3:

Input: amount = 10, coins = [10]
Output: 1


Note:

You can assume that

0 <= amount <= 5000
1 <= coin <= 5000
the number of coins is less than 500
the answer is guaranteed to fit into signed 32-bit integer
*/

package main

import "fmt"

func main() {

	amount := 3999
	coins := []int{200, 217, 234, 251, 268, 285, 302, 319, 336, 353, 370, 387, 404, 421, 438, 455, 472, 489, 506, 523, 540, 557, 574, 591, 608, 625, 642, 659, 676, 693, 710, 727, 744, 761, 778, 795, 812, 829, 846, 863, 880, 897, 914, 931, 948, 965, 982, 999, 1016, 1033, 1050, 1067, 1084, 1101, 1118, 1135, 1152, 1169, 1186, 1203, 1220, 1237, 1254, 1271, 1288, 1305, 1322, 1339, 1356, 1373, 1390, 1407, 1424, 1441, 1458, 1475, 1492, 1509, 1526, 1543, 1560, 1577, 1594, 1611, 1628, 1645, 1662, 1679, 1696, 1713, 1730, 1747, 1764, 1781, 1798, 1815, 1832, 1849, 1866, 1883, 1900, 1917, 1934, 1951, 1968, 1985, 2002, 2019, 2036, 2053, 2070, 2087, 2104, 2121, 2138, 2155, 2172, 2189, 2206, 2223, 2240, 2257, 2274, 2291, 2308, 2325, 2342, 2359, 2376, 2393, 2410, 2427, 2444, 2461, 2478, 2495, 2512, 2529, 2546, 2563, 2580, 2597, 2614, 2631, 2648, 2665, 2682, 2699, 2716, 2733, 2750, 2767, 2784, 2801, 2818, 2835, 2852, 2869, 2886, 2903, 2920, 2937, 2954, 2971, 2988, 3005, 3022, 3039, 3056, 3073, 3090, 3107, 3124, 3141, 3158, 3175, 3192, 3209, 3226, 3243, 3260, 3277, 3294, 3311, 3328, 3345, 3362, 3379, 3396, 3413, 3430, 3447, 3464, 3481, 3498, 3515, 3532, 3549, 3566, 3583, 3600, 3617, 3634, 3651, 3668, 3685, 3702, 3719, 3736, 3753, 3770, 3787, 3804, 3821, 3838, 3855, 3872, 3889, 3906, 3923, 3940, 3957, 3974, 3991, 4008, 4025, 4042, 4059, 4076, 4093, 4110, 4127, 4144, 4161, 4178, 4195, 4212, 4229, 4246, 4263, 4280, 4297, 4314, 4331, 4348, 4365, 4382, 4399, 4416, 4433, 4450, 4467, 4484, 4501, 4518, 4535, 4552, 4569, 4586, 4603, 4620, 4637, 4654, 4671, 4688, 4705, 4722, 4739, 4756, 4773, 4790, 4807, 4824, 4841, 4858, 4875, 4892, 4909, 4926, 4943, 4960, 4977, 4994}

	// 预期结果96650

	fmt.Println(change(amount, coins))
}

// todo 这里dp保存为2维数组，因为dp[i][]只依赖dp[i-1][]
//可以考虑把代码修改为dp仅使用一维数组对形式
func change(amount int, coins []int) int {

	coinCnt := len(coins)

	if 0 == amount {
		return 1
	} else if 0 == coinCnt {
		return 0
	}

	dp := make([][]int, coinCnt)

	for i := 0; i < coinCnt; i++ {
		dp[i] = make([]int, amount+1)
	}

	//初始化第一行
	for j := 1; j <= amount; j++ {
		coinOne := coins[0]
		if 0 == j%coinOne {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	for i := 1; i < coinCnt; i++ {
		par := coins[i]
		for j := 1; j <= amount; j++ {
			pages := j / par
			for k := pages; k >= 0; k-- {
				h := j - par*k
				if h > 0 {
					dp[i][j] += dp[i-1][j-par*k]
				} else {
					// 这里考虑到剩余为0的情况，直接对可选情况+1
					dp[i][j] += 1
				}

			}
		}
	}
	return dp[coinCnt-1][amount]
}
