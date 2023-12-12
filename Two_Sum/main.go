package main

import (
	"fmt"
)

func main() {
	var target int
	var n int

	fmt.Println("Enter the range of array")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Print("Enter the target: ")
	fmt.Scan(&target)

	fmt.Print("Enter the array: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		sum := arr[i]
		for j := i + 1; j <= n; j++ {
			if sum == target {
				fmt.Printf("Subarray found with sum %d at positions %d and %d\n", target, i, j-1)
				return // Terminate after finding the subarray
			}
			if sum > target || j == n {
				break
			}
			sum += arr[j]
		}
	}

	fmt.Printf("Subarray with sum %d not found.\n", target)
}
