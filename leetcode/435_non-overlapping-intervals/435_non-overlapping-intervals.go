/*
link : https://leetcode.com/problems/non-overlapping-intervals/

参考： https://labuladong.gitbook.io/algo/dong-tai-gui-hua-xi-lie/tan-xin-suan-fa-zhi-qu-jian-tiao-du-wen-ti

*/

/*
当前结果：
Runtime: 356 ms, faster than 5.14% of Go online submissions for Non-overlapping Intervals.
Memory Usage: 8.1 MB, less than 5.14% of Go online submissions for Non-overlapping Intervals.
Next challenges:

/*

// todo  后续优化，在操作之前，先根据每个元素的right值进行排序后再进行处理

package main

import "fmt"

func main() {
	intervals := [][]int {{1,2},{2,3},{3,4},{1,3}}
	// intervals := [][]int {{1,2},{2,3},{3,4}}
	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	
	
	l := len(intervals)
	erase := 0


	if l < 2 {
		return erase
	}


	min_index  := find_end_first(intervals)
	// min_left := intervals[min_index][0]
	min_right := intervals[min_index][1]
	next_intervals := make([][]int,0)
	for i:=0; i < l ; i++ {
		if intervals[i][0] < min_right {
			if i != min_index {
				erase++
			} 
		}  else {
			next_intervals = append(next_intervals, intervals[i])
		}
	}

	return erase + eraseOverlapIntervals(next_intervals)


}

func find_end_first(intervals [][]int)  int {
	min := intervals[0][1]
	min_index := 0
	l := len(intervals)
	for i:=1; i < l ; i++ {
		if intervals[i][1] < min {
			min = intervals[i][1]
			min_index =  i
		}
	}
	return min_index
}