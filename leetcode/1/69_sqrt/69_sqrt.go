/*

link: 69. Sqrt(x)
Easy

1574

2065

Add to List

Share
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.


*/

package main

import "fmt"

func main() {
	fmt.Println(mySqrt(1))
}

/*

使用二分法

提交结果：
Runtime: 4 ms, faster than 52.42% of Go online submissions for Sqrt(x).
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Sqrt(x).

*/
func mySqrt(x int) int {

	left := 0
	right := x
	for right > left {
		// bug1 这里使用+1 防止left+1 = right时出现死循环
		mid := (left + right + 1) / 2
		if mid*mid <= x {
			left = mid
		} else {
			right = mid - 1
		}
		fmt.Println("left,right:", left, right)
	}
	return left
}
