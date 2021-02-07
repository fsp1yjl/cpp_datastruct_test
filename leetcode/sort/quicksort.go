package sort

//less compare, if a < b  return true, else return false
func lessOrEqual(a int, b int)  bool {
	return a <= b
}

func QuickSortPartition(nums []int, left int ,right int ) int {
	pivot := nums[left]
	i:= left+1  // (left, i-1)区间值小于pivot
	j := right // （j
	for {
		//循环后i表示大于哨兵都第一个索引
		for lessOrEqual(nums[i], pivot) {
			i++
			if i >= right {
				break
			}
		}

		// 循环后j表示小于哨兵的第一个索引
		for lessOrEqual(pivot,nums[j]) {
			j--
			if j<= left {
				break
			}
		}

		//相遇或交叉，表示节点已经遍历完毕
		if i >= j {
			break
		}

		// 如果i,j未相遇或者交叉，则进行一次交换
		nums[i], nums[j] = nums[j], nums[i] // swap i,j

	}


	//经过循环处理，j+1到right均大于pivot， j表示最后一个不大于pivot的值
	// swap pivot and j
	nums[left], nums[j] = nums[j],nums[left]
	return j
}

func QuickSort(nums []int)  []int{
	l := len(nums)
	if l <= 1 {
		return nums
	}

	var helper func([]int, int,int)
	helper = func(nums []int, left int, right int) {
		//fmt.Println("LEFT:, RIGHT:", left, right)
		if left >= right {

			return
		}
		//fmt.Println("LEFT:, RIGHT:", left, right)


		div := QuickSortPartition(nums, left, right)

		//fmt.Println("DIV:", div)
		helper(nums, left, div-1)
		helper(nums, div+1, right)
	}

	helper(nums, 0, len(nums)-1)

	return nums



}