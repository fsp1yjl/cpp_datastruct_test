/*
link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).



Example 1:

Input: prices = [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
Example 2:

Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
Example 4:

Input: prices = [1]
Output: 0


说明：
题目是股票买卖最大收益的变种，这里最多进行两次交易


思路：
dp[i][j]表示第i天到第j天最多进行一次交易的最大收益

如果要收益最大，要不进行两次，要么进行一次交易
一次交易最大收益为dp[0][len]
两次交易最大收益为{dp[0][i] + dp[i+1][len] | 1<i<len}

*/

package main

import (
	"fmt"
)

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4, 10}

	prices := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0} // 13

	// prices := []int{1, 2, 3, 4, 5}

	fmt.Println(maxProfit(prices))
}

/*
最开始选择使用dp[i][j] 保存i到j期间进行一次买卖到最大收益，
这种方式在提交到继时候容易内存空间占用过大超时
故而分别使用一个dpLeft, dpRight构建两个一位数组


运行结果：
Runtime: 124 ms, faster than 26.62% of Go online submissions for Best Time to Buy and Sell Stock III.
Memory Usage: 10.5 MB, less than 5.19% of Go online submissions for Best Time to Buy and Sell Stock III.
*/

func maxProfit(prices []int) int {
	l := len(prices)

	if l < 2 {
		return 0
	}

	// fmt.Println("before len:", l)

	// fmt.Println("after len:", len(sample(prices)))

	type itemLeft struct {
		min    int
		profit int
	}

	type itemRight struct {
		max    int
		profit int
	}

	dpLeft := make([]itemLeft, l)   // dpLeft[i] 表示0到i期间进行一次买卖到最大收益
	dpRight := make([]itemRight, l) // dpRight[i]表示i到l-1 期间进行一次买卖的最大收益

	dpLeft[0].min = prices[0]
	dpLeft[0].profit = 0

	dpRight[l-1].max = prices[l-1]
	dpRight[l-1].profit = 0

	for j := 1; j < l; j++ {

		preMin := dpLeft[j-1].min
		preProfit := dpLeft[j-1].profit
		curPrice := prices[j]

		// fmt.Println("_|_|_", preMin, preProfit, curPrice)
		if curPrice < preMin {

			dpLeft[j].min = curPrice
			dpLeft[j].profit = preProfit
		} else {
			dpLeft[j].min = preMin
			if curPrice-preMin > preProfit {
				dpLeft[j].profit = curPrice - preMin
			} else {
				dpLeft[j].profit = preProfit
			}
		}
	}

	for j := l - 2; j >= 0; j-- {
		preMax := dpRight[j+1].max
		preProfit := dpRight[j+1].profit

		curPrice := prices[j]

		if curPrice > preMax {
			dpRight[j].max = curPrice
			dpRight[j].profit = preProfit
		} else {
			dpRight[j].max = preMax
			if preMax-curPrice > preProfit {
				dpRight[j].profit = preMax - curPrice
			} else {
				dpRight[j].profit = preProfit
			}
		}
	}

	max := dpRight[0].profit
	for i := 0; i < l-1; i++ {
		l := dpLeft[i].profit
		r := dpRight[i+1].profit

		if l+r > max {
			max = l + r
		}
	}

	return max
}

// 下面这个分治法容易递归运行超时
func maxProfit_2(prices []int) int {

	l := len(prices)
	if l < 2 {
		return 0
	}
	profit_1 := maxProfit1(prices)
	profit_2 := 0

	for i := 0; i < l-1; i++ {
		left_profit := maxProfit1(prices[:i+1])
		right_profit := maxProfit1(prices[i+1:])

		if left_profit+right_profit > profit_2 {
			profit_2 = left_profit + right_profit
		}
	}

	if profit_1 > profit_2 {
		return profit_1
	} else {
		return profit_2
	}

}

func maxProfit1(prices []int) int {

	l := len(prices)

	if l < 1 {
		return 0
	}

	min := prices[0]
	tmpMaxProfit := 0

	for i := 1; i < l; i++ {
		curPrice := prices[i]

		if curPrice < min {
			min = curPrice
		} else if curPrice-min > tmpMaxProfit {
			tmpMaxProfit = curPrice - min
		}
	}

	return tmpMaxProfit

}
