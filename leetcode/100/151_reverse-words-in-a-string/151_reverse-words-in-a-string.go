/*
link:https://leetcode.com/problems/reverse-words-in-a-string/

Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.



Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
Example 4:

Input: s = "  Bob    Loves  Alice   "
Output: "Alice Loves Bob"
Example 5:

Input: s = "Alice does not even like bob"
Output: "bob like even not does Alice"

Follow up:

Could you solve it in-place with O(1) extra space?

要求：
 文章的单词反转
 去除首位空格
 多个空格需要处理掉

 原地处理，o(1)的额外空间使用

*/

package main

import "fmt"

func main() {
	s := " Alice does not even like bob"
	fmt.Println(reverseWords(s))
}

/*

思路： 先对数组做空格处理，后进行两次反转，一次整体反转，一次逐个单词对字符反转

提交结果：不是很理想：
Runtime: 12 ms, faster than 10.49% of Go online submissions for Reverse Words in a String.
Memory Usage: 7.2 MB, less than 5.56% of Go online submissions for Reverse Words in a String.
*/
func reverseWords(s string) string {

	l := len(s)
	//先处理字符串多余空格，并trim

	pre := " "
	newS := ""
	for i := 0; i < l; i++ {
		if string(pre) == " " && string(s[i]) == " " {
			continue
		} else {
			newS += string(s[i])
			pre = string(s[i])
		}
	}
	newLen := len(newS)

	if newLen == 0 {
		return ""
	}
	if string(newS[newLen-1]) == " " {
		newS = newS[:newLen-1]
	}
	// 0 1 2 3
	//所有字符串进行一次反转
	ll := len(newS)
	r := []rune(newS)
	for i := 0; i < ll/2; i++ {
		// tmp := newS[i]
		r[i], r[ll-1-i] = r[ll-1-i], r[i]
	}

	// s = string(r)

	//对每个单词再进行一次反转
	left := 0
	cnt := 0
	for left < ll {

		// fmt.Println("cnt", cnt)

		// 注意，这里为了防止越界， left+cnt == ll的判断要放前面
		if left+cnt == ll || r[left+cnt] == ' ' {
			// fmt.Println("here", left)

			start := left
			end := left + cnt - 1
			for start < end {
				r[start], r[end] = r[end], r[start]
				start++
				end--
			}

			left = left + cnt + 1
			cnt = 0

		} else {
			cnt++
			continue
		}

	}

	return string(r)
}
