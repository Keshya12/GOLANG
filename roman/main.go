package main

import (
	"fmt"
	"strings"
)

func romanToDecimal(roman string) int {
	romanNumerals := map[byte]int{
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

	for i := len(roman) - 1; i >= 0; i-- {
		current := romanNumerals[roman[i]]
		if current < prev {
			result -= current
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

	// Convert Roman numeral to decimal number
	result := romanToDecimal(strings.ToUpper(s))
	fmt.Println("Decimal number:", result)
}
