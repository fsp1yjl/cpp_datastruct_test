/*
link: https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.

//说明，找到字符串数组的最长公共前缀，所有字符串为小写英文字母


*/

package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

}

/*

提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.4 MB, less than 52.02% of Go online submissions for Longest Common Prefix.

*/
func longestCommonPrefix(strs []string) string {
	prefix := ""
	l := len(strs)
	if 0 == l {
		return prefix
	}

	minLen := len(strs[0])
	for i := 1; i < l; i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
I:
	for i := 0; i < minLen; i++ {
		c := strs[0][i]
		for j := 1; j < l; j++ {
			if strs[j][i] != c {
				break I
			}
		}
		prefix += string(c)
	}
	return prefix
}
