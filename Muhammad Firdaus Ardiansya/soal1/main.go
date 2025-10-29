package main

import "fmt"

func main() {
	var n, hasil int
	fmt.Scan (&n)

	for i := 1; i <= n; i++{
		hasil += 2
		fmt.Print(hasil, " ")
	}
}