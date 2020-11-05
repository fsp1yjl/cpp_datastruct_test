/*
linK :https://leetcode.com/problems/longest-substring-without-repeating-characters/

无重复字符的最长子串

Given a string s, find the length of the longest substring without repeating characters.

 

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
Example 4:

Input: s = ""
Output: 0


思路：
滑动窗口 
https://www.youtube.com/watch?v=LupZFfCCbAU
*/

package main 

import (
	"fmt"
)
func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}
/*
用一个哈希表存储right元素上一次出现的位置x，下次移动滑动窗口left从x+1开始

提交结果：
Runtime: 8 ms, faster than 72.47% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 3.8 MB, less than 9.00% of Go online submissions for Longest Substring Without Repeating Characters.

*/
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l <= 1 {
		return l
	}

	left:=0
	right := 0 
	res := 1

	m := make(map[string]int)



	for right < l {
		char  := string(s[right])
		if index, ok := m[char];ok {
			if index >= left {
				//滑动窗口内存在right对应元素，需要缩小滑动窗口，并进行一次长度统计
				// fmt.Println("left, right", left, right)
				curLength := right - left 
				if curLength > res {
					res = curLength
				}
				left = index+1
				m[char] = right
				right++
				continue
			}
		} 

		curLength := right - left + 1
		if curLength > res {
			res = curLength
		}
		m[char] = right
		right++
	}


	return res 
}