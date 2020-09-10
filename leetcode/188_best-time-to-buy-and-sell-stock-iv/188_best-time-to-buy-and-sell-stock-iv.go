/*

link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/


Say you have an array for which the i-th element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most k transactions.

Note:
You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

Example 1:

Input: [2,4,1], k = 2
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
Example 2:

Input: [3,2,6,5,0,3], k = 2
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
			 Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.



思路1: 将dp[i][k] 用来存储到第i天最多进行k次交易的最大收益,
状态转移方程可以为 dp[i][k] = max {tmp = dp[j][k-1] + onceMaxProfit(j+1, i, prices),  0 <j < i}

但是这个思路的时间复杂度太高， 即使将dp从二维数组降低为1维数组也无法满足要求
如果问题较小，可以这样处理，先pass


思路2:
dpHold[i][k]表示0到i区间最多进行k次买入且最终持仓情况下的最小成本
dpEmpty[i][k]表示0到i区间最多进行k次买入且最终空仓情况下的最大收益

状态转移方程:
	dpHold[i][k] = max (dpHold[i-1][k], dpEmpty[i-1][k-1] - prices[i])
	dpEmpty[i][k] = max (dpEmpty[i-1][k], dpHold[i-1][k] + prices[i]) //注意这里的状态方程是dpHold[i-1][k]而非dpHold[i-1][k-1]



*/

package main

import "fmt"

func main() {
	// prices := []int{3, 2, 6, 5, 0, 3}
	// prices := []int{2, 2, 4, 1, 5}
	// prices := []int{3, 3, 5, 0, 0, 3, 1, 4} //k=2, output: 6

	prices := []int{3, 3, 3, 5, 0, 0, 3, 1, 4}
	// k := 1000000000
	k := 2
	// fmt.Println("max_k:", max_k(prices))
	fmt.Println(maxProfit(k, prices))
}

func maxProfit(k int, prices []int) int {
	l := len(prices)

	if k < 1 || l < 2 {
		return 0
	}

	if 2*k >= l {
		// 这里要处理k值异常大的情况，防止内存溢出，也为了防止无效计算
		// 如果交易次数大于天数，则问题退化为maxProfitWithoutLImit， 参考leetcode122
		return maxProfitWithoutLImit(prices)
	}

	//dpHold[i][k]表示0到i区间最多进行k次买入且最终持仓情况下的最小成本
	dpHold := make([][]int, l)

	//dpEmpty[i][k]表示0到i区间最多进行k次买入且最终空仓情况下的最大收益
	dpEmpty := make([][]int, l)

	for i := 0; i < l; i++ {
		dpHold[i] = make([]int, k+1)
		dpEmpty[i] = make([]int, k+1)
	}

	// 先对0次交易及第一天完结的相关情况做初始化
	for i := 0; i <= k; i++ {
		dpEmpty[0][i] = 0
		if i != 0 {
			dpHold[0][i] = 0 - prices[0]
		}

	}

	for i := 0; i < l; i++ {
		dpEmpty[i][0] = 0
	}

	/* 状态转移方程
	dpHold[i][k] = max (dpHold[i-1][k], dpEmpty[i-1][k-1] - prices[i])
	dpEmpty[i][k] = max (dpEmpty[i-1][k], dpHold[i-1][k] + prices[i]) //注意这里的状态方程是dpHold[i-1][k]而非dpHold[i-1][k-1]

	*/

	for i := 1; i < l; i++ {
		for t := 1; t <= k; t++ {
			// 计算dpHold[i][t]
			if dpHold[i-1][t] > dpEmpty[i-1][t-1]-prices[i] {
				// 本次不买，则0到i区间持仓状态下的最小成本转换到0到i-1区间
				dpHold[i][t] = dpHold[i-1][t]
			} else {
				// 本次买入，则0到i区间持仓状态下的最小成本转换到 0到i-1区间空仓最大盈利 减去本次买入价格
				dpHold[i][t] = dpEmpty[i-1][t-1] - prices[i]
			}

			//计算dpEmpty[i][t]

			if dpEmpty[i-1][t] > dpHold[i-1][t]+prices[i] {
				// 本次不卖出的最大收益更大
				dpEmpty[i][t] = dpEmpty[i-1][t]
			} else {
				// 本次卖出的最大收益更大
				dpEmpty[i][t] = dpHold[i-1][t] + prices[i]
			}

		}
	}

	// fmt.Println("empty:", dpEmpty)
	// fmt.Println("hold:", dpHold)

	return dpEmpty[l-1][k]

}

// 不限制交易次数的最大收益
func maxProfitWithoutLImit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}

	// dpHold[i]表示0到i期间，如果为持仓状态的最小成本
	dpHold := make([]int, l)
	// dpEmpty[i]表示0到i期间，空仓状态下的最大收益
	dpEmpty := make([]int, l)

	dpHold[0] = 0 - prices[0]
	dpEmpty[0] = 0

	for i := 1; i < l; i++ {
		// 计算dpHold[i], 有两种可能，本天买或者本天不买
		if dpHold[i-1] > dpEmpty[i-1]-prices[i] {
			// 本次不买，则0到i区间持仓状态下的最小成本转换到0到i-1区间
			dpHold[i] = dpHold[i-1]
		} else {
			// 本次买入，则，则0到i区间持仓状态下的最小成本转换到 0到i-1区间空仓最大盈利 减去本次买入价格
			dpHold[i] = dpEmpty[i-1] - prices[i]
		}

		//计算dpEmpty[i], 有两种可能，本天卖或者不卖

		if dpEmpty[i-1] > dpHold[i-1]+prices[i] {
			// 本次不卖出的最大收益更大
			dpEmpty[i] = dpEmpty[i-1]
		} else {
			// 本次卖出的最大收益更大， 为0到i-1区间持仓的最小成本 加上本次卖出的价格
			dpEmpty[i] = dpHold[i-1] + prices[i]
		}
	}

	// 最终0到l-1区间最终空仓的最大收益即为所需结果
	return dpEmpty[l-1]

}
