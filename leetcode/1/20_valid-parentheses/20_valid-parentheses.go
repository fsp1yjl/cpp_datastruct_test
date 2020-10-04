/*
link: https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
Example 4:

Input: s = "([)]"
Output: false
Example 5:

Input: s = "{[]}"
Output: true


*/

package main

import "fmt"

func main() {
	s := "((()[]))"
	fmt.Println(isValid(s))
}

/*

运行结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 2.1 MB, less than 24.70% of Go online submissions for Valid Parentheses.
*/
func isValid(s string) bool {
	l := len(s)

	if l%2 != 0 {
		return false
	}

	stack := []string{}

	for i := 0; i < l; i++ {
		c := string(s[i])
		if "(" == c ||
			"[" == c ||
			"{" == c {
			//左括号入栈
			stack = append(stack, c)
		} else {
			stack_len := len(stack)
			if 0 == stack_len {
				return false
			}
			last := stack[stack_len-1]
			if (")" == c && "(" == last) ||
				("]" == c && "[" == last) ||
				("}" == c && "{" == last) {
				stack = stack[:stack_len-1]
			} else {
				return false
			}
		}
	}

	if 0 == len(stack) {
		return true
	} else {
		return false
	}
}
