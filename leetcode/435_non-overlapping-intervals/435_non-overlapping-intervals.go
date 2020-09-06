/*
link : https://leetcode.com/problems/non-overlapping-intervals/

参考： https://labuladong.gitbook.io/algo/dong-tai-gui-hua-xi-lie/tan-xin-suan-fa-zhi-qu-jian-tiao-du-wen-ti

*/




package main

import
("fmt"
"sort"
)

func main() {
	intervals := [][]int {{1,2},{2,3},{3,4},{1,3}}
	// intervals := [][]int {{1,2},{2,3},{3,4}}
	// 方法1
	fmt.Println(eraseOverlapIntervals(intervals))

	// 方法2
	fmt.Println(eraseOverlapIntervals_2(intervals))

}


// todo  后续优化，在操作之前，先根据每个元素的right值进行排序后再进行处理
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

func find_end_first(intervals [][]int) int {
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

// 另外一种简洁思路，先按照right 进行排序后，统计不重叠的间隔数
func eraseOverlapIntervals_2(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	// fmt.Println(points)

	l := len(points)

	if l < 2 {
		return 0
	}

	min_right := points[0][1]
	count := 1
	for i := 1; i < l; i++ {
		if points[i][0] >= min_right {
			count++
			min_right = points[i][1]
		}
	}

	return l - count
}