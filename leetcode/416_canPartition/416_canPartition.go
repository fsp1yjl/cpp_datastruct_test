/*
Given a non-empty array containing only positive integers,
find if the array can be partitioned into two subsets
such that the sum of elements in both subsets is equal.

Note:
Each of the array element will not exceed 100.
The array size will not exceed 200.


Example 1:
Input: [1, 5, 11, 5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11]
*/

package main

import "fmt"

func main() {
	nums := []int{1, 5, 4, 7, 5}

	res := canPartition(nums)
	fmt.Println("res:", res)

}

func canPartition(nums []int) bool {
	
	sum := 0
	// l := len(nums)
	for _, v := range nums {
		sum += v
	}
	if 0 != sum%2  {
		return false 
	}

	W := sum/2
	dp := make(map[int]bool)
	dp[0] = true

	
	
	// dp[i]表示使用当前可选的重量，是否可以刚好将背包装满
	// 最初dp这个map中只有dp[0]一个key, 依次使用1，5， 4，7，5这几个重量对dp进行更新
	for _, num := range nums {
		for i:=W; i >= num; i-- {
		// 每次重量从后向前是因为小于选中背包质量的dp 不需要再本次循环中更新
			if v, ok := dp[i-num]; ok {
				dp[i] = v
			}
		}
	}
	
	fmt.Println("final:", dp)
	if v, ok := dp[W]; ok {
		return v
	} else {
		return false 
	}

 
}