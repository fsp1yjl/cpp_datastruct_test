/*

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。



示例 1：


输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func main() {

}

/*
思路：dp

提交结果：
Runtime: 8 ms, faster than 87.57% of Go online submissions for Minimum Path Sum.
Memory Usage: 3.9 MB, less than 54.59% of Go online submissions for Minimum Path Sum.

*/
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([]int, n)

	//初始化第一行
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for j := 1; j < m; j++ {
		dp[0] = dp[0] + grid[j][0]
		for i := 1; i < n; i++ {
			if dp[i] < dp[i-1] {
				dp[i] = dp[i] + grid[j][i] //上方更小
			} else {
				dp[i] = dp[i-1] + grid[j][i] //左边更小
			}
		}
	}

	return dp[n-1]
}
