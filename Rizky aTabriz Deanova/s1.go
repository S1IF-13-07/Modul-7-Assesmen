package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	i := 0
	x := 1
	for i < n {
		if x%2 != 0 {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(x)
			i++
		}
		x++
	}
	fmt.Println()
}
