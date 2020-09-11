/*
link : https://leetcode.com/problems/permutations/

Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {

	res := combin([]int{}, nums)
	return res
}

func combin(visited []int, unvisited []int) [][]int {

	if len(unvisited) == 1 {
		prefix := append(visited, unvisited...)
		return [][]int{prefix}
	} else {
		arr := make([][]int, 0)
		for i, v := range unvisited {
			newPrefix := append(visited, v)
			newUnvisited := make([]int, 0)
			newUnvisited = append(newUnvisited, unvisited[:i]...)
			newUnvisited = append(newUnvisited, unvisited[i+1:]...)

			subSet := combin(newPrefix, newUnvisited)

			for _, once := range subSet {
				arr = append(arr, once)
			}
		}
		return arr
	}

}
