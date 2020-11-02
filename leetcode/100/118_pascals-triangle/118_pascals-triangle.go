/*

link:https://leetcode.com/problems/pascals-triangle/

Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

/*
本题即为构造一个杨辉三角，
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Pascal's Triangle.
*/
func generate(numRows int) [][]int {

	arr := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		// fmt.Println("i:", i)
		arr[i] = make([]int, i+1)
		arr[i][0] = 1
		arr[i][i] = 1
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
	}

	return arr
}
