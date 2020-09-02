/*


link: https://leetcode-cn.com/problems/longest-common-subsequence/

Given two strings text1 and text2, return the length of their longest common subsequence.

A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.



If there is no common subsequence, return 0.



Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3

*/

package main

import "fmt"

func main() {

	// text1 := "abcde"
	// text2 := "dacd"
	// 最长子序列为3 acd

	text1 := "mhunuzqrkzsnidwbun"
	text2 := "szulspmhwpazoxijwbq"
	// 6
	fmt.Println(longestCommonSubsequence(text1, text2))
}

func longestCommonSubsequence(text1 string, text2 string) int {

	len1 := len(text1)
	len2 := len(text2)

	mem := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		mem[i] = make([]int, len2+1)
	}

	for i := 0; i < len1+1; i++ {
		for j := 0; j < len2+1; j++ {
			mem[i][j] = -1
		}
	}

	return dp(text1, text2, len1, len2, mem)

}

// 使用动态规划思维
//如果text1,text2的最后一位相同，则问题转化为text1-1, text2-1
// 如果最后一位不相同，则问题转换为 (text1, text2-1 ) (text1 -1 ,text2-1) 两个子问题的最大值
// 这里为了防止递归子问题的多次重复计算，使用一个mem缓存存储子问题的解

//这个题也可以考虑自底向上，一次从(dp[0][0]...dp[0][right])到(dp[left][0]...dp[left][right])，记录动态规划表去解决
// link: https://labuladong.gitbook.io/algo/dong-tai-gui-hua-xi-lie/zui-chang-gong-gong-zi-xu-lie

func dp(text1 string, text2 string, left int, right int, mem [][]int) int {

	if mem[left][right] != -1 {
		return mem[left][right]
	}

	if 0 == left || 0 == right {
		mem[left][right] = 0
		return 0
	}

	var res int
	if text1[left-1] == text2[right-1] {
		res = dp(text1, text2, left-1, right-1, mem) + 1
	} else {
		res1 := dp(text1, text2, left-1, right, mem)
		res2 := dp(text1, text2, left, right-1, mem)

		if res1 > res2 {
			res = res1
		} else {
			res = res2
		}
	}

	mem[left][right] = res

	return res
}
