/*
link: https://leetcode.com/problems/valid-palindrome-ii/

Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:
Input: "aba"
Output: True
Example 2:
Input: "abca"
Output: True
Explanation: You could delete the character 'c'.
Note:
The string will only contain lowercase characters a-z. The maximum length of the string is 50000.


//这道题不要思考太多，比如最开始考虑使用动态规划就很不划算
//validCheck函数为标准的字符串回文判断函数，无删除多余字符的操作
//validPalindrome在左右不相等时，问题退化为删除左侧或者右侧多余字符后的标准回文字符串判断子问题

运行结果：
Runtime: 16 ms, faster than 72.50% of Go online submissions for Valid Palindrome II.
Memory Usage: 6.3 MB, less than 70.50% of Go online submissions for Valid Palindrome II.
*/

package main

import "fmt"

func main() {
	s := "abaca"
	fmt.Println(validPalindrome(s))
}

func validPalindrome(s string) bool {
	l := len(s)
	// if l < 2 {
	// 	return true
	// }

	left := 0
	right := l - 1

	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return validCheck(s, left+1, right) || validCheck(s, left, right-1)
		}
	}
	return true

}

func validCheck(s string, start int, end int) bool {

	left := start
	right := end
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
