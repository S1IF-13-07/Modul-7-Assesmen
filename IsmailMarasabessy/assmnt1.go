package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	fmt.Print("Output: ")
	for i := 1; i <= n; i++ {
		fmt.Print(2*i - 1)
		if i < n {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
