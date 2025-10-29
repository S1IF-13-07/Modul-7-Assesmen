package main

import "fmt"

func main() {
	var start, end, hasil int
	fmt.Scan(&start)
	fmt.Scan(&end)
	hasil = 1
	for i := start; i <= end; i++ {
		hasil *= i
	}
	fmt.Print(hasil)
}
