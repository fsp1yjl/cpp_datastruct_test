/*
给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

package main

func main() {
	m := make([][]int, 3)
	m[0] = []int{1, 2, 3}
	m[1] = []int{4, 5, 6}
	m[2] = []int{7, 8, 9}

	rotate(m)
}

/*
提交结果：
Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate Image.
Memory Usage: 2.2 MB, less than 15.83% of Go online submissions for Rotate Image.

思路：
由于题目不许使用额外矩阵存储，考虑进行swap
step1: 依次进行四个角上元素swap
step2: 依次进行除角之外对边上元素swap
这里设上右下左四条边为a,b,c,d 思路就是依次swap(a,b), swap(a,c),swap(a,d)


*/
func rotate(matrix [][]int) {
	n := len(matrix)

	if n < 2 {
		return
	}

	left := 0
	right := n - 1

	for left < right {
		//先进行四个角的移动
		matrix[left][left], matrix[left][right] = matrix[left][right], matrix[left][left]
		matrix[left][left], matrix[right][right] = matrix[right][right], matrix[left][left]
		matrix[left][left], matrix[right][left] = matrix[right][left], matrix[left][left]

		//后进行除了角的边上节点移动
		for i := left + 1; i < right; i++ {
			//swap top, right
			matrix[left][i], matrix[i][right] = matrix[i][right], matrix[left][i]

			// fmt.Println(matrix)

			// swap top,down
			matrix[left][i], matrix[n-1-left][n-1-i] = matrix[n-1-left][n-1-i], matrix[left][i]

			// fmt.Println(matrix)

			// swap top,left
			matrix[left][i], matrix[n-1-i][left] = matrix[n-1-i][left], matrix[left][i]
			// fmt.Println(matrix)

		}
		left++
		right--
	}
	// fmt.Println(matrix)
}
