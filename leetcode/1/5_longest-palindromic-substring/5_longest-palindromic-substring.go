package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abbc"))
}

func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}

	max := 1

	res := string(s[0])

	dp := make([][]bool,l)
	for i:=0; i <l;i++ {
		dp[i] = make([]bool,l)
		dp[i][i] = true
	}
	for i:=0; i < l-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			
			if 2 > max {
				max = 2
				if i+1 < l-1 {
					res = s[i:i+2]
				} else {
					res = s[i:]
				}
			}
		}
	}


	for span := 3; span <=l; span++ {
		for i:=0; i<= l-span;i++ {
			j := i+span-1
			left := s[i]
			right := s[j]
			dp[i][j] = (left == right && dp[i+1][j-1])

			if dp[i][j] && j-i+1  > max {
				max = j-i+1
				if j < l-1 {
					res = s[i:j+1]
				} else {
					res = s[i:]
				}
			}
		}
	}

	return res


}