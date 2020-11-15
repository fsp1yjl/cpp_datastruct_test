/*
link: https://leetcode.com/problems/merge-intervals/

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.



Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.


Constraints:

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104

*/

package main

import (
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {9, 1}, {2, 6}, {8, 10}, {15, 18}}
	merge(intervals)
}

/*
提交结果： Runtime: 8 ms, faster than 92.06% of Go online submissions for Merge Intervals.
Memory Usage: 4.7 MB, less than 83.73% of Go online submissions for Merge Intervals.

*/
func merge(intervals [][]int) [][]int {
	//先对数组按左边元素排序

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// fmt.Println("slice:", intervals)

	res := make([][]int, 0)

	l := len(intervals)
	if l <= 1 {
		return intervals
	}

	pre := intervals[0]

	for i := 1; i < l; i++ {
		preRight := pre[1]
		cur := intervals[i]
		curLeft := cur[0]
		curRight := cur[1]

		if curLeft <= preRight {
			//需要进行合并,
			if curRight > preRight {
				pre[1] = curRight
			}
		} else {
			//不需要合并，将pre入结果
			res = append(res, pre)
			pre = cur
		}
	}

	res = append(res, pre)
	return res

}
