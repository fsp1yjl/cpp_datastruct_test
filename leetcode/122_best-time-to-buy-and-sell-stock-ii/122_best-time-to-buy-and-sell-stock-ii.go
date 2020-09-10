/*

link : https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/


Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.


Constraints:

1 <= prices.length <= 3 * 10 ^ 4
0 <= prices[i] <= 10 ^ 4

*/

/*

分析，本题和121题类似，但是可以多次交易，要做到收益最大
要做多次买卖成对交易， 在每次单调增的发起点买，结束点卖
基本思路，构建dp, dp[i][j]表示i到j的价格是否单调递增，是保存true，不是保存false
*/
package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	l := len(prices)

	if l < 2 {
		return 0
	}

	// start := false
	buyFlag := false
	buy := 0
	profit := 0

	for i := 0; i < l; i++ {
		if !buyFlag {
			if i < l-1 && prices[i+1] > prices[i] { // 进行一次买入
				buyFlag = true
				// start = true

				buy = prices[i]
				// fmt.Println("buy:", buy)

			}
		} else {
			if (i+1 < l && prices[i+1] < prices[i]) || i == l-1 {
				if buyFlag {
					// fmt.Println("SELL:", prices[i])

					profit += prices[i] - buy // 进行一次卖出
					buyFlag = false
					buy = 0
				}
			}
		}
	}

	return profit
}

/*

output:

Runtime: 4 ms, faster than 96.19% of Go online submissions for Best Time to Buy and Sell Stock II.
Memory Usage: 3.1 MB, less than 76.19% of Go online submissions for Best Time to Buy and Sell Stock II.
*/
