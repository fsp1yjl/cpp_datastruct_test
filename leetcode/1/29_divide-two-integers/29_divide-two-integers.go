/*
link:  https://leetcode.com/problems/divide-two-integers/

Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.

Note:

Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For this problem, assume that your function returns 231 − 1 when the division result overflows.


Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = truncate(3.33333..) = 3.
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = truncate(-2.33333..) = -2.
Example 3:

Input: dividend = 0, divisor = 1
Output: 0
Example 4:

Input: dividend = 1, divisor = 1
Output: 1


*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(6 << 3)
	fmt.Println(divide(2147483647, 1))
	// fmt.Println("------------------")
	// fmt.Println(divide2(-31, 2))
}

/*
// 题目要求不能使用乘法/除法/取模操作，这里使用减法先写一次， 提交结果不理想

todo: 考虑使用位运算，每次移动一位去减少减法处理，先找到除数第一个移位后大于被除数的项目

提交结果：
Runtime: 2512 ms, faster than 11.72% of Go online submissions for Divide Two Integers.
Memory Usage: 2.5 MB, less than 21.88% of Go online submissions for Divide Two Integers.
*/
func divide2(dividend int, divisor int) int {
	max := math.MaxInt32
	min := math.MinInt32

	// fmt.Println("max:", max)

	// fmt.Println("min:", min)

	negtiveFlag := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		negtiveFlag = true
	}
	res := 0
	if 0 == dividend {
		return res
	}

	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor
	}

	// bug1 不要忽略等于的情况
	for dividend >= divisor {
		res++
		dividend = dividend - divisor
	}

	// bug2 不要忽略数值越界的情况
	if negtiveFlag {
		res = 0 - res
		if res < min {
			return min
		} else {
			return res
		}
	} else {
		if res > max {
			return max
		} else {
			return res
		}
	}

}

/*
这里先使用移位的方式，减少执行减法操作的次数

进一步: 可以考虑将所有的减法操作都替换成移位操作，进一步提高运算效率
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Divide Two Integers.
Memory Usage: 2.5 MB, less than 21.88% of Go online submissions for Divide Two Integers.
Next challenges:
*/
func divide(dividend int, divisor int) int {
	max := math.MaxInt32
	min := math.MinInt32

	// fmt.Println("max:", max)

	// fmt.Println("min:", min)

	negtiveFlag := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		negtiveFlag = true
	}
	res := 0
	if 0 == dividend {
		return res
	}

	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor
	}

	if dividend < divisor {
		return 0
	}

	move := 0
	dd := divisor
	for dividend >= dd {
		move++
		dd = dd << 1
		// fmt.Println("ddd:", dd)
	}

	if move > 0 {
		dd = dd >> 1
		move--
		res = 1 << move
	} else {
		// 这里需要考虑 除数绝对值大于被除数都情况
		dd = 0
	}

	// fmt.Println("res move:", move)
	// fmt.Println("res move:", dd)

	dividend = dividend - dd
	// dd == dd >> 1
	// for dividend > 0 && dividend >= dd {
	// 	move--
	// }
	fmt.Println("after dividend:", dividend, divisor)

	if dividend >= divisor {
		res += divide(dividend, divisor)
	}

	// for dividend >= divisor {
	// 	res++
	// 	dividend = dividend - divisor
	// }

	// bug2 不要忽略数值越界的情况
	if negtiveFlag {
		res = 0 - res
		if res < min {
			return min
		} else {
			return res
		}
	} else {
		if res > max {
			return max
		} else {
			return res
		}
	}

}
