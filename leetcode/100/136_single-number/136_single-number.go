/*
link: https://leetcode.com/problems/single-number/

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

Follow up: Could you implement a solution with a linear runtime complexity and without using extra memory?



Example 1:

Input: nums = [2,2,1]
Output: 1
Example 2:

Input: nums = [4,1,2,1,2]
Output: 4
Example 3:

Input: nums = [1]
Output: 1
*/

package main

import "fmt"

func main() {
	n := []int{2,2,3}

	fmt.Println(singleNumber(n))
}
/*
其他元素都成对出现，仅有一个元素单独出现，则考虑进行异或处理, 因为相同元素异或后会置0
single number与0进行异或后，还是数字本身

提交结果：
Runtime: 16 ms, faster than 53.26% of Go online submissions for Single Number.
Memory Usage: 6 MB, less than 71.88% of Go online submissions for Single Number.
*/
func singleNumber(nums []int) int {

 l:= len(nums)
 ret := 0
 for i:=0; i<l; i++ {
	ret ^= nums[i] 
 }

 return ret
}