/*

link:https://leetcode.com/problems/letter-combinations-of-a-phone-number/

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

package main

import "fmt"

func main() {
	str := "23"

	fmt.Println(letterCombinations(str))
}

/*
运行结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
Memory Usage: 2 MB, less than 68.38% of Go online submissions for Letter Combinations of a Phone Number.

思路：
遍历字符串，保存已经遍历部分对笛卡尔积数组

*/
func letterCombinations(digits string) []string {
	l := len(digits)
	if 0 == l {
		return []string{}
	}

	return dfs(digits, 0, []string{})
}

func mapping(c string) []string {
	res := []string{}
	switch c {
	case "2":
		res = []string{"a", "b", "c"}
	case "3":
		res = []string{"d", "e", "f"}
	case "4":
		res = []string{"g", "h", "i"}
	case "5":
		res = []string{"j", "k", "l"}
	case "6":
		res = []string{"m", "n", "o"}
	case "7":
		res = []string{"p", "q", "r", "s"}
	case "8":
		res = []string{"t", "u", "v"}
	case "9":
		res = []string{"w", "x", "y", "z"}
	}
	return res
}

func dfs(digits string, index int, pre []string) []string {
	l := len(digits)
	if index >= l {
		return pre
	}

	d := string(digits[index])
	chars := mapping(d)
	lenPre := len(pre)
	if lenPre == 0 {
		pre = chars
	} else {
		tmp := []string{}
		for j := 0; j < lenPre; j++ {
			s := pre[j]
			for k := 0; k < len(chars); k++ {
				tmp = append(tmp, s+chars[k])
			}

		}
		pre = tmp
	}

	return dfs(digits, index+1, pre)

}
