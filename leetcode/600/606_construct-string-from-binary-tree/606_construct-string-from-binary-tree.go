/*
link: https://leetcode.com/problems/construct-string-from-binary-tree/

606. Construct String from Binary Tree
Easy

827

1120

Add to List

Share
You need to construct a string consists of parenthesis and integers from a binary tree with the preorder traversing way.

The null node needs to be represented by empty parenthesis pair "()". And you need to omit all the empty parenthesis pairs that don't affect the one-to-one mapping relationship between the string and the original binary tree.

Example 1:
Input: Binary tree: [1,2,3,4]
       1
     /   \
    2     3
   /
  4

Output: "1(2(4))(3)"

Explanation: Originallay it needs to be "1(2(4)())(3()())",
but you need to omit all the unnecessary empty parenthesis pairs.
And it will be "1(2(4))(3)".
Example 2:
Input: Binary tree: [1,2,3,null,4]
       1
     /   \
    2     3
     \
      4

Output: "1(2()(4))(3)"

Explanation: Almost the same as the first example,
except we can't omit the first parenthesis pair to break the one-to-one mapping relationship between the input and the output.


注意：左子为空右子不为空的时候，不可省略括号
这里最开始觉得是不是需要额外使用一个先进先出队列，思考了一下，直接递归处理把

运行结果：
Runtime: 12 ms, faster than 62.50% of Go online submissions for Construct String from Binary Tree.
Memory Usage: 7 MB, less than 66.07% of Go online submissions for Construct String from Binary Tree.

*/

package main

import "strconv"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	res := ""

	if t == nil {
		return res
	}
	c := strconv.Itoa(t.Val)
	res += c

	left := tree2str(t.Left)
	right := tree2str(t.Right)

	if left != "" || right != "" {
		res = res + "(" + left + ")"
	}

	if right != "" {
		res = res + "(" + right + ")"
	}

	return res

}
