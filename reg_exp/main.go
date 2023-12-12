package main

import (
	"fmt"
)

func isMatch(s, p string) bool {
	if p == "" {
		return s == ""
	}

	match := len(s) > 0 && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (match && isMatch(s[1:], p))
	} else {
		return match && isMatch(s[1:], p[1:])
	}
}

func main() {
	var inputStr, pattern string

	fmt.Println("Enter the input string:")
	fmt.Scanln(&inputStr)

	fmt.Println("Enter the pattern to match:")
	fmt.Scanln(&pattern)

	match := isMatch(inputStr, pattern)

	if match {
		fmt.Println("The pattern matches the entire input string.")
	} else {
		fmt.Println("The pattern does not match the entire input string.")
	}
}
