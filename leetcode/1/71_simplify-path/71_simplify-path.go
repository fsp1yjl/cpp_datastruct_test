/*
link: https://leetcode.com/problems/simplify-path/

Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.

In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level.

Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. The last directory name (if it exists) must not end with a trailing /. Also, the canonical path must be the shortest string representing the absolute path.



Example 1:

Input: "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
Example 2:

Input: "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
Example 3:

Input: "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
Example 4:

Input: "/a/./b/../../c/"
Output: "/c"
Example 5:

Input: "/a/../../b/../c//.//"
Output: "/c"
Example 6:

Input: "/a//b////c/d//././/.."
Output: "/a/b/c"

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/../"
	fmt.Println(simplifyPath(s))
}

/*
思路1：
逐个判断字符，用一个堆栈存放常规字符,/,.,
	1.如果为'/':
		如果堆栈最后一位为'/'则进行下一个循环
		如果堆栈内最后一位为'.',则将点退站
	2. 如果为'.':
		如果堆栈内最后一位为'.'，表示返回上级目录，则栈内退至倒数第二个‘/’， 如果不满足两个‘/’，则退至根目录
		如果堆栈顶不是’.', 则将'.' 入堆栈
	3. 其他情况则直接入堆栈
	最后再处理一下结束遍历后，堆栈顶部为/和.的情况即可

	bug: 这里忽略了一个极其恶心的事情，就是超过两个.的组合可以作为路径名，比如：
	/aaa/.../ccc
	放弃思路1 。。。。。

思路2 ：
直接按'/'对字符串进行分割，存入一个字符串数组，再逐个进行上面类似对处理

提交结果：
Runtime: 4 ms, faster than 18.18% of Go online submissions for Simplify Path.
Memory Usage: 4.3 MB, less than 8.18% of Go online submissions for Simplify Path.
Next challenges:

*/

func simplifyPath(path string) string {
	res := "/"

	l := len(path)

	if l == 1 {
		return res
	}
	//FieldsFunc函数对字符串按照特定字符进行分段截取，空段会被抛弃
	paths := strings.FieldsFunc(path, func(s rune) bool {
		if s == '/' {
			return true
		}
		return false
	})

	fmt.Println("dirs:", paths)
	depth := len(paths)

	//p用来存放最终的各段绝对路径名
	p := make([]string, 0)
	for i := 0; i < depth; i++ {
		name := paths[i]
		if name == "." {
			continue
		} else if name == ".." {
			length := len(p)
			if length >= 1 {
				p = p[:length-1]
			}
		} else {
			p = append(p, name)
		}
	}

	// bug1, 要处理最终的绝对路径数组为空的情况，这个时候返回根目录
	if len(p) == 0 {
		return res
	}
	res = ""
	for i := 0; i < len(p); i++ {
		res = res + "/" + p[i]
	}

	return res

}
