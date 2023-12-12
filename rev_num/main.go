package main

import (
	"fmt"
)

func main() {
	var rev, rem, num int
	fmt.Println("Enter the number")
	fmt.Scan(&num)
	for num != 0 {
		rem = num % 10
		rev = rev*10 + rem
		num /= 10
	}
	fmt.Println(rev)
}
