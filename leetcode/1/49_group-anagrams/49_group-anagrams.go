/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(1 + 'a')
	strs := []string{"hello", "ddd", "hlloe"}
	groupAnagrams(strs)
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	l := len(strs)
	if 0 == l {
		return res
	}
	m := make(map[string]int)
	for i := 0; i < l; i++ {
		summary := make([]int, 26)
		s := []rune(strs[i])
		for j := 0; j < len(s); j++ {
			c := s[j]
			index := c - 'a'
			summary[index]++
		}
		// fmt.Println(summary)
		key := ""
		for k := 0; k < 26; k++ {
			if summary[k] != 0 {
				key += string(k+'a') + strconv.Itoa(summary[k])
			}
		}

		index, ok := m[key]
		if ok {
			res[index] = append(res[index], strs[i])
		} else {
			l := len(res)
			m[key] = l
			tmp := []string{strs[i]}
			res = append(res, tmp)
		}
	}

	fmt.Println("res:", res)
	return res
}
