/*

link: https://leetcode.com/problems/palindrome-number

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Follow up:

Coud you solve it without converting the integer to a string?
*/

// 题目中说了，尽量尝试不将数组转换为字符串的解法

/*
运行结果：
Runtime: 28 ms, faster than 24.36% of Go online submissions for Palindrome Number.
Memory Usage: 5.2 MB, less than 82.77% of Go online submissions for Palindrome Number.

*/
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(1000021))
}

//
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	div := 1
	for x/div > 9 {
		div *= 10
	}

	for x > 0 {

		left := x / div
		right := x % 10
		// fmt.Println("left:,right:", left, right)
		if left != right {
			return false
		} else {
			x = (x - right - left*div) / 10
			// 下次运算，除数自动缩小100，因为头尾各去掉了一位
			// 同时，如果每次基于x重新计算div，注意1011这样的数字检测问题 ，要考虑1011收缩为01后
			div = div / 100
			// fmt.Println("x::", x)
		}

	}

	return true
}
