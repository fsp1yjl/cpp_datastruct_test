/*
link : https://leetcode.com/problems/stone-game/

参考：
https://labuladong.gitbook.io/algo/dong-tai-gui-hua-xi-lie/dong-tai-gui-hua-zhi-bo-yi-wen-ti
https://www.youtube.com/watch?v=xJ1Rc30Pyes

Alex and Lee play a game with piles of stones.  There are an even number of piles arranged in a row, and each pile has a positive integer number of stones piles[i].

The objective of the game is to end with the most stones.  The total number of stones is odd, so there are no ties.

Alex and Lee take turns, with Alex starting first.  Each turn, a player takes the entire pile of stones from either the beginning or the end of the row.  This continues until there are no more piles left, at which point the person with the most stones wins.

Assuming Alex and Lee play optimally, return True if and only if Alex wins the game.



Example 1:

Input: piles = [5,3,4,5]
Output: true
Explanation:
Alex starts first, and can only take the first 5 or the last 5.
Say he takes the first 5, so that the row becomes [3, 4, 5].
If Lee takes 3, then the board is [4, 5], and Alex takes 5 to win with 10 points.
If Lee takes the last 5, then the board is [3, 4], and Alex takes 4 to win with 9 points.
This demonstrated that taking the first 5 was a winning move for Alex, so we return true.


Constraints:

2 <= piles.length <= 500
piles.length is even.
1 <= piles[i] <= 500
sum(piles) is odd.


*/

package main

func main() {
	piles := []int{3, 9, 1, 2}
	stoneGame(piles)
}

func stoneGame(piles []int) bool {

	l := len(piles)
	if 0 == l {
		return false
	}
	type score struct {
		first  int
		second int
	}

	dp := make([][]score, l)

	for i := 0; i < l; i++ {
		dp[i] = make([]score, l)
	}
	// 初始化数组长度为1的子问题
	for i := 0; i < l; i++ {
		dp[i][i].first = piles[i]
		dp[i][i].second = 0
	}

	for length := 2; length <= l; length++ {
		for i := 0; i <= l-length; i++ {
			j := i + length - 1 // 注意这里要减1
			tmp1 := dp[i][j-1].second + piles[j]
			tmp2 := dp[i+1][j].second + piles[i]

			// 先选择右侧
			if tmp1 > tmp2 {
				dp[i][j].first = tmp1
				dp[i][j].second = dp[i][j-1].first
			} else {
				//先选择左侧
				dp[i][j].first = tmp2
				dp[i][j].second = dp[i+1][j].first
			}
		}
	}

	if dp[0][l-1].first > dp[0][l-1].second {
		return true
	} else {
		return false
	}

}
