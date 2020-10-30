/*

link: https://leetcode.com/problems/add-binary/

Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"


Constraints:

Each string consists only of '0' or '1' characters.
1 <= a.length, b.length <= 10^4
Each string is either "0" or doesn't contain any leading zero.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1010"
	b := "1011"
	// 10101

	// 11101

	fmt.Println(addBinary(a, b))
}

/*
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
Memory Usage: 2.6 MB, less than 58.24% of Go online submissions for Add Binary.



*/
func addBinary(a string, b string) string {
	l_last := len(a) - 1
	r_last := len(b) - 1

	up := 0 // up代表进位，注意每次计算后更新进位

	s := ""
	l := l_last
	r := r_last

	for l >= 0 && r >= 0 {
		ia, _ := strconv.Atoi(string(a[l]))
		ib, _ := strconv.Atoi(string(b[r]))

		// fmt.Println("ia:", ia)
		// fmt.Println("ib:", ib)

		sum := ia + ib + up

		// fmt.Println("sum:", sum)

		if sum/2 > 0 {
			up = 1
		} else {
			up = 0
		}

		t := sum % 2

		if 0 == t {
			s = "0" + s
		} else {
			s = "1" + s
		}

		l--
		r--
	}

	for l >= 0 {
		i_a, _ := strconv.Atoi(string(a[l]))
		sum := i_a + up

		if sum/2 > 0 {
			up = 1
		} else {
			up = 0
		}

		t := sum % 2

		if 0 == t {
			s = "0" + s
		} else {
			s = "1" + s
		}
		l--
	}

	for r >= 0 {

		i_b, _ := strconv.Atoi(string(b[r]))
		sum := i_b + up

		if sum/2 > 0 {
			up = 1
		} else {
			up = 0
		}

		t := sum % 2

		if 0 == t {
			s = "0" + s
		} else {
			s = "1" + s
		}
		r--
	}

	if up > 0 {
		s = "1" + s
	}
	return s
}
