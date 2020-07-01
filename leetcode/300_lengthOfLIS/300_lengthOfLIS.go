/*
link: https://leetcode.com/problems/longest-increasing-subsequence/

Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.

*/
package main

import "fmt"

func main() {
	//nums := []int{4, 10, 4, 3, 8, 9}
	// [10,9,2,5,3,7,101,18]

	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	l := lengthOfLIS(nums)

	fmt.Println("len :", l)

}

// 核心思路是先依次计算出nums[i]结尾的最大递增子序列，然后一次循环，计算出最大的一个
// 通过缓存历史结果，进行剪枝
func lengthOfLIS(nums []int) int {

	length := len(nums)
	if 0 == length {
		return 0
	}
	l := LIS{
		nums: nums,
		bp:   make(map[int]int),
	}

	l.InitBp()
	max := 0
	for _, b := range l.bp {
		if b > max {
			max = b
		}
	}

	// fmt.Println(l.bp)
	return max
}

type LIS struct {
	nums []int
	bp   map[int]int // bp[i]存放以nums[i]结尾的子数组LIS 长度
}

func (l *LIS) InitBp() {
	l.GetBp(len(l.nums) - 1)
}

// 这里关键是准确找到bp问题的定义，即能准确覆盖问题，又能写出状态方程
func (l *LIS) GetBp(i int) int {

	if v, ok := l.bp[i]; ok {
		return v
	}

	if i == 0 {
		l.bp[i] = 1

		return 1
	} else {

		//tmpLIS 存放nums[i] 依次与bp[0]...bp[i-1]比较后，临时最大lis
		tmpMaxLIS := 1
		curLIS := 1 // curLIS表示与bp[j]比较之后的临时bp[i]
		for j := 0; j < i; j++ {

			subBp := l.GetBp(j)
			//fmt.Println("j:bp", j, subBp)
			if l.nums[i] > l.nums[j] {
				curLIS = subBp + 1
			} else if l.nums[i] == l.nums[j] {
				curLIS = subBp
			} else {
				curLIS = 1
			}

			if tmpMaxLIS < curLIS {
				tmpMaxLIS = curLIS
			}
		}
		l.bp[i] = tmpMaxLIS
		return tmpMaxLIS
	}

}
