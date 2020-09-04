/*
正则表达式

https://leetcode.com/problems/regular-expression-matching/

Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false

*/

package main

import "fmt"

func main() {
	// false
	s := "mississippi"
	p := "mis*is*p*."

	//false
	// s := "a"
	// p := "ab*a"

	//true
	// s := "a"
	// p := "a*"

	//true
	// s := "aa"
	// p := "c*a*"

	// false
	// s := "aab"
	// p := "b.*"

	// ture
	// s := ""
	// p := "c*c*"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	len_s := len(s)
	len_p := len(p)

	if 0 == len_p {
		if 0 == len_s {
			// 这里要考虑是不进行replace的情况下
			//操作之后 s, p是否都为空，如果都为空，则返回true
			return true
		} else {
			return false
		}
	}

	if 1 == len_p {
		if len_p == len_s {
			if p[0] == s[0] || '.' == p[0] {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	}

	// 正则表达式长度至少为2的情况
	if '*' == p[len_p-1] {
		p_pre := p[len_p-2]
		if '*' == p_pre {
			return false
		} else if '.' == p_pre {
			replace_0 := false
			replace_1 := false
			replace_more := false
			if 0 == len_s {
				if 2 == len_p { // 这里要判断，s为空的时候，p经过replace是否为空
					return true
				} else {
					return false
				}
			} else {
				replace_1 = isMatch(s[:len_s-1], p[:len_p-2])
				replace_0 = isMatch(s, p[:len_p-2])
				replace_more = isMatch(s[:len_s-1], p)
			}

			if replace_0 || replace_1 || replace_more {
				return true
			} else {
				return false
			}
		} else {
			// '*'前为普通字符的情况
			replace_0 := false
			replace_1 := false
			replace_more := false
			if 0 == len_s {
				return isMatch(s, p[:len_p-2])
			} else {
				if p_pre == s[len_s-1] {
					// '*'前字符与s末尾字符匹配的情况
					replace_1 = isMatch(s[:len_s-1], p[:len_p-2])
					replace_0 = isMatch(s, p[:len_p-2])
					replace_more = isMatch(s[:len_s-1], p)
				} else {
					// '*'前字符与s末尾字符不匹配的情况
					replace_0 = isMatch(s, p[:len_p-2])
				}
				if replace_0 || replace_1 || replace_more {
					return true
				} else {
					return false
				}
			}

		}
	} else if '.' == p[len_p-1] {
		if 0 == len_s {
			return false
		} else {
			return isMatch(s[:len_s-1], p[:len_p-1])
		}
	} else {
		if 0 == len_s {
			return false
		}
		if p[len_p-1] == s[len_s-1] {
			return isMatch(s[:len_s-1], p[:len_p-1])
		} else {
			return false
		}
	}

	return false

}
