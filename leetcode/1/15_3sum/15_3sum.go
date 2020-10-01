/*
link: https://leetcode.com/problems/3sum/

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Example 2:

Input: nums = []
Output: []
Example 3:

Input: nums = [0]
Output: []

*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}
	fmt.Println(threeSum(nums))
}
/*
// 解法1: 思路，先对数组排序，依次确定左右值，查找中间值是否存在，同时注意删除重复对解
// 提交结果显示运行效率很差

Runtime: 1148 ms, faster than 9.93% of Go online submissions for 3Sum.
Memory Usage: 7 MB, less than 43.57% of Go online submissions for 3Sum.
*/
func threeSum(nums []int) [][]int {

	res := [][]int{}
	l := len(nums)
	if l < 3 {
		return res 
	}

	// m := make(map[int]int)


	//数组排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// fmt.Println(nums)


	right := l-1 
	for right >=2 {
		left := 0
		for left < right  {
			l := nums[left] // 代表最小值
			r := nums[right] // 代表最大值
	
	
			// 如果最左大于0，或者最右小于0，则不可能再有解，结束循环
			if l>0 || r < 0 {
				break

			}
	
			sum := l +r 
	
			target := 0 - sum
			// fmt.Println("left, right:, target", l, r,target)

			if target < l || target > r  {
				left++
				continue
			}
			
			for i:= left+1; i< right; i++ {
				// fmt.Println("hello ", target, nums[i])
				dupFlag := true

				if nums[i] == target {
					// fmt.Println("hello:num:", nums[i])
					tmp:=[]int{l, target, r}
				
					if len(res) != 0 {
						for k:=0; k < len(res); k++ {
							pre := res[k]

							if pre[0]== l && pre[1] == target  { // 这里是为了防止解重复
								dupFlag = false
								break
							}
						}
					

						
						// fmt.Println("pre:", pre)
						// fmt.Println("tmp:", tmp)

					} 
					
					if dupFlag {
						res = append(res, tmp)

					}
					
					break
				}

				if nums[i] > target {
					break
				}
			}
	
			left++
	
		}
		right--
	}


	return res
}