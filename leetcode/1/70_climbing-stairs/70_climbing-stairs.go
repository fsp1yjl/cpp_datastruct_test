/*
link:
https://leetcode.com/problems/climbing-stairs/

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


Constraints:

1 <= n <= 45


*/

package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}

/*
本题即是基本的 斐波那契数列

提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Climbing Stairs.

*/
func climbStairs(n int) int {

	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		m[i] = -1
	}

	m[1] = 1
	m[2] = 2

	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]

		// fmt.Println(m)
	}

	return m[n]
}
