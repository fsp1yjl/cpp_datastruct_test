/**

link: https://leetcode.com/problems/pascals-triangle-ii/

Given an integer rowIndex, return the rowIndexth row of the Pascal's triangle.

Notice that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Follow up:

Could you optimize your algorithm to use only O(k) extra space?



Example 1:

Input: rowIndex = 3
Output: [1,3,3,1]
Example 2:

Input: rowIndex = 0
Output: [1]
Example 3:

Input: rowIndex = 1
Output: [1,1]

*/

package main

import "fmt"

func main() {

	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	pre := make([]int, rowIndex+1)

	pre[0] = 1

	if rowIndex == 0 {
		return pre
	}

	for i := 1; i <= rowIndex; i++ {
		pre[0] = 1
		pre[i] = 1
		// 这里倒序处理是为了节省空间
		for j := i - 1; j > 0; j-- {
			pre[j] = pre[j-1] + pre[j]
		}

	}

	return pre
}
