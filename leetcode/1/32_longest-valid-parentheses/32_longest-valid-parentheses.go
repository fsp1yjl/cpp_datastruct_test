/*

给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。


输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

思路： dp,  dp[i]表示i开头的最大有效括号长度
*/

package main

import "fmt"

func main() {
	s := "((()(())"
	fmt.Println(longestValidParentheses(s))
}

/*

提交结果有些惨：
执行用时：
12 ms, 在所有 Go 提交中击败了9.46%的用户内存消耗：2.8 MB
, 在所有 Go 提交中击败了80.35%的用户
*/
func longestValidParentheses1(s string) int {
	l := len(s)

	if l < 2 {
		return 0
	}

	dp := make([]int, l-1)

	flag := false
	//初始化长度为2的子问题
	for i := 0; i < l-1; i++ {
		if string(s[i]) == "(" && string(s[i+1]) == ")" {

			dp[i] = 2
			flag = true
		}
	}
	// 当遍历后，发现dp无法再延展更新，则退出循环
	for flag {
		flag = false
		//如果一次循环后，不再发生长度更新，则flag为false ,退出外层循环
		for i := 0; i < l-1; {

			tmp := dp[i]

			// dp[i] + "()"的情况
			if tmp > 0 && i+tmp < l-1 && dp[i+tmp] > 0 {
				dp[i] = tmp + dp[i+tmp]
				i += tmp
				flag = true
				// fmt.Println(dp)

				continue
			}

			// "(" + dp[i] + ")"的情况
			if i-1 >= 0 && i+tmp <= l-1 && dp[i] > 0 && string(s[i-1]) == "(" && string(s[i+tmp]) == ")" {
				dp[i-1] = tmp + 2
				i += tmp
				flag = true
				// fmt.Println(dp)

				continue
			}

			if tmp > 0 {
				i += tmp
			} else {
				i++
			}
		}

	}

	max := 0
	// fmt.Println(dp)
	for i := 0; i < l-1; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

/*
同上，依旧用dp[i]表示i开头的最长括号，但是运行思路改为倒着遍历，
这样计算dp[i]时，就可以利用已经求得的dp[i+1], dp[i+2]
当s[i] = (, s[i+1]=) ,dp[i] = 2 + dp[i+2]   注，因为s[i+1] 为右括号时，dp[i+1]为0
当s[i]= (,  dp[i+1] =k,k> 0时，如果s[i+k+1] = ),则 dp[i] = 2 + dp[i+1] + dp[i+k+2]


提交性能大幅度提升

Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Valid Parentheses.
Memory Usage: 2.8 MB, less than 56.12% of Go online submissions for Longest Valid Parentheses.
*/
func longestValidParentheses(s string) int {
	l := len(s)

	if l < 2 {
		return 0
	}

	dp := make([]int, l)

	//初始化长度为2的子问题
	for i := l - 2; i >= 0; i-- {
		if string(s[i]) == "(" && string(s[i+1]) == ")" {
			if i+2 >= l {
				dp[i] = 2
			} else {
				dp[i] = 2 + dp[i+2]
			}
			continue
		}

		if string(s[i]) == "(" && dp[i+1] > 0 {
			k := dp[i+1]

			if i+k+1 > l-1 {
				continue
			}
			if string(s[i+k+1]) == ")" {
				dp[i] = 2 + dp[i+1]

				if i+k+2 < l {
					dp[i] += dp[i+k+2]
				}
			}
		}
	}

	max := 0
	// fmt.Println(dp)
	for i := 0; i < l-1; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
