/*
YouTube理解参考：https://www.youtube.com/watch?v=UQF2VSGbZjw

题目link: https://leetcode.com/problems/super-egg-drop/
You are given K eggs, and you have access to a building with N floors from 1 to N.

Each egg is identical in function, and if an egg breaks, you cannot drop it again.

You know that there exists a floor F with 0 <= F <= N such that any egg dropped at a floor higher than F will break, and any egg dropped at or below floor F will not break.

Each move, you may take an egg (if you have an unbroken one) and drop it from any floor X (with 1 <= X <= N).

Your goal is to know with certainty what the value of F is.

What is the minimum number of moves that you need to know with certainty what F is, regardless of the initial value of F?



Example 1:

Input: K = 1, N = 2
Output: 2
Explanation:
Drop the egg from floor 1.  If it breaks, we know with certainty that F = 0.
Otherwise, drop the egg from floor 2.  If it breaks, we know with certainty that F = 1.
If it didn't break, then we know with certainty F = 2.
Hence, we needed 2 moves in the worst case to know what F is with certainty.
Example 2:

Input: K = 2, N = 6
Output: 3
Example 3:

Input: K = 3, N = 14
Output: 4
*/
package main

import "fmt"

func main() {
	K:= 4
	N:= 25
	fmt.Println(superEggDrop(K, N))
	fmt.Println(superEggDrop2(K, N))
}

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1) //dp[i][j] 表示i鸡蛋，j次drop可以确定的楼层数
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
		// for j := 0; j <= N; j++ {
		// 	dp[i][j] = -1
		// }
	}

	dp[0][0] = 0
	if K == 0 || N == 0 {
		return 0
	}
	if K == 1 {
		return N
	}
	//如果只有一个鸡蛋，dp[1][j]均初始化为drop次数j
	for i := 0; i <= N; i++ {
		dp[1][i] = i
		dp[0][i] = 0
	}

	//如果drop次数为0，则可检测楼数为0
	for i := 0; i <= K; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= K; i++ {
		for j := 1; j <= N; j++ {
			dp[i][j] = dp[i-1][j-1] + 1 + dp[i][j-1]
			if dp[i][j] >= N {

				// fmt.Println("i,j", i, j)
				break
			}
		}
	}
	fmt.Println("dp:", dp)
	for j := 1; j <= N; j++ {
		if dp[K][j] >= N {
			return j
		}
	}

	return N

}


func superEggDrop2(K int , N int) int {
	dp := make([][]int, K+1)
	for i:=1; i<=K; i++ {
		dp[i] = make([]int, N+1)
	}
	//只有一个鸡蛋，按楼高更新
	for j:=1; j <=N; j++ {
		dp[1][j] = j
	}

	//如果楼高为1， 则需要1次
	for i:=1;i<=K;i++ {
		dp[i][1] = 1
	}
	var helper  func(int,  int) int
	helper = func(k int, n int) int {
		if n == 0 {
			return 0
		}
		if k > n {
			return helper(n, n)
		}
		if dp[k][n] != 0 {
			fmt.Println("HERE")
			return dp[k][n]
		}

		min := n
		for t:= 1; t<=n; t++{
			// if drop egg broke, check lower floor
			brokeCase := helper(k-1,t-1) + 1
			// if not break, check upper floor
			safeCase := helper(k, n-t) + 1

			curMax := brokeCase
			if safeCase > curMax {
				curMax = safeCase
			}

			curRes := curMax
			if curRes < min {
				min = curRes
			}
		}

		return min
	}

	return helper(K, N)
}
