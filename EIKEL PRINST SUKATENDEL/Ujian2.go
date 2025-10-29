package main

import "fmt"

func main() {
	var x, y, total int = 0, 0, 1
	fmt.Print("Input x y: ")
	fmt.Scan(&x, &y)

	for i := x; i <= y; i++ {
		total *= i
	}
	fmt.Println(total)
}