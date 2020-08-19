package main

import (
	"fmt"
)

func main() {
	nums := []int {3, 5, 2, 19, 6, 9}
	fmt.Println(select_sort(nums))
	fmt.Println(bubble_sort(nums))
	fmt.Println(insert_sort(nums))

}

func select_sort(nums []int ) []int {

	l := len(nums)

	for i := 0; i < l; i++ {
		min_index := i
		min := nums[i]

		for j :=i+1; j < l; j++ {
			if nums[j] < min { //注意这里每次都是和暂时的最小值做对比
				min_index = j
				min = nums[j]
			}
		}
		if i != min_index {
			nums[min_index] = nums[i]
			nums[i] = min
		}
	}

	return nums

}


func bubble_sort(nums []int) []int {
	l := len(nums)

	for i:=0; i < l; i++ {
		for j:=i+1; j < l; j++ {
			if nums[j] < nums[j-1] {
				temp := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = temp
			}
		}
	}

	return nums
}


func insert_sort(nums []int) []int {
	l := len(nums)

	for i := 1; i < l; i++ {
		for j:=0; j < i; j++ {
			if nums[i] < nums[j] {
				tmp := nums[i]
				for t := i-1; t <= j; t-- {
					nums[t+1] = nums[t]
				}
				nums[j] = tmp
			}
		}
	}
	return nums
}