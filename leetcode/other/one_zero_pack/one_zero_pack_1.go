/*

给你一个可装载重量为 W 的背包和 N 个物品，
每个物品有重量和价值两个属性。
其中第 i 个物品的重量为 wt[i]，价值为 val[i]，
现在让你用这个背包装物品，最多能装的价值是多少？


N = 3, W = 4
wt = [2, 1, 3]
val = [4, 2, 3]
*/

package main

import "fmt"

var (
	// wt  = []int{2, 1, 3}
	// val = []int{4, 2, 3}
	wt  = []int{5, 4, 7, 2, 6}
	val = []int{12, 3, 10, 3, 6}
)

func main() {
	const w int = 15

	n := len(wt)

	dpMem := make([][]int, n)

	for i := 0; i < n; i++ {
		dpMem[i] = make([]int, w+1)
	}

	//初始化第一行
	for i := wt[0]; i <= w; i++ {
		dpMem[0][i] = val[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= w; j++ {
			dp_1 := dpMem[i-1][j]
			dp_2 := 0
			if j >= wt[i] { // 这里一定要注意不要忘记等于的情况
				dp_2 = dpMem[i-1][j-wt[i]] + val[i]
			}

			if dp_1 > dp_2 {
				dpMem[i][j] = dp_1
			} else {
				dpMem[i][j] = dp_2
			}
		}
	}

	fmt.Println("mem:", dpMem)

	fmt.Println("result:", dpMem[n-1][w])

	// ret := dp(dpMem, n-1, w)

	// fmt.Println("ret:", ret)

}
