/*
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
示例 4：

输入：nums = [1]
输出：[1]


实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
示例 4：

输入：nums = [1]
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 2, 0, 2, 3, 2, 0}
	// nums := []int{1, 3, 2}

	nextPermutation(nums)

	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	/*
		思路，从后向前，找到递减序列的结束点j，在递减序列中互相交换无法产生更大到值
		所以可以确定发生交换一定是0到j-1 和j到最后的两个区间各选一个节点。
		为了影响最小，就是拿前面的最低位j-1 和 后面区间中大于nums[j-1]的第一个数进行交换

		交换之后，再对后面区间进行调整，做成一个单调递增区间

	*/
	l := len(nums)
	if l == 1 {
		return
	}

	for i := l - 1; i > 0; i-- {
		//

		if nums[i] <= nums[i-1] {
			continue
		} else {

			j := l - 1
			for nums[j] <= nums[i-1] {
				j--
			}

			// fmt.Println("i,j:", i-1, j)
			//因为后一个区间是递减的，这里对j将代表后一个区间第一个大于i-1的数字
			nums[i-1], nums[j] = nums[j], nums[i-1]
			n := nums[i:]
			sort.Slice(n, func(i, j int) bool {
				// fmt.Println("dddddd")

				return n[i] < n[j]
			})

			return
		}

	}

	for i := 0; i < l/2; i++ {
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}

}
