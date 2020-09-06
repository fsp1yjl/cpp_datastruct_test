/*
link : https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.


*/

package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4, 6}
	fmt.Println(maxProfit(prices))
}

/*
思考过程，最开始考虑用动态规划，依次作出长度为1 到 len的 dp表，
dp每个元素存放当前的最小价格和最大收益， dp[i][j]= {min_price, max_profit}
状态转移方程为
if prices[j+1] < dp[i][j].min_price {
	// 如果当前价格低于前面的最小价格，要更新dp[i][j+1]的最小价格，但是最大收益不变
	dp[i][j+1].min_price = prices[j+1]
	dp[i][j+1].max_profit = dp[i][j].max_profit
} else {
	//如果当前价格高于dp[i][j]的最小价格，且差价大于dp[i][j]的最大差价，则更新最大差价，但最小价格不变
	if prices[j+1] - dp[i][j].min_price >  dp[i][j].max_profit {
		dp[i][j+1].min_price = dp[i][j].min_price
		dp[i][j+1].max_profit = prices[j+1] - dp[i][j].min_price
	} else {
		//如果当前价格高于dp[i][j]的最小价格，但差价不大于dp[i][j]的最大差价
		dp[i][j+1].min_price = dp[i][j].min_price
		dp[i][j+1].max_profit = dp[i][j].max_profit
	}
}


简化：
假设有5个价格，这里我们要的结果是dp[0][4],从dp分析看出，这里只要依次计算dp[0][1], dp[0][2], dp[0][3]即可

*/

func maxProfit(prices []int) int {

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

/*
测试结果：
Runtime: 4 ms, faster than 95.25% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 3.4 MB, less than 11.35% of Go online submissions for Best Time to Buy and Sell Stock.

*/
