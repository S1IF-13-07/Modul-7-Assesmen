package main

import "fmt"

func main() {
	var x, y int64
	fmt.Scan(&x, &y)
	var result int64 = 1
	for i := x; i <= y; i++ {
		result *= i
	}
	fmt.Println(result)
}
