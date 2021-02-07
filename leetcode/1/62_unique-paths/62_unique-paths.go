/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



*/

package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

/*
思路：
从网格左上到右下的所有可选路径数。 dp ,dp[i][j]表示到达nums[i][j]的可选方案数，
dp[i][j] = dp[i-1][j] + dp[i][j-1]  从第一行依次向后。 可以将二维化为一维， dp[i] = dp[i-1] + dp[i]

提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
Memory Usage: 1.9 MB, less than 88.89% of Go online submissions for Unique Paths.

*/

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			dp[i] = dp[i-1] + dp[i]
		}
	}

	return dp[n-1]
}
