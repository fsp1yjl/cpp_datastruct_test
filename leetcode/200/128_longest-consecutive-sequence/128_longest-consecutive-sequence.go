/*

link: https://leetcode.com/problems/longest-consecutive-sequence/

参考： https://leetcode.com/problems/longest-consecutive-sequence/


Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

Follow up: Could you implement the O(n) solution?



Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
*/

package main

import "fmt"

func main() {
	nums := []int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {

	res := 0
	l := len(nums)

	if 0 == l {
		return res
	}

	m := make(map[int]int)

	for i := 0; i < l; i++ {
		cur := nums[i]

		leftFlag := false
		curFlag := false
		rightFlag := false

		if _, ok := m[cur-1]; ok {
			leftFlag = true
		}
		if _, ok := m[cur+1]; ok {
			rightFlag = true
		}
		if _, ok := m[cur]; ok {
			curFlag = true
		}

		// 已经存在则不再处理
		if curFlag {
			continue
		}

		//如果左右均不存在
		if !leftFlag && !rightFlag {
			m[cur] = 1
			continue
		}

		//如果仅有左
		if leftFlag && !rightFlag {
			span := m[cur-1]

			m[cur] = span + 1
			m[cur-span] = m[cur]
			continue
		}

		//如果仅有右
		if !leftFlag && rightFlag {
			span := m[cur+1]
			m[cur] = span + 1
			m[cur+span] = m[cur]
			continue
		}

		//如果左右均存在
		if leftFlag && rightFlag {
			leftSpan := m[cur-1]
			rightSpan := m[cur+1]

			newSpan := leftSpan + rightSpan + 1
			m[cur-leftSpan] = newSpan
			m[cur+rightSpan] = newSpan

			// 注意，这里将m[cur]赋值为1很关键，防止重复出现的元素造成错误结果
			m[cur] = 1

			continue
		}

	}

	for _, v := range m {
		if v > res {
			res = v
		}
	}

	fmt.Println("m::", m)

	return res

}
