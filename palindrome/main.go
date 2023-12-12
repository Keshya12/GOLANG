package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var start, end int
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // Check palindrome for odd length substrings
		len2 := expandAroundCenter(s, i, i+1) // Check palindrome for even length substrings
		length := max(len1, len2)

		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println("Longest palindromic substring for 'babad':", longestPalindrome("babad")) // Output: "bab" or "aba"
	fmt.Println("Longest palindromic substring for 'cbbd':", longestPalindrome("cbbd"))   // Output: "bb"
}
