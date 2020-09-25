/*
link:  https://leetcode.com/problems/valid-palindrome/

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false


Constraints:

s consists only of printable ASCII characters.

*/

package main

import (
	"fmt"
	"strings"
)

func main () {
	// c := ':'
	// // b:= 'a'
	
	// if (c >= 'a' &&  c <= 'z') || (c >= 'A'  &&  c <=  'z' ) {
	// 	fmt.Println("dddd")
	// }
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0p"))


}

/*

运行结果：
Runtime: 592 ms, faster than 5.18% of Go online submissions for Valid Palindrome.
Memory Usage: 52.4 MB, less than 5.18% of Go online submissions for Valid Palindrome.

从结果看，这种方式时间和空间效率都不高

*/
func isPalindrome(s string) bool {
	    // bug1 题目中声明忽略大小写，这里先转换字符串全部为小写
	s = strings.ToLower(s)
	ss := ""
	l := len(s)
	for i:=0; i < l; i++ {
		c := s[i]
		if (c >= 'a' &&  c <= 'z')  || (c >= '0' && c <= '9') {
			ss += string(c)
		}
	}

	// fmt.Println("ss:", ss)
	l = len(ss)
    //  fmt.Println("lll:", l)
	for i:=0; i < l/2; i++ {
		left := ss[i]
		right := ss[l-i-1]
		// fmt.Println("left:", string(left))
		// fmt.Println("right:", string(right))

		if left != right {
			return false
		}
	}
	return true 
}