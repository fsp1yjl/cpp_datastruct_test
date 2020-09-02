/*

字符串最大回文子数

Given a string s, find the longest palindromic subsequence's length in s. You may assume that the maximum length of s is 1000.

Example 1:
Input:

"bbbab"
Output:
4
One possible longest palindromic subsequence is "bbbb".


Example 2:
Input:

"cbbd"
Output:
2
One possible longest palindromic subsequence is "bb".


Constraints:

1 <= s.length <= 1000
s consists only of lowercase English letters.
*/

package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("bbbabb"))
}

func longestPalindromeSubseq(s string) int {
	l := len(s)
	if 0 == l {
		return 0
	}
	dp := make([][]int, l) // dp[i][j] 表示s[i:j+1]对应字符串的最大回文长度
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	// 初始化长度为1的子字符串回文为1
	for i := 0; i < l; i++ {
		dp[i][i] = 1
	}

	// 依次求出从1到l-1 长度到子字符串到最大回文数， 最终l长度到最大回文就包括在l-1的问题中
	for length := 2; length <= l; length++ {
		for i := 0; i <= l-length; i++ { // 注意这里的判断条件是小于等于
			j := i + length - 1

			if s[i] == s[j] {
				if i+1 <= j-1 {
					dp[i][j] = dp[i+1][j-1] + 2
				} else {
					dp[i][j] = 2
				}
			} else {
				dp_1 := dp[i+1][j]
				dp_2 := dp[i][j-1]

				if dp_1 > dp_2 {
					dp[i][j] = dp_1
				} else {
					dp[i][j] = dp_2
				}
			}
		}
	}

	// fmt.Println("dp::", dp)
	return dp[0][l-1]
}
