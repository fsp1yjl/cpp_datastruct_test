/*
https://leetcode.com/problems/multiply-strings/

Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.



Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"


Constraints:

1 <= num1.length, num2.length <= 200 (不要考虑字符串为空的情况)
num1 and num2 consist of digits only.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
*/

package main

import (
	"strconv"
)

func main() {
	num1 := "123"
	num2 := "456"

	multiply(num1, num2)

	// a := nums1[1] - '0'
	// b := a * 3
	// fmt.Println("a:", b)

}

func multiply(num1 string, num2 string) string {
	len1 := len(num1)

	len2 := len(num2)

	arr := make([]int, len1+len2)

	// 先处理低位，所以要倒序处理
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			c1, _ := strconv.Atoi(string(num1[i]))
			c2, _ := strconv.Atoi(string(num2[j]))
			tmp := arr[i+j+1] + c1*c2
			arr[i+j+1] = tmp % 10
			arr[i+j] += tmp / 10
		}
	}

	str := ""
	flag := false

	for i := 0; i < len1+len2; i++ {
		if !flag {
			if arr[i] == 0 {
				continue
			} else {
				flag = true
			}
		}

		str += strconv.Itoa(arr[i])

	}

	// bug1 : 考虑结果为0的情况
	if str == "" {
		str = "0"
	}
	// fmt.Println("str::", str)
	return str

}
