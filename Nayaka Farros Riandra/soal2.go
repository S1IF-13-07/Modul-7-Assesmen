package main

import "fmt"

func main() {
	var a, b int
	var hasil int

	fmt.Scan(&a)
	fmt.Scan(&b)

	hasil =1
	for i := a; i <= b; i++ {
		hasil *= i
	}
	fmt.Print(hasil)
}