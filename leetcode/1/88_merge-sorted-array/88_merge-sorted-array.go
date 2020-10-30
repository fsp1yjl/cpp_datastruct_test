/*
link: https://leetcode.com/problems/merge-sorted-array/

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]

*/
package main

func main() {

}

/*
题目的意思是将nums2的元素也合并到nums1的剩余空间中
为了不使用额外空间，可以从倒数第一个元素依次往前处理

提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Sorted Array.
Memory Usage: 2.3 MB, less than 6.43% of Go online submissions for Merge Sorted Array.

*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	end := m + n - 1
	l := m - 1
	r := n - 1
	for l >= 0 && r >= 0 {
		if nums1[l] > nums2[r] {
			nums1[end] = nums1[l]
			l--
		} else {
			nums1[end] = nums2[r]
			r--
		}
		end--
	}

	for l >= 0 {
		nums1[end] = nums1[l]
		l--
		end--
	}

	for r >= 0 {
		nums1[end] = nums2[r]
		r--
		end--
	}

}
