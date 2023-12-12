package main

import (
	"fmt"
	"strings"
)

func romanToDecimal(roman string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0

	for _, r := range roman {
		current := romanNumerals[r]
		if current > prev {
			result += current - 2*prev
		} else {
			result += current
		}
		prev = current
	}

	return result
}

func main() {
	var s string

	fmt.Println("Enter the Roman numeral:")
	fmt.Scan(&s)

	result := romanToDecimal(strings.ToUpper(s))
	fmt.Println("Decimal number:", result)
}
