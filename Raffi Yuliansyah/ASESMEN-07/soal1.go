package main

import "fmt"

func main() {
	var n, ganjil int
	fmt.Scan(&n)
	ganjil = 1
	for i := 1; i <= n; i++ {
		fmt.Print(ganjil, " ")
		ganjil += 2
	}
}
