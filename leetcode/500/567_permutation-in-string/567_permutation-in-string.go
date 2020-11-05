/*
link:
https://leetcode.com/problems/permutation-in-string/

Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.



Example 1:

Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input:s1= "ab" s2 = "eidboaoo"
Output: False



*/

package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidboaooo"
	fmt.Println(checkInclusion(s1, s2))

}

/*
思路： 类似滑动窗口，每次从s2截取s1长度的子串，统计各个元素出现的个数，
如果出现不存在的元素或者统计结果超过的元素，则更新窗口起始位置
用一个map记录s1每个元素出现的次数


提交结果：
Runtime: 16 ms, faster than 38.78% of Go online submissions for Permutation in String.
Memory Usage: 4.5 MB, less than 11.22% of Go online submissions for Permutation in String.
*/
func checkInclusion(s1 string, s2 string) bool {

	len1 := len(s1)
	len2 := len(s2)
	if len1 > len2 {
		return false
	}

	m := make(map[string]int)
	for i := 0; i < len1; i++ {
		char := string(s1[i])

		if _, ok := m[char]; ok {
			m[char]++
		} else {
			m[char] = 1
		}
	}

	start := 0
	for start < len2-len1+1 {
		mm := make(map[string]int)

		for k, v := range m {
			mm[k] = v
		}

		cnt := 0

		//这里倒着处理，可以加快处理速度，只要出现不合理的元素，立即从下一位开始新的窗口
		for j := start + len1 - 1; j >= start; j-- {
			char := string(s2[j])

			if _, ok := mm[char]; !ok {
				start = j + 1
				break
			}

			val := mm[char]
			if val < 1 {
				start = j + 1
			} else {
				mm[char]--
				cnt++

			}
		}

		if cnt == len1 {
			return true
		}
	}

	return false
}
