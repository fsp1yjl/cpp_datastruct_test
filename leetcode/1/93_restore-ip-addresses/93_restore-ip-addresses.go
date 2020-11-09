/*

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。



示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
示例 2：

输入：s = "0000"
输出：["0.0.0.0"]
示例 3：

输入：s = "1111"
输出：["1.1.1.1"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-ip-addresses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "0000"
	fmt.Println(restoreIpAddresses(s))
}

/*
 提交结果：
 Runtime: 0 ms, faster than 100.00% of Go online submissions for Restore IP Addresses.
Memory Usage: 2.1 MB, less than 14.06% of Go online submissions for Restore IP Addresses.

*/
func restoreIpAddresses(s string) []string {
	return helper(s, 0, 4)
}

func helper(s string, left int, nums int) []string {
	res := make([]string, 0)

	// r := []rune(s) // string to rune slice
	l := len(s)
	span := l - left

	//如果 字符长度不在nums 到 3*nums之间，则非法
	if span > 3*nums || span < nums {
		return res
	}

	//表示最后一位ip
	if nums == 1 {
		//多于一位，且首字符为0,非法返回空
		if span > 1 && string(s[left]) == "0" {
			return res
		}
		num, _ := strconv.Atoi(s[left:])
		if num < 0 || num > 255 {
			return res
		} else {
			return append(res, s[left:])
		}
	}

	// 非最后一位ip， 则有多种可选情况

	//取一位数字
	c1 := string(s[left])
	sub := helper(s, left+1, nums-1)
	if len(sub) != 0 {
		for i := 0; i < len(sub); i++ {
			res = append(res, c1+"."+sub[i])
		}
	}

	// 如果首位不为0，则可以考虑取2位或者3位作为数字,否则直接返回
	if string(s[left]) == "0" {
		return res
	}

	// 取两位作为数字
	if left+1 < l-1 {
		c2 := s[left : left+2]
		sub := helper(s, left+2, nums-1)
		if len(sub) != 0 {
			for i := 0; i < len(sub); i++ {
				res = append(res, c2+"."+sub[i])
			}
		}
	}

	// 取3位作为数字
	if left+1 < l-2 {
		c3 := s[left : left+3]

		num, _ := strconv.Atoi(c3)
		if num < 0 || num > 255 {
			return res
		}

		sub := helper(s, left+3, nums-1)
		if len(sub) != 0 {
			for i := 0; i < len(sub); i++ {
				res = append(res, c3+"."+sub[i])
			}
		}
	}
	return res
}
