/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false


提示：

board 和 word 中只包含大写和小写英文字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

package l79

// import "fmt"


func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	// l := len(word)

	// var helper func(int, int, int ) bool


	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			// visited := make([][]bool, m)
			// for i := 0; i < m; i++ {
			// 	visited[i] = make([]bool, n)
			// 	for j := 0; j < n; j++ {
			// 		visited[i][j] = false
			// 	}
			// }
			ret := helper(i, j, 0,board, word)
			if true == ret {
				return true
			}
		}
	}
	return false

}

func helper(i int, j int, index int, board [][]byte, word string) bool {

	m := len(board)
	n := len(board[0])
	l := len(word)

	if i < 0 || i >= m || j < 0 || j >= n {
		//索引越界，返回false
		return false
	}

	// i,j表示board当前开始元素， index 表示对应word的字符索引
	if board[i][j] != word[index] {
		//首字符不同，可能已经遍历过，或者字符不匹配，直接返回false
		return false
	}

	//如果字符相同，且为最后一个word元素，则返回true
	if index == l-1 {
		return true
	}

	//首字符相同,且匹配未结束，则对上下左右进行递归
	board[i][j] = '#'
	
	// 注意，这里不要进行top ,right,down, left 四个结果取或返回的操作，这样会成倍放大计算次数
	top := helper(i-1, j, index+1, board, word)
	if  top {
		return true 
	}
	right := helper(i, j+1, index+1,board, word )
	if right {
		return true 
	}
	down := helper(i+1, j, index+1, board, word )
	if down {
		return true 
	}
	left := helper(i, j-1, index+1, board, word)
	if left {
		return true 
	}

	// ret := top || right || down || left

	// 如果上下左右均无解，i，j重置,返回false
	board[i][j] = word[index]

	return false
}



