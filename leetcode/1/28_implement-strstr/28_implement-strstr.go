/*
link: https://leetcode.com/problems/implement-strstr/


Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().



Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1



*/

package main

import "fmt"

func main() {
	s1 := "hello"
	s2 := "ll"

	fmt.Println(strStr(s1, s2))
}

/*
// 暴力解法：
运行结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
Memory Usage: 2.3 MB, less than 5.02% of Go online submissions for Implement strStr()

*/
func strStr(haystack string, needle string) int {
	l1 := len(haystack)

	l2 := len(needle)

	if 0 == l2 {
		return 0
	}

	if 0 == l1 {
		return -1
	}
	// i表示指针的偏移量
	for i := 0; i < l1; i++ {
		end := i + l2
		// fmt.Println("end:", end, i)
		if end > l1 {
			break
		}
		for j := 0; j < l2; j++ {
			if haystack[j+i] != needle[j] {
				break
			} else {
				if j == l2-1 {
					return i
				}
			}
		}
	}

	return -1
}
