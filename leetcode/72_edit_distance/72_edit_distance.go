/*

link: https://leetcode.com/problems/edit-distance/


Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/

package main

import "fmt"

func main() {
	// 预期值 7
	word1 := "dinitrophenylhydrazine"
	word2 := "benzalphenylhydrazone"

	// word1 := "horse"
	// word2 := "ros"

	fmt.Println("min:", minDistance(word1, word2))

}

func minDistance(word1 string, word2 string) int {

	lastI := len(word1)
	lastJ := len(word2)

	dpMem := make([][]int, lastI+1)
	for i := 0; i <= lastI; i++ {
		dpMem[i] = make([]int, lastJ+1)
	}

	for i := 0; i <= lastI; i++ {
		for j := 0; j <= lastJ; j++ {
			dpMem[i][j] = -1
		}
	}

	// fmt.Println("dpmem:", dpMem)

	// fmt.Println("lasti:", lastI)
	// fmt.Println("lastJ:", lastJ)

	return dp(dpMem, word1, word2, lastI, lastJ)
}

func dp(dpMem [][]int, word1 string, word2 string, i int, j int) int {
	// 这里增加了一个对已经处理子问题结果对缓存，否则，会递归多次处理同一个子问题，严重减慢计算速度
	if dpMem[i][j] != -1 {
		return dpMem[i][j]
	}
	if 0 == i {
		dpMem[i][j] = j
		return j
	}

	if 0 == j {
		dpMem[i][j] = i
		return i
	}
	if word1[i-1] == word2[j-1] {
		tmp := dp(dpMem, word1, word2, i-1, j-1)
		dpMem[i][j] = tmp
		return tmp
	}

	//如果i,j对应的字符不同，则有三种方式
	dp_insert := dp(dpMem, word1, word2, i, j-1) + 1    //word1插入一个word2[j-1]
	dp_replace := dp(dpMem, word1, word2, i-1, j-1) + 1 //word1替换i字符为word2[j-1]
	dp_delete := dp(dpMem, word1, word2, i-1, j) + 1    // word1删除i,使用i-1与j进行比较

	// fmt.Println("i:", i, "j:", j)
	dp_min := dp_insert
	if dp_replace < dp_min {
		dp_min = dp_replace
	}

	if dp_delete < dp_min {
		dp_min = dp_delete
	}

	dpMem[i][j] = dp_min

	return dp_min
}
