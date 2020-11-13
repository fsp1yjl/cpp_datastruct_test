/*
link: https://leetcode.com/problems/kth-largest-element-in-an-array/

Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.

*/

package main

import "fmt"

func main() {
	// nums := []int{3, 2, 1, 5, 6, 4}
	nums := []int{2, 1}
	fmt.Println(findKthLargest(nums, 2))
}

/*

利用类似快速排序的方式去处理，哨兵的位置即为最终第k大的数据

这里构建一个降序排列的数组

提交结果不理想：
Runtime: 24 ms, faster than 6.81% of Go online submissions for Kth Largest Element in an Array.
Memory Usage: 4.3 MB, less than 9.26% of Go online submissions for Kth Largest Element in an Array.
*/
func findKthLargest(nums []int, k int) int {
	l := len(nums)
	// index := k - 1

	// return helper(nums, 0, l-1, index)

	return partition(nums, 0, l-1, k-1)

}

/*
另外一种利用哨兵的快拍定位方式，提交结果更理想一些：

Runtime: 12 ms, faster than 31.87% of Go online submissions for Kth Largest Element in an Array.
Memory Usage: 4.3 MB, less than 22.25% of Go online submissions for Kth Largest Element in an Array.

进一步的优化思路，k与数组长度的二分之一进行比较，如果k大于二分之一，则将问题转化为求第l-k小的元素
*/
func partition(nums []int, left int, right int, index int) int {
	if left == right && left == index {
		return nums[left]
	}

	storeIndex := left //表示第一个未确定的元素位置
	pivot := nums[right]
	for i := left; i < right; i++ {
		// fmt.Println("i---", i)
		// fmt.Println("num,pivot", nums[i], pivot)

		if nums[i] > pivot {
			tmp := nums[i]
			nums[i] = nums[storeIndex]
			nums[storeIndex] = tmp
			storeIndex++ //每发现一个大于哨兵的元素，未确定大小的元素索引就前进一位
			// fmt.Println("nums---:", nums)
		}
	}

	//循环结束后的storeindex就表示第一个大于哨兵的元素位置 （这个位置要么小于right，要么和right重叠）
	tmp := nums[storeIndex]
	nums[storeIndex] = pivot
	nums[right] = tmp

	// fmt.Println("nums:", nums)
	// fmt.Println("store-index-----:", storeIndex)

	if storeIndex == index {
		return nums[storeIndex]
	} else if storeIndex < index {
		// fmt.Println("herer----------")

		return partition(nums, storeIndex+1, right, index)
	} else {
		// fmt.Println("herer")
		return partition(nums, left, storeIndex-1, index)
	}

}

func helper(nums []int, start int, end int, index int) int {
	divider := nums[start]
	left := start + 1
	right := end
	// fmt.Println("left,right:", left, right)

	// 注意这里要带=号，处理类似[2,1] 1，这样的情况，此时初始哨兵为0， start,end均为1
	for left <= right {

		if nums[left] >= divider {
			left++
		}

		if nums[right] <= divider {
			right--
		}

		if left < right && nums[left] < divider && nums[right] > divider {
			tmp := nums[left]
			nums[left] = nums[right]
			nums[right] = tmp

		}
		// swap tmp left,right

	}

	nums[start] = nums[right]
	nums[right] = divider

	//右侧正好有k个元素，则返回右侧最小值
	if right == index {
		//k为右侧最小值
		return nums[right]
		// return min
	} else if right < index {
		return helper(nums, right+1, end, index)
	} else {
		// 0,1,2,3,4,5,6,7
		// k=
		// left=

		return helper(nums, start, right-1, index)
	}

}
