/*

link: https://leetcode.com/problems/plus-one/
Given a non-empty array of digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Example 3:

Input: digits = [0]
Output: [1]

*/

package main

import "fmt"

func main() {
	d := []int{4, 3, 2, 1}
	fmt.Println(plusOne(d))
}

/*
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Plus One.

*/

func plusOne(digits []int) []int {
	l := len(digits)

	d := make([]int, l, l+1)

	// fmt.Println(len(d))
	up := 1 // up用来表示进位
	for i := l - 1; i >= 0; i-- {
		tmp := up + digits[i]
		if tmp > 9 {
			up = 1
		} else {
			up = 0
		}

		d[l-i-1] = tmp % 10
	}
	if up > 0 {
		d = append(d, up)
	}

	l = len(d)
	// fmt.Println(d)
	//数组反转
	for i := 0; i <= (l-1)/2; i++ { //  bug1,注意这里是l-1,而不是l
		tmp := d[i]

		d[i] = d[l-i-1]
		d[l-i-1] = tmp
	}
	return d
}