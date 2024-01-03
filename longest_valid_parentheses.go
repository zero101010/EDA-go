//https://leetcode.com/problems/longest-valid-parentheses/
package main

import "fmt"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// public int longestValidParentheses(String s) {
// 	int maxans = 0;
// 	int dp[] = new int[s.length()];
// 	for (int i = 1; i < s.length(); i++) {
// 		if (s.charAt(i) == ')') {
// 			if (s.charAt(i - 1) == '(') {
// 				dp[i] = (i >= 2 ? dp[i - 2] : 0) + 2;
// 			} else if (i - dp[i - 1] > 0 && s.charAt(i - dp[i - 1] - 1) == '(') {
// 				dp[i] = dp[i - 1] + ((i - dp[i - 1]) >= 2 ? dp[i - dp[i - 1] - 2] : 0) + 2;
// 			}
// 			maxans = Math.max(maxans, dp[i]);
// 		}
// 	}
// 	return maxans;
// }

func longestValidParentheses(s string) int {
	maxans := 0
	var dp = []int{}
	for i, v := range s {
		fmt.Println(string(v))
		if string(v) == ")" {
			if string(s[i-1]) == "(" {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 0 + 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if dp[i-1]+(i-dp[i-1]) >= 2 {
					dp[i] = dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = 0 + 2
				}
			}
			maxans = Max(maxans, dp[i])
		}
	}
	return maxans
}
func main() {
	longestValidParentheses("(()")
}
