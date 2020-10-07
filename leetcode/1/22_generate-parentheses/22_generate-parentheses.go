/*

link:
https://leetcode.com/problems/generate-parentheses/
*/

package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(4))
}

/*
运行结果：

Runtime: 0 ms, faster than 100.00% of Go online submissions for Generate Parentheses.
Memory Usage: 2.8 MB, less than 45.78% of Go online submissions for Generate Parentheses.
*/
func generateParenthesis(n int) []string {
	// left := 0
	// right := 0
	res := []string{}

	res = dfs("", n, 0, 0, res)

	return res
}

/*
// 这里res即为结果，注意考虑到slice作为形参时，如果进行append，
// 必须返回结果，否则无法被外部记录更新
// left表示左括号，right表示右括号。
// 这里可以画出分枝树， 规则主要有如下：
	1. left如果小于right，则为错误解，退出递归
	2. 如果left== right，则下一步必定只能选择left+1
	3. 如果left == n则表示左括号已经完成，只需填补必要的右括号即可结束递归
*/
func dfs(pre string, n, left, right int, res []string) []string {

	if right > left {
		return res
	}
	str := pre
	if left == n {
		for right < n {
			str += ")"
			right++
		}
		// fmt.Println("str::", str)
		res = append(res, str)
		return res
	} else if left < n {
		// fmt.Println("herer,", left)
		res = dfs(pre+"(", n, left+1, right, res)
		res = dfs(pre+")", n, left, right+1, res)
		return res
	}
	return res

}
