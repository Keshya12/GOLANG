package main

import (
	"fmt"
)

func main() {
	var str string
	var flag bool // Change the flag type to bool

	fmt.Println("Enter the combination")
	fmt.Scan(&str)

	n := len(str)
	flag = false // Initialize the flag as false before checking

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ { // Corrected loop condition j < n
			if str[i] == str[j] {
				flag = true // Set flag to true if a duplicate character is found
				break       // Break out of the loop if duplicate is found
			}
		}
		if flag { // Check flag value after inner loop
			break // Break the outer loop if duplicate is found
		}
	}

	if flag {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
