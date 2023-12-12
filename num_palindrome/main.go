package main

import (
	"fmt"
)

func main() {
	var num int
	var rev int
	fmt.Println("Enter the number")
	fmt.Scan(&num)
	ori := num
	for num != 0 {
		rem := num % 10
		rev = rev*10 + rem
		num /= 10
	}
	if rev == ori {
		fmt.Println("Yes")
	} else {
		fmt.Println("NO")
	}
}
