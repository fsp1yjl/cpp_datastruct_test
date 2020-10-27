/*
leetcode 58

给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。



示例:

输入: "Hello World"
输出: 5


*/

package main

import "fmt"

func main() {
	s := "haha hehee"

	fmt.Println(lengthOfLastWord(s))
}

/*
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Length of Last Word.

*/
func lengthOfLastWord(s string) int {
	l := len(s)

	if 0 == l {
		return 0
	}
	start := false
	startIndex := -1
	endIndex := -1 // 默认空格开始的索引位置

	for i := l - 1; i >= 0; i-- {
		c := s[i]
		if !start && c != byte(' ') {
			start = true
			startIndex = i
		}

		fmt.Println("startIndex:", startIndex)

		if start {
			if c == byte(' ') {
				endIndex = i
				fmt.Println("end:", endIndex)

				break
			}
		}
	}

	if start {
		return startIndex - endIndex
	} else {
		return 0
	}
}
