/*
https://leetcode-cn.com/problems/max-area-of-island/

题目说明：
一个0，1组成的二维举证， 求连城一片的1对应的最大面积，每个格子面积为1

思路：
对为1的元素进行dfs, 已经遍历过的为1节点在统计之后值置0.

注意：入栈时元素值需要立即置0，防止重复入站

提交结果：
Runtime: 12 ms, faster than 77.17% of Go online submissions for Max Area of Island.
Memory Usage: 5.7 MB, less than 29.35% of Go online submissions for Max Area of Island.

*/

package main

import "fmt"

func main() {
	// grid := [][]int{
	// 	{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	// 	{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
	// 	{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1}}

	// fmt.Println(len(grid))
	fmt.Println(maxAreaOfIsland(grid))

}

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	m := len(grid)
	if 0 == m {
		return res
	}
	for i := 0; i < m; i++ {
		n := len(grid[i])
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				stack := make([][]int, 0)
				count := 0
				stack = append(stack, []int{i, j})
				grid[i][j] = 0 // 入站节点数据置0
				// fmt.Println("-------------------")
				for len(stack) != 0 {
					pop := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					count++
					y := pop[0]
					x := pop[1]
					// fmt.Println("here,y,x:", y, x)

					// fmt.Println("stack:", stack)

					//上存在且为1，则入栈. 且同时元素值设为0，防止重复入栈
					if y-1 >= 0 && 1 == grid[y-1][x] {
						grid[y-1][x] = 0
						stack = append(stack, []int{y - 1, x})

					}

					//下存在，且为1，则入站
					if y+1 < m && 1 == grid[y+1][x] {
						grid[y+1][x] = 0
						stack = append(stack, []int{y + 1, x})
					}

					// 左
					if x-1 >= 0 && 1 == grid[y][x-1] {
						grid[y][x-1] = 0
						stack = append(stack, []int{y, x - 1})
					}

					//右
					if x+1 < n && 1 == grid[y][x+1] {
						grid[y][x+1] = 0
						stack = append(stack, []int{y, x + 1})
					}

					//遍历过的节点清除为0
					// fmt.Println("stack after:", stack)
				}

				// fmt.Println("-------------------")

				if count > res {
					res = count
				}
			}
		}
	}

	return res
}
